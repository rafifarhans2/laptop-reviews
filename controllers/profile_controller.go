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
	Bio      string `json:"bio"`
}

// CreateProfile godoc
// @Summary Create a new profile.
// @Description Create a new profile for a user.
// @Tags Profile
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param Body body ProfileInput true "the body to create a profile"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /profiles [post]
func CreateProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input ProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile := models.Profile{
		Fullname: input.Fullname,
		Bio:      input.Bio,
		UserID:   userID,
	}

	if err := db.Create(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile created successfully", "profile": profile})
}

// GetProfile godoc
// @Summary Get a profile.
// @Description Get a profile by user ID.
// @Tags Profile
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /profiles [get]
func GetProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var profile models.Profile
	if err := db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": profile})
}

// UpdateProfile godoc
// @Summary Update a profile.
// @Description Update a profile by user ID.
// @Tags Profile
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param Body body ProfileInput true "the body to update a profile"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /profiles [put]
func UpdateProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input ProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully", "profile": profile})
}
