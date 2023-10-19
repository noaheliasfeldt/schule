package Model

import (
	"gorm.io/gorm"
	"time"
)

type Item struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ItemName  string         `gorm:"index"`
	ItemBBD   time.Time      `gorm:"index"` // englisch für mhd
}
