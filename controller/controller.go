package controller

import (
	"database/sql"
	"log"
	"net/http"
	"example/go-aircraft-api/data"
	"github.com/gin-gonic/gin"

	_ "modernc.org/sqlite"
)

func GetAicraftDataByIcoa24(c *gin.Context) {
	var hex string = c.Param("hex")
	var aircraft data.AircraftData
	
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

func loadDataModelFromRow(row *sql.Rows, aircraft *data.AircraftData) {
	if err := row.Scan(&aircraft.Icao24, &aircraft.Owner, &aircraft.City, &aircraft.State, 
		&aircraft.AircraftMake, &aircraft.AircraftModel, &aircraft.YearBuilt, &aircraft.AirworthyDate, 
		&aircraft.WeightClass, &aircraft.EngineCount, &aircraft.SeatCount, &aircraft.EngineCategory, 
		&aircraft.EngineMake, &aircraft.EngineModel, &aircraft.Horsepower, &aircraft.Thrust); 
		err != nil {
			log.Fatal(err)
	}
}