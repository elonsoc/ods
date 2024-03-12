package common

import (
	logrusLoki "github.com/schoentoon/logrus-loki"
	"github.com/sirupsen/logrus"
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
