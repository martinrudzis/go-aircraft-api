package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type AircraftData struct {
	Icao24         string `json:"icao24"`
    Owner          string `json:"owner"`
    City           string `json:"city"`
    State          string `json:"state"`
    AircraftMake   string `json:"aircraftMake"`
    AircraftModel  string `json:"aircraftModel"`
    YearBuilt      string `json:"yearBuilt"`
    AirworthyDate  string `json:"airworthyDate"`
    WeightClass    string `json:"weightClass"`
    EngineCount    string `json:"engineCount"`
    SeatCount      string `json:"seatCount"`
    EngineCategory string `json:"engineCategory"`
    EngineMake     string `json:"engineMake"`
    EngineModel    string `json:"engineModel"`
    Horsepower     string `json:"horsepower"`
    Thrust         string `json:"thrust"`
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

func GetAircraftData(input string) {
	var aircraft AircraftData
	db, err := sql.Open("sqlite", "data/FaaData.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	statement, err := db.Prepare("SELECT * FROM FAA_DATA WHERE icao24 = ?")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer statement.Close()

	row, err := statement.Query(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer row.Close()

	loadAircraftDataFromRow(row, &aircraft)

	jsonData, err := json.Marshal(aircraft)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}