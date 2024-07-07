package models

import "time"

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	Content   string
	Rating    int
	UserID    uint
	User      User
	LaptopID  uint
	Laptop    Laptop
	CreatedAt time.Time
	UpdatedAt time.Time
}
