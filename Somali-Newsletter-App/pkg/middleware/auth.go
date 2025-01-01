package middleware

import (
	"Somali-Newsletter-App/internal/repository"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenRepo repository.TokenRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from authorization header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			log.Println("Missing Authorization header")
			c.JSON(401, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}
		// the token is usauly in the format "Bearer <token>"
		parts := strings.Split(authHeader, "")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Println("Invalid Authorization format ")
			c.JSON(401, gin.H{"error": "Invalid Authorization format"})
			c.Abort()
			return
		}
		tokenString := parts[1]

		//Fetch the token from the database
		token, err := tokenRepo.FindToken(tokenString)
		if err != nil {
			log.Printf("Token lookup failed: %v", err)
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check if the token expired
		if token.ExpiredAt.Before(time.Now()) {
			log.Printf("Token expired at: %v", token.ExpiredAt)
			c.JSON(401, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}
		// set the user ID in the contex
		c.Set("userID", token.UserID)

		// Proceed with the request
		c.Next()
	}
}
