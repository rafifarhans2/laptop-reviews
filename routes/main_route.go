package routes

import (
	"final-project-rest-api/controllers"
	"final-project-rest-api/middleware"

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
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(corsConfig))

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// User routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes with JWT middleware
	auth := r.Group("/auth")
	auth.Use(middleware.JwtAuthMiddleware())
	{
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

		// Brand
		api.GET("/brands", controllers.GetBrands)
		api.GET("/brand/:id", controllers.GetBrandByID)
		api.POST("/brand", middleware.JwtAuthMiddleware(), controllers.CreateBrand)
		api.PUT("/brand/:id", middleware.JwtAuthMiddleware(), controllers.UpdateBrand)
		api.DELETE("/brand/:id", middleware.JwtAuthMiddleware(), controllers.DeleteBrand)

		// Laptop
		api.GET("/laptops", controllers.GetLaptops)
		api.GET("/laptop/:id", controllers.GetLaptopById)
		api.POST("/laptop", middleware.JwtAuthMiddleware(), controllers.CreateLaptop)
		api.PUT("/laptop/:id", middleware.JwtAuthMiddleware(), controllers.UpdateLaptop)
		api.DELETE("/laptop/:id", middleware.JwtAuthMiddleware(), controllers.DeleteLaptop)

		// Profile
		api.GET("/profiles", controllers.GetProfile)
		api.POST("/profile", middleware.JwtAuthMiddleware(), controllers.CreateProfile)
		api.PUT("/profile/:id", middleware.JwtAuthMiddleware(), controllers.UpdateProfile)

		// Comment
		api.GET("/comments", controllers.GetComments)
		api.GET("/comment/:id", controllers.GetCommentById)
		api.POST("/comment", middleware.JwtAuthMiddleware(), controllers.CreateComment)
		api.PUT("/comment/:id", middleware.JwtAuthMiddleware(), controllers.UpdateComment)
		api.DELETE("/comment/:id", middleware.JwtAuthMiddleware(), controllers.DeleteComment)
	}

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
