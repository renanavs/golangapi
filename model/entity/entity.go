package entity

import (
	"time"

	"gorm.io/gorm"
)

type BasicEntity struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
