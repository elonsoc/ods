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
	err = prepareStatements(connection, ctx)
	if err != nil {
		log.Fatal(err)
	}

	return &Db{db: connection}
}

func prepareStatements(connection *pgx.Conn, ctx context.Context) (err error) {
	_, err = connection.Prepare(ctx, "insert_into_applications", "INSERT INTO applications (app_ID, app_name, description, owners, team_name, api_key, is_valid) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}
	return nil
}

func (s *Db) GetConn() *pgx.Conn {
	return s.db
}

// NewApp stores the information about a new application in the database.
func (db *Db) NewApp(name string, ID string, desc string, owners string, tname string, key string, valid bool) error {
	ctx := context.Background()

	// Storing all new app info into the applications table.
	db.db.Exec(ctx, "insert_into_applications", ID, name, desc, owners, tname, key, valid)

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
	err := db.db.QueryRow(ctx, query, newGen).Scan()
	if err == pgx.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
