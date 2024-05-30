package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Photos struct {
	Uuid      uuid.UUID `json:"uuid" gorm:"type:char(36);primaryKey"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uuid.UUID `json:"user_id" validate:"required" gorm:"type:char(36);not null"`
	User      *Users    `gorm:"foreignKey:UserID;references:Uuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (photos *Photos) BeforeCreate(tx *gorm.DB) (err error) {
	photos.Uuid = uuid.New()
	return
}
