package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "modernc.org/sqlite"
)

func GetAicraftDataByIcoa24(c *gin.Context) {
	var hex string = c.Param("hex")
	var aircraft AircraftData
	
	db, err := sql.Open("sqlite", "data/FaaData.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	statement, err := db.Prepare("SELECT * FROM FAA_DATA WHERE icao24 = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer statement.Close()

	row, err := statement.Query(hex)
	if err != nil || !row.Next() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aircraft not found"})
		return
	}
	defer row.Close()

	loadDataModelFromRow(row, &aircraft)

	c.JSON(http.StatusOK, aircraft)
}

func loadDataModelFromRow(row *sql.Rows, aircraft *AircraftData) {
	if err := row.Scan(&aircraft.Icao24, &aircraft.Owner, &aircraft.City, &aircraft.State, 
		&aircraft.AircraftMake, &aircraft.AircraftModel, &aircraft.YearBuilt, &aircraft.AirworthyDate, 
		&aircraft.WeightClass, &aircraft.EngineCount, &aircraft.SeatCount, &aircraft.EngineCategory, 
		&aircraft.EngineMake, &aircraft.EngineModel, &aircraft.Horsepower, &aircraft.Thrust); 
		err != nil {
			log.Fatal(err)
	}
}

func loadAircraftDataFromRow(row *sql.Rows, aircraft *AircraftData) {
	if row.Next() { 
		if err := row.Scan(&aircraft.Icao24, &aircraft.Owner, &aircraft.City, &aircraft.State, 
			&aircraft.AircraftMake, &aircraft.AircraftModel, &aircraft.YearBuilt, &aircraft.AirworthyDate, 
			&aircraft.WeightClass, &aircraft.EngineCount, &aircraft.SeatCount, &aircraft.EngineCategory, 
			&aircraft.EngineMake, &aircraft.EngineModel, &aircraft.Horsepower, &aircraft.Thrust); 
			err != nil {
				log.Fatal(err)
		}
	}
}

func GetAircraftData(input string) []byte {
	var aircraft AircraftData
	db, err := sql.Open("sqlite", "data/FaaData.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	statement, err := db.Prepare("SELECT * FROM FAA_DATA WHERE icao24 = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer statement.Close()

	row, err := statement.Query(input)
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()

	loadAircraftDataFromRow(row, &aircraft)

	jsonData, err := json.Marshal(aircraft)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	return jsonData
}