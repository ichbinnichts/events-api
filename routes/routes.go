package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ichbinnichts/events-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	//Creates a group where this group will use the authenticate middleware

	authenticated := server.Group("/")

	//Adds the middle ware to the group
	authenticated.Use(middlewares.Authenticate)

	//Events routes
	server.GET("/", getHome)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	//User routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
