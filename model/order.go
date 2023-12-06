package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID          uint            `json:"id" gorm:"not null"`
	AccountID   uint            `json:"account_id" gorm:"not null"`
	Address     string          `json:"address"`
	PhoneNumber string          `json:"phone_number" gorm:"not null;size:255"`
	Total       float64         `json:"total"  gorm:"not null"`
	Status      string          `json:"status"  gorm:"not null;size:255;default:ORDERED"`
	CreatedBy   string          `json:"created_by" gorm:"size:255;default:SYSTEM"`
	UpdatedBy   string          `json:"updated_by" gorm:"size:255;default:SYSTEM"`
	DeletedBy   *string         `json:"deleted_by" gorm:"size:255"`
	CreatedAt   *time.Time      `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt   *time.Time      `json:"updated_at" gorm:"default:current_timestamp"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at"`

	Account      *Item         `json:"account,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:AccountID;references:ID"`
	OrderItem    *OrderItem    `json:"order_item,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:OrderID;references:ID"`
	OrderVoucher *OrderVoucher `json:"order_voucher,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:OrderID;references:ID"`
}
