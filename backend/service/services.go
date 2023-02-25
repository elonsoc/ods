package service

import "log"

// Service, here, describes
type Service struct {
	Logger *log.Logger
}

func NewService() *Service {
	return &Service{}
}
