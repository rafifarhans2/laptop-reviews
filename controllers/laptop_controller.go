package controllers

import (
	"final-project-rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LaptopInput struct {
	Name        string  `json:"name" binding:"required"`
	ReleaseYear int     `json:"release_year"`
	Spec        string  `json:"spec"`
	Price       float64 `json:"price"`
	BrandID     uint    `json:"brand_id" binding:"required"`
	CategoryID  uint    `json:"category_id" binding:"required"`
}

// CreateLaptop godoc
// @Summary Create a new laptop.
// @Description Create a new laptop.
// @Tags Laptop
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param Body body LaptopInput true "the body to create a laptop"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/laptop [post]
func CreateLaptop(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LaptopInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	laptop := models.Laptop{
		Name:        input.Name,
		ReleaseYear: input.ReleaseYear,
		Spec:        input.Spec,
		Price:       input.Price,
		BrandID:     input.BrandID,
		CategoryID:  input.CategoryID,
	}

	if err := db.Create(&laptop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create laptop"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Laptop created successfully", "laptop": laptop})
}

// GetLaptops godoc
// @Summary Get all laptops.
// @Description Get a list of all laptops.
// @Tags Laptop
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/laptops [get]
func GetLaptops(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var laptops []models.Laptop

	if err := db.Preload("Brand").Preload("Category").Preload("Comments").Find(&laptops).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve laptops"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"laptops": laptops})
}

// GetLaptopById godoc
// @Summary Get a laptop.
// @Description Get a laptop by ID.
// @Tags Laptop
// @Param id path string true "Laptop ID"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/laptop/{id} [get]
func GetLaptopById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var laptop models.Laptop

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := db.Preload("Brand").Preload("Category").Preload("Comments.User").Where("id = ?", id).First(&laptop).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Laptop not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"laptop": laptop})
}

// UpdateLaptop godoc
// @Summary Update a laptop.
// @Description Update a laptop by ID.
// @Tags Laptop
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Laptop ID"
// @Param Body body LaptopInput true "the body to update a laptop"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/laptop/{id} [put]
func UpdateLaptop(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LaptopInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var laptop models.Laptop

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := db.Where("id = ?", id).First(&laptop).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Laptop not found"})
		return
	}

	laptop.Name = input.Name
	laptop.ReleaseYear = input.ReleaseYear
	laptop.Spec = input.Spec
	laptop.Price = input.Price
	laptop.BrandID = input.BrandID
	laptop.CategoryID = input.CategoryID

	if err := db.Save(&laptop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update laptop"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Laptop updated successfully", "laptop": laptop})
}

// DeleteLaptop godoc
// @Summary Delete a laptop.
// @Description Delete a laptop by ID.
// @Tags Laptop
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Laptop ID"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/laptop/{id} [delete]
func DeleteLaptop(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var laptop models.Laptop

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := db.Where("id = ?", id).First(&laptop).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Laptop not found"})
		return
	}

	if err := db.Delete(&laptop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete laptop"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Laptop deleted successfully"})
}
