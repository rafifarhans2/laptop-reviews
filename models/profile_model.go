package models

import (
	"time"
)

type Profile struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id" gorm:"not null;unique"`
	Fullname  string    `json:"fullname" gorm:"size:255"`
	Bio       string    `json:"bio" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"-" gorm:"foreignKey:UserID"`
}
