package service

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func initDb(databaseURL string, log *logrus.Logger) *Db {
	connection, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	return &Db{db: connection}
}

type DbIFace interface {
	GetConn() *pgx.Conn
}

type Db struct {
	db *pgx.Conn
}

func (s *Db) GetConn() *pgx.Conn {
	return s.db
}
