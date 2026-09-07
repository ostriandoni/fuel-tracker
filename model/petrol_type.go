package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type PetrolType struct {
	ID         uint            `gorm:"primaryKey" json:"id"`
	LocationID uint            `gorm:"not null" json:"location_id"`
	Location   Location        `gorm:"foreignKey:LocationID" json:"location,omitempty"`
	Name       string          `gorm:"type:varchar(255);not null" json:"name"`
	Price      decimal.Decimal `gorm:"type:decimal(12,2);not null" json:"price"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	DeletedAt  gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}

func (PetrolType) TableName() string { return "petrol_types" }
