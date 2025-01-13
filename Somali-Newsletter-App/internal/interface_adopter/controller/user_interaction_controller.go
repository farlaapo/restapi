package controller

import (
	"Somali-Newsletter-App/internal/entity"
	"Somali-Newsletter-App/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// Controller is a struct that holds the service instance
type UserInteractionController struct {
	userInteractionService service.UserInteractionService
}

// NewUserInteractionController returns a new instance of UserInteractionController
func NewUserInteractionController(userInteractionService  service.UserInteractionService) *UserInteractionController {
	return &UserInteractionController{
		userInteractionService: userInteractionService,
	}
}

// CreatUserInteractions returns a list of user interactions
func (US *UserInteractionController) CreateUserInteraction(c *gin.Context) {
	var userInteraction entity.UserInteraction
	// bind the request body to the userInteraction struct
	if err := c.BindJSON(&userInteraction); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// call the service to create a new user interaction
	createdUserInteraction, err := US.userInteractionService.CreateUserInteraction(userInteraction.NewsletterID, userInteraction.UserID, userInteraction.Liked, userInteraction.Read, userInteraction.Comment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// return
	c.JSON(201, createdUserInteraction)
}

 // GetUserInteractionsByID returns a list of user interactions
func (US *UserInteractionController) GetUserInteractionByID(c *gin.Context) {
	// param 
	userInteractionParam := c.Param("id")
	userInteractionID, err := uuid.FromString(userInteractionParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// call the service to get by id
	userInteraction, err := US.userInteractionService.GetUserInteractionByID(userInteractionID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return
	c.JSON(200, userInteraction)
}

// GetUserInteractions returns a list of user interactions
func (US *UserInteractionController) GetAllUserInteraction(c *gin.Context) {
	// call the service to get all
	userInteractions, err := US.userInteractionService.GetAllUserInteraction()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// return
	c.JSON(200, userInteractions)
}

// UpdateUserInteractions returns a list of user interactions
func (US *UserInteractionController) UpdateUserInteraction(c *gin.Context) {
	var userInteraction entity.UserInteraction
	// param
	userInteractionParam := c.Param("id")
	userInteractionID, err := uuid.FromString(userInteractionParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// bind
	if err := c.BindJSON(&userInteraction); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} 

	userInteraction.ID = userInteractionID
	// call the service to update
	if err := US.userInteractionService.UpdateUserInteraction(&userInteraction); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return
	c.JSON(200, gin.H{"message": " User interaction updated successfully"})
	
}

func (US *UserInteractionController) DeleteUserInteraction(c *gin.Context) {
	// param 
	userInteractionParam := c.Param("id")
	userInteractionID, err := uuid.FromString(userInteractionParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// call the service to delete
	if err := US.userInteractionService.DeleteUserInteraction(userInteractionID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return
	c.JSON(200, gin.H{"message": "User interaction deleted"})
}







