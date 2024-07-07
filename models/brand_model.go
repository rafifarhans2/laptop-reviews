package models

type Brand struct {
	ID        uint   `gorm:"primaryKey"`
	BrandName string `gorm:"size:255"`
	Laptops   []Laptop
}
