package routes

import (
	"Somali-Newsletter-App/internal/interface_adopter/controller"
	"Somali-Newsletter-App/internal/repository"
	"Somali-Newsletter-App/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserInteractionRoutes(router *gin.Engine, useInteractionController *controller.UserInteractionController, tokenRepo repository.TokenRepository) {
	 // apply middleware
	 authMiddleware := middleware.AuthMiddleware(tokenRepo)

	 useInteractionRoutes := router.Group("/useInteractions")
	 {
		// protected

		useInteractionRoutes.Use(authMiddleware)
		{
			useInteractionRoutes.POST("", useInteractionController.CreateUserInteraction)
			useInteractionRoutes.GET("/", useInteractionController.GetAllUserInteraction)
			useInteractionRoutes.GET(":id", useInteractionController.GetUserInteractionByID)
			useInteractionRoutes.PUT("/:id", useInteractionController.UpdateUserInteraction)
			useInteractionRoutes.DELETE(":id", useInteractionController.DeleteUserInteraction)
		}

	 }
	 
}