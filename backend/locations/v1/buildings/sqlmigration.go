package migration

// import MSSQL and PSQL drivers and libraries that allow us to connect
import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

func migrate() {
	// connecting to MSSQL Server Database (Elon's in the future)
	// MSSQL connection string w/ server address, user, pass, database name
	mssqlConnString := "server=localhost;user id = sa; password=<odspass12345>;database=[databasename]"
	// mssqlconn holds the connection to the mssql server
	// mssqlconn is used for reference to exectute commands on database
	mssqlConn, err := sql.Open("sqlserver", mssqlConnString)
	// sql.Open returns an error value nil if connection success
	// if connections fails, err is set to non-nil
	if err != nil {
		log.Fatal("Error connecting to MSSQL Server:", err.Error())
		return
	}
	// defer closing the connection of database until end of migration/function
	defer mssqlConn.Close()

	// migrating data from MSSQL to PSQL using a SQL query to MSSQL database
	// query selects all columns from data table
	mssqlQuery := "SELECT * FROM elonbuildings"
	// mssqlRows will hold all rows from data table retrieved from query
	mssqlRows, err := mssqlConn.Query(mssqlQuery)
	if err != nil {
		log.Fatal("Error querying MSSQL Server:", err.Error())
		return
	}
	defer mssqlRows.Close()

	// define connection string for PSQL database
	// connecting to a localhost psql database here
	pgConnString := "postgres://postgres:odspass12345@localhost/postgres?sslmode=disable"
	// connecting to PostgreSQL server using Go driver from package
	pgConn, err := sql.Open("postgres", pgConnString)
	// checking if connection is unsuccessful
	if err != nil {
		log.Fatal("Error connecting to PSQL Server", err.Error())
	}
	defer pgConn.Close()

	// create PSQL database w/ same schema as database
	// empty _ because Exec returns both a sql result obj & err value
	// we're evaluating err value, hence the _
	_, err = pgConn.Exec(`
    CREATE TABLE elonbuildingspg (
        buildings_id varchar(4) NOT NULL,
        bldg_desc varchar(255) NULL,
        bldg_location varchar(5) NULL,
        bldg_location_representation varchar(30) NULL,
        bldg_type varchar(10) NULL,
        bldg_type_representation varchar(32) NULL,
        bldg_long_desc varchar(1996) NULL,
        bldg_city varchar(25) NULL,
        bldg_state varchar(2) NULL,
        bldg_zip varchar(10) NULL,
        bldg_sector varchar(10) NULL,
        bldg_latitude numeric NULL,
        bldg_longitude numeric NULL,
        buildings_add_date timestamp NULL,
        buildings_chgdate timestamp NULL,
        PRIMARY KEY (buildings_id)
    );
`)

	// if error, logs error message and termianates
	if err != nil {
		log.Fatal("Error creating PostgreSQL table: ", err.Error())
	}

	// inserting data into PSQL database
	// for loop to iterate over rows of MSSQL

	for mssqlRows.Next() {
		// declaration of variables to hold values of each column in row
		var buildings_id string
		var bldg_desc sql.NullString
		var bldg_location sql.NullString
		var bldg_location_representation sql.NullString
		var bldg_type sql.NullString
		var bldg_type_representation sql.NullString
		var bldg_long_desc sql.NullString
		var bldg_city sql.NullString
		var bldg_state sql.NullString
		var bldg_zip sql.NullString
		var bldg_sector sql.NullString
		var bldg_latitude sql.NullFloat64
		var bldg_longitude sql.NullFloat64
		var buildings_add_date sql.NullTime
		var buildings_chgdate sql.NullTime
		// scan method retrieves values from MSSQL rows & assings to Go's memory
		err = mssqlRows.Scan(&buildings_id, &bldg_desc, &bldg_location,
			&bldg_location_representation, &bldg_type,
			&bldg_type_representation, &bldg_long_desc, &bldg_city,
			&bldg_state, &bldg_zip, &bldg_sector, &bldg_latitude,
			&bldg_longitude, &buildings_add_date, &buildings_chgdate)
		// check for errors, terminates if error
		if err != nil {
			log.Fatal("Error scanning row: ", err.Error())
		}
		// inserts values as arguments into PSQL
		// inserting values as arguments into PSQL
		_, err = pgConn.Exec(`
	INSERT INTO elonbuildingspg (buildings_id, bldg_desc, bldg_location, bldg_location_representation, 
		bldg_type, bldg_type_representation, bldg_long_desc, bldg_city, bldg_state, bldg_zip, 
		bldg_sector, bldg_latitude, bldg_longitude, buildings_add_date, buildings_chgdate)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`,
			buildings_id, bldg_desc, bldg_location, bldg_location_representation, bldg_type,
			bldg_type_representation, bldg_long_desc, bldg_city, bldg_state, bldg_zip, bldg_sector,
			bldg_latitude, bldg_longitude, buildings_add_date, buildings_chgdate)
		if err != nil {
			log.Fatal("Error inserting into PostgreSQL table: ", err.Error())
		}
	}
}
