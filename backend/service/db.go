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
	_, err := connection.Prepare(ctx, "insert_into_keys", "INSERT INTO keys (app_ID, api_key, is_valid) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}

	_, err := connection.Prepare(ctx, "insert_into_applications", "INSERT INTO applications (app_ID, app_name, description, owners, team_name) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		log.Fatal(err)
	}
	return &Db{db: connection}
}

type DbIFace interface {
	GetConn() *pgx.Conn
	NewApp()  error
}

type Db struct {
	db *pgx.Conn
}

func (s *Db) GetConn() *pgx.Conn {
	return s.db
}

func (db *Db) NewApp() error {
	return nil
}