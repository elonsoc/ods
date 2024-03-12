package common

import (
	"context"
	"fmt"

	"github.com/elonsoc/ods/src/common/types"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	errors "github.com/pkg/errors"
)

// DbIFace is an interface for the database
type DbIFace interface {
	NewApp(string, string, string) (string, error)
	GetApplications(string) ([]types.Application, error)
	GetApplication(string) (types.Application, error)
	UpdateApplication(string, types.Application) error
	DeleteApplication(string) error
	NewUser(string, string, string, string, string) error
	IsUser(string) bool
	IsValidApiKey(string) bool
	GetUserInformation(string) (*ExternalUserInformation, error)
}

type ExternalUserInformation struct {
	FirstName   string `json:"given_name"`
	LastName    string `json:"family_name"`
	Email       string `json:"email"`
	OdsId       string `json:"id"`
	Affiliation string `json:"affiliation"`
}

// Db is a struct that contains a pointer to the database connection
type Db struct {
	db *pgxpool.Pool
}

func InitDb(databaseURL string, log LoggerIFace) *Db {
	ctx := context.Background()

	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	cfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		err := prepareStatements(ctx, conn)
		pgxuuid.Register(conn.TypeMap())
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

	_, err = connection.Prepare(ctx, "relate_new_application", "INSERT INTO app_owner (app_id, user_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}

	_, err = connection.Prepare(ctx, "select_all_applications_by_user",
		"SELECT api_key, description, id, is_valid, name FROM applications WHERE id IN (SELECT app_id FROM app_owner WHERE user_id = $1);")
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

	_, err = connection.Prepare(ctx, "delete_app_owner_relation",
		"DELETE FROM app_owner WHERE app_id=$1")
	if err != nil {
		return err
	}

	_, err = connection.Prepare(ctx, "delete_application",
		"DELETE FROM applications WHERE id=$1")
	if err != nil {
		return err
	}

	_, err = connection.Prepare(ctx, "create_new_user", "INSERT INTO users (email, given_name, family_name, affiliation) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return err
	}

	_, err = connection.Prepare(ctx, "relate_elon_ods", "INSERT INTO elon_ods (elon_id, ods_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	return nil
}

// NewApp stores the information about a new application in the database.

func (db *Db) NewApp(name string, desc string, userId string) (string, error) {
	ctx := context.Background()

	tx, err := db.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return "", err
	}

	defer tx.Rollback(ctx)

	var id string

	err = tx.QueryRow(ctx, "insert_into_applications", name, desc).Scan(&id)
	if err != nil {
		return "", err
	}

	_, err = tx.Exec(ctx, "relate_new_application", id, userId)
	if err != nil {
		return "", err
	}

	tx.Commit(ctx)

	return id, nil
}

// UserApps returns a list of all the applications that exist right now.
// Once accessing the users email is figured out, this function will return
// a list of all the applications that the user owns.
func (db *Db) GetApplications(userId string) ([]types.Application, error) {
	ctx := context.Background()
	rows, err := db.db.Query(ctx, "select_all_applications_by_user", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
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

	defer rows.Close()

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

	_, err = tx.Exec(txContext, "delete_app_owner_relation", applicationId)
	if err != nil {
		return err
	}
	tx.Exec(txContext, "delete_application", applicationId)
	if err != nil {
		return err
	}
	err = tx.Commit(txContext)
	if err == nil {
		return nil
	}
	return err
}

func (db *Db) NewUser(uid, givenName, surname, email, affiliation string) error {
	txContext := context.Background()
	tx, err := db.db.BeginTx(txContext, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback(txContext)

	var ods_id_uuid pgxuuid.UUID
	err = tx.QueryRow(txContext, "create_new_user", email, givenName, surname, affiliation).Scan(&ods_id_uuid)
	if err != nil {
		return err
	}

	_, err = tx.Exec(txContext, "relate_elon_ods", uid, ods_id_uuid)
	if err != nil {
		return err
	}

	err = tx.Commit(txContext)
	if err != nil {
		return err
	}

	return nil

}

func (db *Db) IsUser(elon_id string) bool {
	var count int
	// this is nasty and not cool
	err := db.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM elon_ods where elon_id=$1", elon_id).Scan(&count)
	if count == 0 || err != nil {
		// this might be a bad idea but in my mind if we error here then the user doesn't exist
		// but we could, in fact, error for other reasons.
		return false
	}
	return true
}

func (db *Db) IsValidApiKey(key string) bool {
	if key == "" {
		return false
	}

	var count int
	// this is nasty and not cool
	err := db.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM applications WHERE api_key=$1", key).Scan(&count)
	if count == 0 || err != nil {
		return false
	}

	return true
}

func (db *Db) GetUserInformation(elon_id string) (*ExternalUserInformation, error) {
	if elon_id == "" {
		return nil, errors.New("elon_id is empty")
	}

	var userInfo ExternalUserInformation
	err := db.db.QueryRow(context.Background(),
		"SELECT u.id, u.email, u.given_name, u.family_name, u.affiliation FROM users u JOIN elon_ods eo ON u.id = eo.ods_id WHERE eo.elon_id = $1",
		elon_id).Scan(&userInfo.OdsId, &userInfo.Email, &userInfo.FirstName, &userInfo.LastName, &userInfo.Affiliation)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not get user information for elonId %s: %v", elon_id, err))
	}

	return &userInfo, nil
}
