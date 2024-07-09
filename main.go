package main

import (
	"net/http"
	"strconv"

	"farlaap99/rest-api/db"
	"farlaap99/rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE
	server.GET("/event/:id", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080

}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id."})
		return
	}

	event, err := models.GetEventByID(int16(eventId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"massege": "Could not fetch events. try again later."})
		return
	}

	context.JSON(http.StatusOK, events)

}
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"massage": "could not parse request data!"})
		return
	}

	event.ID = 1
	event.UserID = 2

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"massege": "Could not create events. try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"massage": "Event created!", "event": event})

}
