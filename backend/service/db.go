package service

import (
	"context"

	pgx "github.com/jackc/pgx/v5"
)

// DbIFace is an interface for the database
type DbIFace interface {
	GetConn() *pgx.Conn
	NewApp(string, string, string, string, string, string, bool) error
	UserApps() (pgx.Rows, error)
	CheckDuplicate(string, string) (bool, error)
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
	_, err = connection.Prepare(ctx, "insert_into_keys", "INSERT INTO keys (app_ID, api_key, is_valid) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = connection.Prepare(ctx, "insert_into_applications", "INSERT INTO applications (app_ID, app_name, description, owners, team_name) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		log.Fatal(err)
	}
	return &Db{db: connection}
}

func (s *Db) GetConn() *pgx.Conn {
	return s.db
}

// NewApp stores the information about a new application in the database.
func (db *Db) NewApp(name string, ID string, desc string, owners string, tname string, key string, valid bool) error {
	ctx := context.Background()
	// Initiating a transaction
	tx, err := db.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Storing info in the keys table
	_, err = tx.Exec(ctx, "insert_into_keys", ID, key, valid)
	if err != nil {
		return err
	}

	// Storing info in the applications table
	_, err = tx.Exec(ctx, "insert_into_applications", ID, name, desc, owners, tname)
	if err != nil {
		return err
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

// UserApps returns a list of all the applications that exist right now.
// Once accessing the users email is figured out, this function will return
// a list of all the applications that the user owns.
func (db *Db) UserApps() (pgx.Rows, error) {
	ctx := context.Background()
	rows, err := db.db.Query(ctx, "SELECT * FROM applications")
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Checks for duplicated app IDs or api keys
// I easily could've used if statements but I've always wanted to use switch statements
// and realistically this could expand to more columns in the future.
func (db *Db) CheckDuplicate(column string, newGen string) (bool, error) {
	ctx := context.Background()
	var query string

	// This switch statement is used to determine which column to check for duplicates
	switch column {
	case "app_ID":
		query = "SELECT * FROM applications WHERE app_ID = $1"
	case "api_key":
		query = "SELECT * FROM applications WHERE api_key = $1"
	}

	// Querying the database for duplicates
	rows, err := db.db.Query(ctx, query, newGen)
	if err != nil {
		return false, err
	}

	// If there are no duplicates, return true
	if rows.Next() {
		return false, nil
	} else {
		return true, nil
	}
}
