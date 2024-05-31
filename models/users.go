package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Uuid      string `json:"uuid" validate:"required" gorm:"type:varchar(255);primaryKey"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"email,required" gorm:"unique"`
	Password  string `json:"password" validate:"required,min=6" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *Users) BeforeCreate(tx *gorm.DB) (err error) {
	//_, errCD := govalidator.ValidateStruct(user)
	//if errCD != nil {
	//	err = errCD
	//	return
	//}
	//err = nil
	user.Uuid = uuid.NewString()
	return
}
