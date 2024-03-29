package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ichbinnichts/events-api/models"
)

func getHome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Everythings fine"})
}

func getEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not convert event id.", "error": err})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not get event.", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"event": event})
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

	event.UserId = context.GetInt64("userId")

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event.", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse id.", "error": err})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event.", "error": err})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not bind event."})
		return
	}

	updatedEvent.ID = id

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse id."})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not get event."})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted."})
}
