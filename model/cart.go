package model

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint            `json:"id" gorm:"not null"`
	AccountID uint            `json:"account_id" gorm:"uniqueIndex:idx_account_item_key;not null"`
	ItemID    uint            `json:"item_id" gorm:"uniqueIndex:idx_account_item_key;not null"`
	Qty       int             `json:"qty"  gorm:"not null"`
	CreatedBy string          `json:"created_by" gorm:"size:255;default:SYSTEM"`
	UpdatedBy string          `json:"updated_by" gorm:"size:255;default:SYSTEM"`
	DeletedBy *string         `json:"deleted_by" gorm:"size:255"`
	CreatedAt *time.Time      `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt *time.Time      `json:"updated_at" gorm:"default:current_timestamp"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`

	Account *Item `json:"account,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:AccountID;references:ID"`
	Item    *Item `json:"item,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:ItemID;references:ID"`
}
