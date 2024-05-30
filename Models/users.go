package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Uuid      uuid.UUID `json:"uuid" validate:"required" gorm:"type:char(36);primaryKey"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,unique,email" gorm:"unique"`
	Password  string    `json:"password" validate:"required,min=6" gorm:"not null"`
	Photos    []Photos  `gorm:"foreignKey:UserID;references:Uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *Users) BeforeCreate(tx *gorm.DB) (err error) {
	user.Uuid = uuid.New()
	return
}
