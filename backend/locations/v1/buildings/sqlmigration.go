package buildings_v1

// import MSSQL and PSQL drivers and libraries that allow us to connect
import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

type Building struct {
	ID			string
	Desc			sql.NullString
	Location		sql.NullString
	LocationRepresentation	sql.NullString
	Type			sql.NullString
	TypeRepresentation	sql.NullString
	LongDesc		sql.NullString
	City			sql.NullString
	State			sql.NullString
	Zip			sql.NullString
	Sector			sql.NullString
	Latitude		sql.NullFloat64
	Longitude		sql.NullFloat64
	AddDate			sql.NullTime
	ChangeDate		sql.NullTime
}
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
	// postgress (prior to :odspass...) is the username
	// odspass12345 is the placeholder pass
	// postgres (prior to ?sslmode..) is placeholder database name
	
	pgConnString := "postgres://postgres:odspass12345@localhost/postgres?sslmode=disable"
	// connecting to PostgreSQL server using Go driver from package
	pgConn, err := sql.Open("postgres", pgConnString)
	// checking if connection is unsuccessful
	if err != nil {
		log.Fatal("Error connecting to PSQL Server", err.Error())
	}
	defer pgConn.Close()

	// inserting data into PSQL database
	// for loop to iterate over rows of MSSQL

	for mssqlRows.Next() {
		// re-use building struct 
		var building Building
		// scan method retrieves values from MSSQL rows & assings to Go's memory
		err = mssqlRows.Scan(&buildings.ID, &building.Desc, &building.Location,
			&building.LocationRepresentation, &building.Type,
			&building.TypeRepresentation, &building.LongDesc, &building.City,
			&building.State, &building.Zip, &building.Sector, &building.Latitude,
			&building.Longitude, &building.AddDate, &building.ChangeDate)
		// check for errors, terminates if error
		if err != nil {
			log.Fatal("Error scanning row: ", err.Error())
		}
		
		// inserting values as arguments into PSQL
		// note: elonbuildingspg is placeholder/assumed postgresql database tablename
		_, err = pgConn.Exec(`
	INSERT INTO elonbuildingspg (buildings_id, bldg_desc, bldg_location, bldg_location_representation, 
		bldg_type, bldg_type_representation, bldg_long_desc, bldg_city, bldg_state, bldg_zip, 
		bldg_sector, bldg_latitude, bldg_longitude, buildings_add_date, buildings_chgdate)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`,
			building.ID, building.Desc, building.Location, building.LocationRepresentation, building.Type,
			building.TypeRepresentation, building.LongDesc, building.City, building.State, building.Zip, building.Sector,
			building.Latitude, building.Longitude, building.AddDate, building.ChangeDate)
		if err != nil {
			log.Fatal("Error inserting into PostgreSQL table: ", err.Error())
		}
	}
}
