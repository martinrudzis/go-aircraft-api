package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type AircraftData struct {
	icao24, owner, city, state, aircraftMake, aircraftModel, yearBuilt, airworthyDate, weightClass, 
	engineCount, seatCount, engineCategory, engineMake, engineModel, horsepower, thrust string
}


func printAircraftData(aircraft *AircraftData) {
	fmt.Printf("icao24: %s\n" + "owner: %s\n" + "city: %s\n" + "state: %s\n" + "aircraft make: %s\n" + 
	"aircraft model: %s\n" + "year built: %s\n" + "airworthy date: %s\n" + "weight class: %s\n" + 
	"engine count: %s\n" + "seat count: %s\n" + "engine category: %s\n" + "engine make: %s\n" + 
	"engine model: %s\n" + "horsepower: %s\n" + "thrust: %s\n", 
	aircraft.icao24, aircraft.owner, aircraft.city, aircraft.state, aircraft.aircraftMake, 
	aircraft.aircraftModel, aircraft.yearBuilt, aircraft.airworthyDate, aircraft.weightClass, 
	aircraft.engineCount, aircraft.seatCount, aircraft.engineCategory,aircraft.engineMake, 
	aircraft.engineModel, aircraft.horsepower, aircraft.thrust)
}

func loadAircraftDataFromRow(row *sql.Rows, aircraft *AircraftData) {
	if row.Next() { 
		if err := row.Scan(&aircraft.icao24, &aircraft.owner, &aircraft.city, &aircraft.state, 
			&aircraft.aircraftMake, &aircraft.aircraftModel, &aircraft.yearBuilt, &aircraft.airworthyDate, 
			&aircraft.weightClass, &aircraft.engineCount, &aircraft.seatCount, &aircraft.engineCategory, 
			&aircraft.engineMake, &aircraft.engineModel, &aircraft.horsepower, &aircraft.thrust); 
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

	printAircraftData(&aircraft)
}