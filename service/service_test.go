package service

import (
	"log"
	"os"
	"testing"

	"github.com/viveknathani/binge/database"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const dsn = "postgres://viveknathani:root@localhost:5432/binge?sslmode=disable"

var service *Service

func TestMain(t *testing.M) {

	service = &Service{}
	db := &database.Database{}
	err := db.Initialize(dsn)
	if err != nil {
		log.Fatal(err)
	}
	service.Database = db
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
	logger, _ := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	service.Logger = logger
	code := t.Run()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
	_ = logger.Sync()
	os.Exit(code)
}
