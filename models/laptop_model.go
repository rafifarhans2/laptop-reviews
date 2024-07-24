package models

import (
	"time"

	"gorm.io/gorm"
)

type Laptop struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	BrandID     uint           `gorm:"not null" json:"brand_id"`
	CategoryID  uint           `gorm:"not null" json:"category_id"`
	Name        string         `gorm:"not null" json:"name"`
	ReleaseYear int            `json:"release_year"`
	Spec        string         `json:"spec"`
	Price       float64        `json:"price"`
	Comments    []Comment      `gorm:"foreignKey:LaptopID" json:"comments"`
	Brand       Brand          `gorm:"foreignKey:BrandID" json:"brand"`
	Category    Category       `gorm:"foreignKey:CategoryID" json:"category"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
