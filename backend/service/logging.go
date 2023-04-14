package service

import (
	logrusLoki "github.com/schoentoon/logrus-loki"
	"github.com/sirupsen/logrus"
)

func initLogging(loggingURL string) *Log {
	log := logrus.New()
	hook, err := logrusLoki.NewLokiDefaults(loggingURL)
	if err != nil {
		log.Fatal(err)
	}
	log.AddHook(hook)
	return &Log{logger: log}
}

type LoggerIFace interface {
	Info(...string)
	InfoWithFields(logrus.Fields, ...string)
	Error(...string)
}

type Log struct {
	logger *logrus.Logger
}

func (s *Log) Info(message ...string) {
	s.logger.Info(message)
}

func (s *Log) InfoWithFields(fields logrus.Fields, message ...string) {
	s.logger.WithFields(fields).Info(message)
}

func (s *Log) Error(message ...string) {
	s.logger.Error(message)
}
