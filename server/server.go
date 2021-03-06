package server

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/viveknathani/binge/service"
	"github.com/viveknathani/binge/shared"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Server holds together all the configuration needed to run this web service.
type Server struct {
	Service *service.Service
	Router  *mux.Router
}

func (s *Server) setupContent(directory string) {

	fileServer := http.FileServer(http.Dir(directory))
	s.Router.PathPrefix("/" + directory + "/").Handler(http.StripPrefix("/"+directory, fileServer))
}

// ServeHTTP is implemented so that Server can be used for listening to requests.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	requestID := uuid.New().String()
	request := r.Clone(shared.WithRequestID(context.Background(), requestID))
	showRequestMetaData(s.Service.Logger, request)
	s.Router.ServeHTTP(w, request)
}

func zapReqID(r *http.Request) zapcore.Field {

	return zapcore.Field{
		Key:    "requestID",
		String: shared.ExtractRequestID(r.Context()),
		Type:   zapcore.StringType,
	}
}

func showRequestMetaData(l *zap.Logger, r *http.Request) {

	reqMethod := zapcore.Field{
		Key:    "method",
		String: r.Method,
		Type:   zapcore.StringType,
	}

	reqPath := zapcore.Field{
		Key:    "path",
		String: r.URL.String(),
		Type:   zapcore.StringType,
	}

	l.Info("incoming request", zapReqID(r), reqMethod, reqPath)
}
