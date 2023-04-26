package service

import (
	"context"

	pgx "github.com/jackc/pgx/v5"
)

func initDb(databaseURL string, log LoggerIFace) *Db {
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
