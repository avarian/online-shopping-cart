package model

import (
	"time"

	"gorm.io/gorm"
)

type OrderItem struct {
	ID          uint            `json:"id" gorm:"not null"`
	OrderID     uint            `json:"order_id" gorm:"not null"`
	ItemID      uint            `json:"item_id" gorm:"not null"`
	Name        string          `json:"name" gorm:"not null;size:255"`
	Description string          `json:"description"`
	Price       float64         `json:"price"  gorm:"not null;size:255"`
	Qty         int             `json:"qty"  gorm:"not null"`
	Total       float64         `json:"total"  gorm:"not null"`
	CreatedBy   string          `json:"created_by" gorm:"size:255;default:SYSTEM"`
	UpdatedBy   string          `json:"updated_by" gorm:"size:255;default:SYSTEM"`
	DeletedBy   *string         `json:"deleted_by" gorm:"size:255"`
	CreatedAt   *time.Time      `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt   *time.Time      `json:"updated_at" gorm:"default:current_timestamp"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at"`

	Order *Order `json:"order,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:OrderID;references:ID"`
	Item  *Item  `json:"item,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:ItemID;references:ID"`
}
