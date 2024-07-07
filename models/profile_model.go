package models

type Profile struct {
	ID       uint   `gorm:"primaryKey"`
	Fullname string `gorm:"size:255"`
	Bio      string
	UserID   uint `gorm:"unique"`
	User     User
}
