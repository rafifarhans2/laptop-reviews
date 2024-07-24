package controllers

import (
	"final-project-rest-api/models"
	"final-project-rest-api/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProfileInput struct {
	Fullname string `json:"fullname" binding:"required"`
	Bio      string `json:"bio" binding:"required"`
}

// CreateProfile godoc
// @Summary Create a new profile.
// @Description Create a new profile for a user.
// @Tags Profile
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param Body body ProfileInput true "the body to create a profile"
// @Produce json
// @Success 201 {object} models.Profile
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 409 {object} map[string]interface{} "Profile already exists for this user"
// @Failure 500 {object} map[string]interface{} "Failed to create profile"
// @Router /api/profile [post]
func CreateProfile(c *gin.Context) {
	var input ProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	var existingProfile models.Profile
	if err := db.Where("user_id = ?", userID).First(&existingProfile).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Profile already exists for this user"})
		return
	} else if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	profile := models.Profile{
		UserID:   userID,
		Fullname: input.Fullname,
		Bio:      input.Bio,
	}

	if err := db.Create(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	c.JSON(http.StatusCreated, profile)
}

// GetProfile godoc
// @Summary Get all profiles.
// @Description Retrieve all profiles.
// @Tags Profile
// @Produce json
// @Success 200 {array} models.Profile
// @Failure 500 {object} map[string]interface{} "Failed to retrieve profiles"
// @Router /api/profiles [get]
func GetProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var profiles []models.Profile
	if err := db.Find(&profiles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve profiles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profiles": profiles})
}

// UpdateProfile godoc
// @Summary Update a profile.
// @Description Update a profile by user ID.
// @Tags Profile
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param Body body ProfileInput true "the body to update a profile"
// @Produce json
// @Success 200 {object} models.Profile
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Profile not found"
// @Failure 500 {object} map[string]interface{} "Failed to update profile"
// @Router /api/profile [put]
func UpdateProfile(c *gin.Context) {
	var input ProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var profile models.Profile
	if err := db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	profile.Fullname = input.Fullname
	profile.Bio = input.Bio

	if err := db.Save(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, profile)
}
