package service

import (
	"context"

	"github.com/elonsoc/ods/backend/applications/types"
	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DbIFace is an interface for the database
type DbIFace interface {
	NewApp(string, string) (string, error)
	GetApplications() ([]types.Application, error)
	GetApplication(string) (types.Application, error)
	UpdateApplication(string, types.Application) error
	DeleteApplication(string) error
}

// Db is a struct that contains a pointer to the database connection
type Db struct {
	db *pgxpool.Pool
}

func initDb(databaseURL string, log LoggerIFace) *Db {
	ctx := context.Background()

	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	cfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		err := prepareStatements(ctx, conn)
		return err
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &Db{db: pool}
}

func prepareStatements(ctx context.Context, connection *pgx.Conn) error {
	_, err := connection.Prepare(ctx, "insert_into_applications",
		"INSERT INTO applications (name, description, is_valid) VALUES ($1, $2, true) RETURNING id")
	if err != nil {
		return err
	}
	_, err = connection.Prepare(ctx, "select_all_applications",
		"SELECT api_key, description, id, is_valid, name FROM applications")
	if err != nil {
		return err
	}

	_, err = connection.Prepare(ctx, "select_application",
		"SELECT api_key, description, id, is_valid, name FROM applications WHERE id=$1")
	if err != nil {
		return err
	}

	_, err = connection.Prepare(ctx, "get_owners",
		"SELECT email FROM users WHERE id IN (SELECT user_id FROM app_owner WHERE app_id=$1)")
	if err != nil {
		return err
	}
	_, err = connection.Prepare(ctx, "update_application",
		"UPDATE applications SET name=$1, description=$2 WHERE id=$3")
	if err != nil {
		return err
	}

	_, err = connection.Prepare(ctx, "first_leg_delete_application",
		"DELETE FROM app_owner WHERE app_id=$1")
	if err != nil {
		return err
	}

	_, err = connection.Prepare(ctx, "second_leg_delete_application",
		"DELETE FROM applications WHERE id=$1")
	if err != nil {
		return err
	}
	return nil
}

// NewApp stores the information about a new application in the database.

func (db *Db) NewApp(name string, desc string) (string, error) {
	ctx := context.Background()

	var id string

	err := db.db.QueryRow(ctx, "insert_into_applications", name, desc).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

// UserApps returns a list of all the applications that exist right now.
// Once accessing the users email is figured out, this function will return
// a list of all the applications that the user owns.
func (db *Db) GetApplications() ([]types.Application, error) {
	ctx := context.Background()
	rows, err := db.db.Query(ctx, "select_all_applications")
	if err != nil {
		return nil, err
	}
	apps := []types.Application{}
	for rows.Next() {
		var app types.Application
		err = rows.Scan(&app.ApiKey, &app.Description, &app.Id, &app.IsValid, &app.Name)
		if err != nil {
			return nil, err
		}

		apps = append(apps, app)
	}
	return apps, nil
}

func (db *Db) GetApplication(applicationId string) (types.Application, error) {
	var app types.Application
	err := db.db.QueryRow(context.Background(), "select_application", applicationId).Scan(
		&app.ApiKey, &app.Description, &app.Id, &app.IsValid, &app.Name)
	if err != nil {
		return app, err
	}

	rows, err := db.db.Query(context.Background(), "get_owners", applicationId)
	if err != nil {
		return app, err
	}

	if rows.Next() {
		var email string
		err = rows.Scan(&email)
		if err != nil {
			return app, err
		}
		app.Owners = email
	}

	return app, nil
}

func (db *Db) UpdateApplication(applicationId string, applicationInfo types.Application) error {
	_, err := db.db.Exec(context.Background(), "update_application",
		applicationInfo.Name,
		applicationInfo.Description,
		applicationId)
	return err
}

func (db *Db) DeleteApplication(applicationId string) error {
	txContext := context.Background()
	tx, err := db.db.BeginTx(txContext, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback(txContext)

	_, err = tx.Exec(txContext, "first_leg_delete_application", applicationId)
	if err != nil {
		return err
	}
	tx.Exec(txContext, "second_leg_delete_application", applicationId)
	if err != nil {
		return err
	}
	err = tx.Commit(txContext)
	if err == nil {
		return nil
	}
	return err
}
