package routes

import (
	"Somali-Newsletter-App/internal/interface_adopter/controller"
	"Somali-Newsletter-App/internal/repository"
	"Somali-Newsletter-App/pkg/middleware"

	"github.com/gin-gonic/gin"
)



func RegisterNewslatterRoutes(router *gin.Engine, newslatterController *controller.NewsletterController, tokenRepo repository.TokenRepository ) {
	// apply  middleware 
 authMiddleware :=  middleware.AuthMiddleware(tokenRepo)

 // newslatter routes
  newslatterRoutes := router.Group("/newslatter")
	{

		// protected 
		newslatterRoutes.Use(authMiddleware)
		{
			newslatterRoutes.POST("", newslatterController.CreateNewslatter)
			newslatterRoutes.GET("/", newslatterController.GetAllNewsletter)
			newslatterRoutes.GET("/:id", newslatterController.GetNewsletterByID)
			newslatterRoutes.PUT("/:id", newslatterController.UpdateNewslatter)
			newslatterRoutes.DELETE("/:id", newslatterController.DeleteNewsletter)

		}
	}

}