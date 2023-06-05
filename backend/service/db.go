package service

import (
	"context"

	pgx "github.com/jackc/pgx/v5"
)

// DbIFace is an interface for the database
type DbIFace interface {
	GetConn() *pgx.Conn
	NewApp(string, string) (string, error)
	GetApplications() (pgx.Rows, error)
}

// Db is a struct that contains a pointer to the database connection
type Db struct {
	db *pgx.Conn
}

func initDb(databaseURL string, log LoggerIFace) *Db {
	ctx := context.Background()
	connection, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	err = prepareStatements(connection, ctx)
	if err != nil {
		log.Fatal(err)
	}

	return &Db{db: connection}
}

func prepareStatements(connection *pgx.Conn, ctx context.Context) (err error) {
	_, err = connection.Prepare(ctx, "insert_into_applications", "INSERT INTO applications (name, description, is_valid) VALUES ($1, $2, true) RETURNING id")
	_, err = connection.Prepare(ctx, "select_all_applications", "SELECT * FROM applications")
	if err != nil {
		return err
	}
	return nil
}

func (s *Db) GetConn() *pgx.Conn {
	return s.db
}

// NewApp stores the information about a new application in the database.
func (db *Db) NewApp(name string, desc string) (string, error) {
	ctx := context.Background()

	var id string

	// Storing all new app info into the applications table.
	err := db.db.QueryRow(ctx, "insert_into_applications", name, desc).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

// UserApps returns a list of all the applications that exist right now.
// Once accessing the users email is figured out, this function will return
// a list of all the applications that the user owns.
func (db *Db) GetApplications() (pgx.Rows, error) {
	ctx := context.Background()
	rows, err := db.db.Query(ctx, "SELECT * FROM applications")
	if err != nil {
		return nil, err
	}
	return rows, nil
}
