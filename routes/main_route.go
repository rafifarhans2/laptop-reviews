package routes

import (
	"final-project-rest-api/controllers"
	"final-project-rest-api/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// CORS configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"} // Ganti dengan origin yang sesuai
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.MaxAge = 12 * 60 * 60 // 12 jam

	r.Use(cors.New(corsConfig))
	r.OPTIONS("/*path", handleOptions)

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// User routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/auth")
	auth.Use(middleware.JwtAuthMiddleware())
	{
		auth.OPTIONS("/*path", handleOptions)
		auth.PUT("/change-password", controllers.ChangePassword)
	}

	api := r.Group("/api")
	{
		// Category
		api.GET("/categories", controllers.GetCategories)
		api.GET("/category/:id", controllers.GetCategoryById)
		api.POST("/category", middleware.JwtAuthMiddleware(), controllers.CreateCategory)
		api.PUT("/category/:id", middleware.JwtAuthMiddleware(), controllers.UpdateCategory)
		api.DELETE("/category/:id", middleware.JwtAuthMiddleware(), controllers.DeleteCategory)
		api.OPTIONS("/category", handleOptions)
		api.OPTIONS("/category/*path", handleOptions)

		// Brand
		api.GET("/brands", controllers.GetBrands)
		api.GET("/brand/:id", controllers.GetBrandByID)
		api.POST("/brand", middleware.JwtAuthMiddleware(), controllers.CreateBrand)
		api.PUT("/brand/:id", middleware.JwtAuthMiddleware(), controllers.UpdateBrand)
		api.DELETE("/brand/:id", middleware.JwtAuthMiddleware(), controllers.DeleteBrand)
		api.OPTIONS("/brand", handleOptions)
		api.OPTIONS("/brand/*path", handleOptions)

		// Laptop
		api.GET("/laptops", controllers.GetLaptops)
		api.GET("/laptop/:id", controllers.GetLaptopById)
		api.POST("/laptop", middleware.JwtAuthMiddleware(), controllers.CreateLaptop)
		api.PUT("/laptop/:id", middleware.JwtAuthMiddleware(), controllers.UpdateLaptop)
		api.DELETE("/laptop/:id", middleware.JwtAuthMiddleware(), controllers.DeleteLaptop)
		api.OPTIONS("/laptop", handleOptions)
		api.OPTIONS("/laptop/*path", handleOptions)

		// Profile
		api.GET("/profiles", controllers.GetProfile)
		api.POST("/profile", middleware.JwtAuthMiddleware(), controllers.CreateProfile)
		api.PUT("/profile", middleware.JwtAuthMiddleware(), controllers.UpdateProfile)
		api.OPTIONS("/profile", handleOptions)
		api.OPTIONS("/profiles", handleOptions)

		// Comment
		api.GET("/comments", controllers.GetComments)
		api.GET("/comment/:id", controllers.GetCommentById)
		api.POST("/comment", middleware.JwtAuthMiddleware(), controllers.CreateComment)
		api.PUT("/comment/:id", middleware.JwtAuthMiddleware(), controllers.UpdateComment)
		api.DELETE("/comment/:id", middleware.JwtAuthMiddleware(), controllers.DeleteComment)
		api.OPTIONS("/comment", handleOptions)
		api.OPTIONS("/comment/*path", handleOptions)
	}

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func handleOptions(c *gin.Context) {
	c.Status(http.StatusOK)
}
