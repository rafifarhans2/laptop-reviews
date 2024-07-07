package controllers

import (
	"final-project-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BrandInput struct {
	BrandName string `json:"name" binding:"required"`
}

// CreateBrand godoc
// @Summary Create a new Brand.
// @Description Create a new Brand.
// @Tags Brand
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param Body body BrandInput true "the body to create a brand"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/brand [post]
func CreateBrand(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input BrandInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	brand := models.Brand{
		BrandName: input.BrandName,
	}

	if err := db.Create(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Brand created successfully", "category": brand})
}

// GetBrands godoc
// @Summary Get all brands.
// @Description Get a list of all brand.
// @Tags Brand
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/brands [get]
func GetBrands(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var brands []models.Brand

	if err := db.Find(&brands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve brands"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"brands": brands})
}

// GetBrandById godoc
// @Summary Get a brand.
// @Description Get a brand by ID.
// @Tags Brand
// @Param id path string true "Brand ID"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/brand/{id} [get]
func GetBrandByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var brand models.Brand
	if err := db.Where("id = ?", c.Param("id")).First(&brand).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": brand})
}

// UpdateBrand godoc
// @Summary Update a brand.
// @Description Update a brand by ID.
// @Tags Brand
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Brand ID"
// @Param Body body BrandInput true "the body to update a brand"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/brand/{id} [put]
func UpdateBrand(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input BrandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var brand models.Brand
	if err := db.Where("id = ?", c.Param("id")).First(&brand).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	brand.BrandName = input.BrandName

	if err := db.Save(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update brand"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Brand updated successfully", "brand": brand})
}

// DeleteBrand godoc
// @Summary Delete a brand.
// @Description Delete a brand by ID.
// @Tags Brand
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Brand ID"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/brand/{id} [delete]
func DeleteBrand(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var brand models.Brand
	if err := db.Where("id = ?", c.Param("id")).First(&brand).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	if err := db.Delete(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete brand"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Brand deleted successfully"})
}
