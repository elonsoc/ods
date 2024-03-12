package common

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	logrusLoki "github.com/schoentoon/logrus-loki"
	"github.com/sirupsen/logrus"
	statsd "github.com/smira/go-statsd"
)

func InitLogging(loggingURL string) *Log {
	log := logrus.New()
	hook, err := logrusLoki.NewLokiDefaults(loggingURL)
	if err != nil {
		log.Fatal(err)
	}
	log.AddHook(hook)
	return &Log{logger: log}
}

type LoggerIFace interface {
	Info(string, logrus.Fields)
	Error(string, logrus.Fields)
	Debug(string, logrus.Fields)
	Fatal(error)
}

type Log struct {
	logger *logrus.Logger
}

func (s *Log) Debug(message string, fields logrus.Fields) {
	s.logger.WithFields(fields).Debug(message)
}

func (s *Log) Info(message string, fields logrus.Fields) {
	s.logger.WithFields(fields).Info(message)
}

func (s *Log) Error(message string, fields logrus.Fields) {
	s.logger.WithFields(fields).Error(message)
}

func (s *Log) Fatal(err error) {
	s.logger.Fatal(err)
}

func CustomLogger(log LoggerIFace, stat StatIFace) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// pass along the http request before we log it
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			scheme := "http"
			if r.TLS != nil {
				scheme = "https"
			}

			log.Info("", logrus.Fields{
				"method":     r.Method,
				"path":       r.URL.Path,
				"request_id": middleware.GetReqID(r.Context()),
				"ip":         r.RemoteAddr,
				"scheme":     scheme,
				"status":     ww.Status(),
			})
			stat.Increment("request", statsd.IntTag("status", ww.Status()), statsd.StringTag("path", r.URL.Path))
		}

		return http.HandlerFunc(fn)
	}
}
