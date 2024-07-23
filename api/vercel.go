package api

import (
	"final-project-rest-api/configs"
	"final-project-rest-api/docs"
	"final-project-rest-api/routes"
	"final-project-rest-api/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	App *gin.Engine
)

func init() {
	start := time.Now()
	App = gin.New()

	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	docs.SwaggerInfo.Title = "Laptop REST API"
	docs.SwaggerInfo.Description = "This is REST API Laptop."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("HOST", "localhost:8080")
	if environment == "development" {
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	log.Println("Connecting to database...")
	db := configs.ConnectDataBase()

	log.Println("Setting up routes...")
	App = routes.SetupRouter(db)

	log.Printf("Initialization completed in %s\n", time.Since(start))
}

// Entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	App.ServeHTTP(w, r)
}
