package main

import (
	"net/http"

	"example.com/go-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // Start the server on port 8080

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error: " + err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1 // later: get from authentication

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
