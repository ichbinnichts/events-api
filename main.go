package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ichbinnichts/events-api/db"
	"github.com/ichbinnichts/events-api/routes"
)

func main() {

	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
