package models

type Laptop struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	ReleaseYear int
	Spec        string
	Price       float64
	BrandID     uint
	Brand       Brand
	CategoryID  uint
	Category    Category
}
