package models

type Category struct {
	ID           uint   `gorm:"primaryKey"`
	CategoryName string `gorm:"not null"`
	Laptops      []Laptop
}
