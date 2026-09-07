package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type FuelLog struct {
	ID            uint            `gorm:"primaryKey" json:"id"`
	Date          time.Time       `json:"date"`
	PricePerLiter decimal.Decimal `gorm:"type:decimal(12,2)" json:"price_per_liter"`
	TotalPaid     decimal.Decimal `gorm:"type:decimal(12,2)" json:"total_paid"`
	LitersFilled  decimal.Decimal `gorm:"type:decimal(10,2)" json:"liters_filled"`
	KmStart       int             `json:"km_start"`
	KmEnd         *int            `json:"km_end"`
	LocationID    uint            `json:"location_id"`
	Location      Location        `gorm:"foreignKey:LocationID" json:"location,omitempty"`
	PetrolTypeID  uint            `json:"petrol_type_id"`
	PetrolType    PetrolType      `gorm:"foreignKey:PetrolTypeID" json:"petrol_type,omitempty"`
	Notes         string          `gorm:"type:text" json:"notes"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	DeletedAt     gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}

func (FuelLog) TableName() string {
	return "fuel_logs"
}
