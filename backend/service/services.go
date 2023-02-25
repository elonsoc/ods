package service

import "log"

type Service struct {
	Logger *log.Logger
}

func NewService(logger *log.Logger) *Service {
	return &Service{Logger: logger}
}
