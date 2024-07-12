package routes

import (
	"event-booking-app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Endpoints/routes and their corresponding Handlers
	server.GET("/events", getEvents)         // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:eventId", getEvent) // Dynamic path parameter (with ':')

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:eventId", updateEvent)
	authenticated.DELETE("/events/:eventId", deleteEvent)
	authenticated.POST("/events/:eventId/register", registerForEvent)
	authenticated.DELETE("/events/:eventId/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)

}
