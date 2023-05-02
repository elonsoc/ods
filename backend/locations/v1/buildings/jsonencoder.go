package buildings_v1


import (
   "database/sql"
   "encoding/json"
   "fmt"
   "log"


   _ "github.com/lib/pq"
)




func encode() {
	// define postgres connection string
	// sql.Open opens a new database connections
   pgConnString := "postgres://postgres:odspass12345@localhost/postgres?sslmode=disable"
   pgConn, err := sql.Open("postgres", pgConnString)

   // check for connection error
   if err != nil {
       log.Fatal("Error connecting to PSQL Server:", err.Error())
   }
   defer pgConn.Close()

   // queries all rows from elonbuildingspg table 
   // elonbuildingspg is our postgres table defined in the sql migration
   rows, err := pgConn.Query("SELECT * FROM elonbuildingspg")
   if err != nil {
       log.Fatal("Error querying PostgreSQL Server:", err.Error())
   }
   defer rows.Close()

   // declare variable buildings as a slice of Building struct
   // NOTE: type Building struct is defined in sql migration file
   var buildings []Building
   for rows.Next() {
		// declare a new instance of building struct to hold row's data
       var b Building
	   // scans rows data into b instance of Building struct
       err = rows.Scan(&b.ID, &b.Desc, &b.Location,
           &b.LocationRepresentation, &b.Type, &b.TypeRepresentation,
           &b.LongDesc, &b.City, &b.State, &b.Zip, &b.Sector,
           &b.Latitude, &b.Longitude, &b.AddDate, &b.ChangeDate)
       if err != nil {
           log.Fatal("Error scanning row:", err.Error())
       }
	   // adds variable b into the buildings slice
       buildings = append(buildings, b)
   }

   // Marshall Indent converts the slice into a JSON string
   // JSON string will be indented with two spaces
   jsonData, err := json.MarshalIndent(buildings, "", "  ")
   if err != nil {
       log.Fatal("Error marshalling JSON:", err.Error())
   }


   fmt.Println(string(jsonData))
}
