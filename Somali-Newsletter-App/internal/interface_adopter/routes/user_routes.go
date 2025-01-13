package routes

import (
	"Somali-Newsletter-App/internal/interface_adopter/controller"
	"Somali-Newsletter-App/internal/repository"
	"Somali-Newsletter-App/pkg/middleware"

	"github.com/gin-gonic/gin"
)



func RegisterUserRoutes(router *gin.Engine, userController *controller.UserController, tokenRepo repository.TokenRepository)  {
	// apply middleware
	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	// user routes
	userRoutes := router.Group("/users")
	{
		// puplisher routes
		userRoutes.POST("", userController.RegesterUser)
		userRoutes.POST("/authenticate", userController.AuthenticateUser)
		
		// protected routes
		userRoutes.Use(authMiddleware)
		{
			userRoutes.GET("", userController.GetAllUser)
			userRoutes.GET("/:id", userController.GetUserByID)
			userRoutes.PUT("/:id", userController.UpdateUser)
			userRoutes.DELETE("/:id", userController.DeleteUser)
		}
	}


}