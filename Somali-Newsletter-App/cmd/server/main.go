package main

import (
	"Somali-Newsletter-App/internal/framework/db"
	"Somali-Newsletter-App/internal/interface_adopter/controller"
	"Somali-Newsletter-App/internal/interface_adopter/gateway"
	"Somali-Newsletter-App/internal/interface_adopter/routes"
	"Somali-Newsletter-App/internal/service"
	"Somali-Newsletter-App/pkg/config"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning  .env file not found. using enviroment variables")
	}
}

func main() {
	// load the database configuration from enviroment
	dbConfig := config.LoadDBConfig()

	// debug: print the loaded database configuration
	log.Printf("DB Config: Host=%s, Port=%s, User=%s, Password=%s, DBName=%s, SSLMode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode)

	// connection to the database
	database, err := db.ConnectDB(dbConfig)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	defer database.Close()

	// Create required tables if they don't exist
	if err := db.CreateTables(database); err != nil {
		log.Fatal("Error creating tables: ", err)
	}

	// Intialize the repisitory
	userRepository := gateway.NewUserRepositoryImpl(database)
	tokenRepository := gateway.NewTokenRepository(database)
	newslatterRepository := gateway.NewNewsletterRepositoryImpl(database)
	userInteractionRepository := gateway.NewUserInteractionRepositoryImpl(database)

	// Intialize the service
	userService := service.NewUserService(userRepository, tokenRepository)
	newslatterService := service.NewNewsletterService(newslatterRepository, tokenRepository)
	userInteractionService := service.NewuserInteractionService(userInteractionRepository, tokenRepository)

	// Intialize the controller
	userController := controller.NewUserController(userService)
	newslatterController := controller.NewNewsletterController(newslatterService)
	userInteractionController := controller.NewUserInteractionController(userInteractionService)

	// Intialize the gin
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://your-frontend-domain.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Regester the routes
	routes.RegisterUserRoutes(r, userController, tokenRepository)
	routes.RegisterNewslatterRoutes(r, newslatterController, tokenRepository)
	routes.RegisterUserInteractionRoutes(r, userInteractionController, tokenRepository)

	// Run the server
	if err := r.Run(":8000"); err != nil {
		log.Fatal("Error running the server: ", err)
	}
}
