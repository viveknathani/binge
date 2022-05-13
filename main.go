package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/viveknathani/binge/database"
	"github.com/viveknathani/binge/processor"
	"github.com/viveknathani/binge/server"
	"github.com/viveknathani/binge/service"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Hold environment variables
var (
	port           string = ""
	databaseServer string = ""
)

// Setup environment variables
func init() {

	port = os.Getenv("PORT")
	databaseServer = os.Getenv("DATABASE_URL")
}

// getLogger will configure and return a uber/zap logger
func getLogger() *zap.Logger {

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevel(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "ts",
			EncodeTime:  zapcore.EpochMillisTimeEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return logger
}

// getDatabase will init and return a db
func getDatabase() *database.Database {

	db := &database.Database{}
	err := db.Initialize(databaseServer)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return db
}

func main() {

	logger := getLogger()
	db := getDatabase()
	pro := processor.New(4, 10)

	// fire the processor
	go pro.Run()

	// Setup the web server
	srv := &server.Server{
		Service: &service.Service{
			Database:  db,
			Logger:    logger,
			Processor: pro,
		},
		Router: mux.NewRouter(),
	}
	srv.SetupRoutes()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Listen
	go func() {

		err := http.ListenAndServe(":"+port, srv)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}()
	fmt.Println("Server started!")
	<-done
	shutdown(srv, db, pro)
}

func shutdown(srv *server.Server, db *database.Database, pro *processor.Processor) {

	err := db.Close()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("goodbye!")
}
