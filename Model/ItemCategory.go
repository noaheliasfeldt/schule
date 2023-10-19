package Model

import (
	"gorm.io/gorm"
	"time"
)

type ItemCategory struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Name           string         `gorm:"index"`
	ItemCategoryID int            `gorm:"index"`
}
