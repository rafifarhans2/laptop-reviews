package controllers

import (
	"final-project-rest-api/models"
	"final-project-rest-api/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentInput struct {
	Content  string `json:"description" binding:"required"`
	Rating   int    `json:"rating" binding:"required"`
	LaptopID uint   `json:"laptop_id" binding:"required"`
}

// CreateComment godoc
// @Summary Create a new comment.
// @Description Create a new comment for a laptop.
// @Tags Comment
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param Body body CommentInput true "the body to create a comment"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /comment [post]
func CreateComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input CommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := models.Comment{
		Content:  input.Content,
		Rating:   input.Rating,
		UserID:   userID,
		LaptopID: input.LaptopID,
	}

	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment created successfully", "comment": comment})
}

// GetComments godoc
// @Summary Get all comments.
// @Description Get a list of all comments.
// @Tags Comment
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/comments [get]
func GetComments(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var comments []models.Comment

	if err := db.Preload("User").Preload("Laptop").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

// GetComment godoc
// @Summary Get a comment.
// @Description Get a comment by ID.
// @Tags Comment
// @Param id path string true "Comment ID"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/comment/{id} [get]
func GetCommentById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var comment models.Comment
	if err := db.Preload("User").Preload("Laptop").Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

// UpdateComment godoc
// @Summary Update a comment.
// @Description Update a comment by ID.
// @Tags Comment
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Comment ID"
// @Param Body body CommentInput true "the body to update a comment"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/comment/{id} [put]
func UpdateComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input CommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var comment models.Comment
	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if comment.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	comment.Content = input.Content
	comment.Rating = input.Rating
	comment.LaptopID = input.LaptopID

	if err := db.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully", "comment": comment})
}

// DeleteComment godoc
// @Summary Delete a comment.
// @Description Delete a comment by ID.
// @Tags Comment
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Comment ID"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/comment/{id} [delete]
func DeleteComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var comment models.Comment
	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if comment.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := db.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
