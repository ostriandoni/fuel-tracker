package dto

import (
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type CreateFuelLogRequest struct {
	Date          time.Time       `json:"date" binding:"required"`
	PricePerLiter decimal.Decimal `json:"price_per_liter" binding:"required"`
	TotalPaid     decimal.Decimal `json:"total_paid" binding:"required"`
	LitersFilled  decimal.Decimal `json:"liters_filled" binding:"required"`
	KmStart       int             `json:"km_start" binding:"required"`
	KmEnd         int             `json:"km_end" binding:"required"`
	LocationID    uint            `json:"location_id"`
	PetrolTypeID  uint            `json:"petrol_type_id"`
	Notes         string          `json:"notes"`
}

type UpdateFuelLogRequest struct {
	Date          time.Time       `json:"date"`
	PricePerLiter decimal.Decimal `json:"price_per_liter"`
	TotalPaid     decimal.Decimal `json:"total_paid"`
	LitersFilled  decimal.Decimal `json:"liters_filled"`
	KmStart       int             `json:"km_start"`
	KmEnd         int             `json:"km_end"`
	LocationID    uint            `json:"location_id"`
	PetrolTypeID  uint            `json:"petrol_type_id"`
	Notes         string          `json:"notes"`
}

type FuelLogResponse struct {
	ID             uint             `json:"id"`
	Month          int              `json:"month"`
	Date           time.Time        `json:"date"`
	PricePerLiter  decimal.Decimal  `json:"price_per_liter"`
	TotalPaid      decimal.Decimal  `json:"total_paid"`
	LitersFilled   decimal.Decimal  `json:"liters_filled"`
	KmStart        int              `json:"km_start"`
	KmEnd          *int             `json:"km_end"`
	Distance       *int             `json:"distance"`
	PaidPerKm      *decimal.Decimal `json:"paid_per_km"`
	UsageKmL       *decimal.Decimal `json:"usage_km_l"`
	LocationID     uint             `json:"location_id"`
	LocationName   string           `json:"location_name"`
	PetrolTypeID   uint             `json:"petrol_type_id"`
	PetrolTypeName string           `json:"petrol_type_name"`
	Notes          string           `json:"notes"`
}

func (f FuelLogResponse) KmEndDisplay() string {
	if f.KmEnd == nil {
		return "-"
	}
	return strconv.Itoa(*f.KmEnd)
}

func (f FuelLogResponse) DistanceDisplay() string {
	if f.Distance == nil {
		return "-"
	}
	return strconv.Itoa(*f.Distance)
}

func (f FuelLogResponse) PaidPerKmDisplay() string {
	if f.PaidPerKm == nil {
		return "-"
	}
	return f.PaidPerKm.StringFixed(2)
}

func (f FuelLogResponse) UsageKmLDisplay() string {
	if f.UsageKmL == nil {
		return "-"
	}
	return f.UsageKmL.StringFixed(2)
}
