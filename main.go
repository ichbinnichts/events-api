package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ichbinnichts/events-api/db"
	"github.com/ichbinnichts/events-api/models"
)

func main() {

	db.InitDB()

	server := gin.Default()

	server.GET("/", getHome)

	server.GET("/events", getEvents)

	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getHome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Everythings fine"})
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events.", "error": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err})
		return
	}

	event.UserId = 1
	event.ID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event.", "error": err})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}
