package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", getHome)

	server.GET("/events", getEvents)

	server.GET("/events/:id", getEventById)

	server.POST("/events", createEvent)
}