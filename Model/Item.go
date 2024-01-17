package Model

import (
	"time"
)

type Item struct {
	ID        uint `gorm:"primaryKey"`
	UpdatedAt time.Time
	ItemName  string `gorm:"index"`
	ItemEAN   int64  `gorm:"column:item_ean;index;quoted"`
	ItemBBD   int64  `gorm:"column:item_bbd;index;quoted"`
	ItemCount uint   `gorm:"index"`
}
