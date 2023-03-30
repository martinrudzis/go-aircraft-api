package main

import (
	"github.com/gin-gonic/gin"
	"example/go-aircraft-api/controller"
)

func main() {
	router := gin.Default()
	router.GET("/icao24/:hex", controller.GetAicraftDataByIcoa24)
	router.Run("localhost:8080")
}