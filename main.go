package main

import (
	//"net/http"
	"github.com/gin-gonic/gin"
	"example/go-aircraft-api/data"
)

func main() {
	// data.GetAircraftData("AA5B92")

	router := gin.Default()
	router.GET("/icao24/:hex", data.GetAicraftDataByIcoa24)
	router.Run("localhost:8080")
}