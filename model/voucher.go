package model

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	ID          uint            `json:"id" gorm:"not null"`
	Code        string          `json:"code" gorm:"not null;size:255"`
	Name        string          `json:"name" gorm:"not null;size:255"`
	Description string          `json:"description"`
	Percentage  float64         `json:"percentage" gorm:"not null;size:255"`
	Max         float64         `json:"price"  gorm:"not null;size:255"`
	IsActive    bool            `json:"bool"  gorm:"not null"`
	CreatedBy   string          `json:"created_by" gorm:"size:255;default:SYSTEM"`
	UpdatedBy   string          `json:"updated_by" gorm:"size:255;default:SYSTEM"`
	DeletedBy   string          `json:"deleted_by" gorm:"size:255"`
	CreatedAt   *time.Time      `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt   *time.Time      `json:"updated_at" gorm:"default:current_timestamp"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at"`

	OrderVoucher *OrderVoucher `json:"order_voucher,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:VoucherID;references:ID"`
}
