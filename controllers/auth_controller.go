package controllers

import (
	"final-project-rest-api/models"
	"final-project-rest-api/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type ChangePasswordInput struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
}

// Login handles user login
// @Summary Login as a user.
// @Description Logging in to get JWT token to access admin or user API by roles.
// @Tags Auth
// @Param Body body LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := models.LoginCheck(input.Username, input.Password, db)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username or password is incorrect"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})
}

// Register handles user registration
// @Summary Register a user.
// @Description Registering a user from public access.
// @Tags Auth
// @Param Body body RegisterInput true "the body to register a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	_, err := u.SaveUser(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

// ChangePassword handles changing user's password
// @Summary Change password for a user.
// @Description Changing password for a logged-in user.
// @Tags Auth
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param Body body ChangePasswordInput true "the body to change password for a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/change-password [put]
func ChangePassword(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input ChangePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var u models.User
	if err := db.First(&u, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := u.VerifyPassword(input.CurrentPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid current password"})
		return
	}

	hashedPassword, err := models.HashPassword(input.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	u.Password = hashedPassword

	if err := db.Save(&u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
