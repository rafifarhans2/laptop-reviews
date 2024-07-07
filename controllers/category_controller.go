package controllers

import (
	"final-project-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryInput struct {
	CategoryName string `json:"name" binding:"required"`
}

// CreateCategory godoc
// @Summary Create a new category.
// @Description Create a new category.
// @Tags Category
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param Body body CategoryInput true "the body to create a category"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/category [post]
func CreateCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input CategoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{
		CategoryName: input.CategoryName,
	}

	if err := db.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully", "category": category})
}

// GetCategories godoc
// @Summary Get all categories.
// @Description Get a list of all categories.
// @Tags Category
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/categories [get]
func GetCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var categories []models.Category

	if err := db.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// GetCategoryById godoc
// @Summary Get a category.
// @Description Get a category by ID.
// @Tags Category
// @Param id path string true "Category ID"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/category/{id} [get]
func GetCategoryById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category})
}

// UpdateCategory godoc
// @Summary Update a category.
// @Description Update a category by ID.
// @Tags Category
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Category ID"
// @Param Body body CategoryInput true "the body to update a category"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/category/{id} [put]
func UpdateCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	category.CategoryName = input.CategoryName

	if err := db.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully", "category": category})
}

// DeleteCategory godoc
// @Summary Delete a category.
// @Description Delete a category by ID.
// @Tags Category
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Category ID"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/category/{id} [delete]
func DeleteCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := db.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
