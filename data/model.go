package data

import (
    "database/sql"
)

type AircraftData struct {
	Icao24         sql.NullString `json:"icao24"`
    Owner          sql.NullString `json:"owner"`
    City           sql.NullString `json:"city"`
    State          sql.NullString `json:"state"`
    AircraftMake   sql.NullString `json:"aircraftMake"`
    AircraftModel  sql.NullString `json:"aircraftModel"`
    YearBuilt      sql.NullString `json:"yearBuilt"`
    AirworthyDate  sql.NullString `json:"airworthyDate"`
    WeightClass    sql.NullString `json:"weightClass"`
    EngineCount    sql.NullString `json:"engineCount"`
    SeatCount      sql.NullString `json:"seatCount"`
    EngineCategory sql.NullString `json:"engineCategory"`
    EngineMake     sql.NullString `json:"engineMake"`
    EngineModel    sql.NullString `json:"engineModel"`
    Horsepower     sql.NullString `json:"horsepower"`
    Thrust         sql.NullString `json:"thrust"`
}