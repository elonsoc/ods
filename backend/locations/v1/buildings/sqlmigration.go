package buildings_v1

// import MSSQL and PSQL drivers and libraries that allow us to connect
import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type mssqlBuilding struct {
	ID                     string        `json:"id" db:"BUILDINGS_ID"`
	Description            string        `json:"description" db:"BLDG_DESC"`
	Location               string        `json:"location" db:"BLDG_LOCATION"`
	LocationRepresentation string        `json:"location_representation" db:"BLDG_LOCATION_REP"`
	Type                   string        `json:"type" db:"BLDG_TYPE"`
	TypeRepresentation     string        `json:"type_representation" db:"BLDG_TYPE_REP"`
	LongDesc               string        `json:"long_description" db:"BLDG_LONG_DESC"`
	City                   string        `json:"city" db:"BLDG_CITY"`
	State                  string        `json:"state" db:"BLDG_STATE"`
	Zip                    string        `json:"zip" db:"BLDG_ZIP"`
	Sector                 string        `json:"sector" db:"BLDG_SECTOR"`
	Latitude               float64       `json:"latitude" db:"BLDG_LATITUDE"`
	Longitude              float64       `json:"longiutde" db:"BLDG_LONGITUDE"`
	AddDate                time.Duration `json:"add_date" db:"BLDG_ADD_DATE"`
	ChangeDate             time.Duration `json:"change_date" db:"BLDG_CHANGE_DATE"`
}

// TODO(@ronydahdal) finish :)
func migrate(sqlConn *sql.DB, pgConn *pgx.Conn) {
	// // connecting to MSSQL Server Database (Elon's in the future)
	// // MSSQL connection string w/ server address, user, pass, database name
	// mssqlConnString := "server=localhost;user id = sa; password=<odspass12345>;database=[databasename]"
	// // mssqlconn holds the connection to the mssql server
	// // mssqlconn is used for reference to exectute commands on database
	// mssqlConn, err := sql.Open("sqlserver", mssqlConnString)
	// // sql.Open returns an error value nil if connection success
	// // if connections fails, err is set to non-nil
	// if err != nil {
	// 	log.Fatal("Error connecting to MSSQL Server:", err.Error())
	// 	return
	// }
	// // defer closing the connection of database until end of migration/function
	// defer mssqlConn.Close()

	// // migrating data from MSSQL to PSQL using a SQL query to MSSQL database
	// // query selects all columns from data table
	mssqlQuery := "SELECT * FROM elonbuildings"
	// mssqlRows will hold all rows from data table retrieved from query
	mssqlRows, err := sqlConn.Query(mssqlQuery)
	// if err != nil {
	// 	log.Fatal("Error querying MSSQL Server:", err.Error())
	// 	return
	// }
	// defer mssqlRows.Close()

	// // define connection string for PSQL database
	// // connecting to a localhost psql database here
	// // postgress (prior to :odspass...) is the username
	// // odspass12345 is the placeholder pass
	// // postgres (prior to ?sslmode..) is placeholder database name

	// pgConnString := "postgres://postgres:odspass12345@localhost/postgres?sslmode=disable"
	// // connecting to PostgreSQL server using Go driver from package
	// pgConn, err := sql.Open("postgres", pgConnString)
	// // checking if connection is unsuccessful
	// if err != nil {
	// 	log.Fatal("Error connecting to PSQL Server", err.Error())
	// }
	// defer pgConn.Close()

	// inserting data into PSQL database
	// for loop to iterate over rows of MSSQL

	for mssqlRows.Next() {
		// re-use building struct
		var building mssqlBuilding
		// scan method retrieves values from MSSQL rows & assings to Go's memory
		err = mssqlRows.Scan(&building)
		// check for errors, terminates if error
		if err != nil {
			log.Fatal("Error scanning row: ", err.Error())
		}

		// inserting values as arguments into PSQL
		// note: elonbuildingspg is placeholder/assumed postgresql database tablename
		_, err = pgConn.Exec(context.Background(), `
	INSERT INTO elonbuildingspg (buildings_id, bldg_desc, bldg_location, bldg_location_representation,
		bldg_type, bldg_type_representation, bldg_long_desc, bldg_city, bldg_state, bldg_zip,
		bldg_sector, bldg_latitude, bldg_longitude, buildings_add_date, buildings_chgdate)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`,
			building.ID, building.Description, building.Location, building.LocationRepresentation, building.Type,
			building.TypeRepresentation, building.LongDesc, building.City, building.State, building.Zip, building.Sector,
			building.Latitude, building.Longitude, building.AddDate, building.ChangeDate)
		if err != nil {
			log.Fatal("Error inserting into PostgreSQL table: ", err.Error())
		}
	}
}
