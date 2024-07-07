package main

import (
	"final-project-rest-api/configs"
	"final-project-rest-api/docs"
	"final-project-rest-api/routes"
	"final-project-rest-api/utils"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	docs.SwaggerInfo.Title = "Laptop Review REST API"
	docs.SwaggerInfo.Description = "This is REST API Laptop Review."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("HOST", "localhost:8080")
	if environment == "development" {
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	db := configs.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run()
}
