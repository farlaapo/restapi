package controller

import (
	"Somali-Newsletter-App/internal/entity"
	"Somali-Newsletter-App/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

//  UserController represent a user controller
type UserController struct {
	userService service.UserService
}
 
// NewUserController return a new instance of UserController
func NewUserController (userService service.UserService) *UserController {
	return &UserController{ userService: userService}
}
// register user func
func (uc *UserController) RegesterUser(c *gin.Context) {
	var user entity.User
	// bind
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// call service
	 registeruser, err := uc.userService.RegesterUser(user.Name, user.Email, user.Password, user.RoleID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return the 
	c.JSON(201, registeruser)

	
}

// authenticateUser func
func (uc *UserController) AuthenticateUser(c *gin.Context) {
	var user entity.User
// bind
if err := c.BindJSON(&user); err != nil {
	c.JSON(400, gin.H{"error": err.Error()})
	return
}

// call service
AuthenticateUser, err := uc.userService.AuthenticateUser(user.Email, user.Password)
if err != nil {
	c.JSON(500, gin.H{"error": err.Error()})
	return
}

// return
c.JSON(200, AuthenticateUser)
}

// get all user func
func (uc *UserController) GetAllUser(c *gin.Context) {
	user, err := uc.userService.GetAllUser()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return
	c.JSON(200, user)

}

// get user by id func
func (uc *UserController) GetUserByID(c *gin.Context) {
	userParam := c.Param("id")
	userID, err := uuid.FromString(userParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// call service
	user, err := uc.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return
	c.JSON(200, user)

	
}

// update user func
func (uc *UserController) UpdateUser(c *gin.Context) {	
	var user entity.User

	// param
	userParam := c.Param("id")
	userID, err := uuid.FromString(userParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// bind
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.ID = userID

	// call service
	if  err := uc.userService.UpdateUser(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return the user
	c.JSON(200, user)


}

// delete user func
func (uc *UserController) DeleteUser(c *gin.Context) {
	// param
	userParam := c.Param("id")
	userID, err := uuid.FromString(userParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// call service
	if err := uc.userService.DeleteUser(userID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return the user
	c.JSON(200, gin.H{"message": "user deleted successfully"})
}



