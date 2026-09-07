package helper

import "github.com/shopspring/decimal"

func CalculateDistance(kmStart int, kmEnd *int) *int {
	if kmEnd == nil {
		return nil
	}
	d := *kmEnd - kmStart
	return &d
}

func CalculatePaidPerKm(totalPaid decimal.Decimal, distance *int) *decimal.Decimal {
	if distance == nil || *distance == 0 {
		return nil
	}
	v := totalPaid.Div(decimal.NewFromInt(int64(*distance)))
	return &v
}

func CalculateUsageKmL(distance *int, liters decimal.Decimal) *decimal.Decimal {
	if distance == nil || liters.IsZero() {
		return nil
	}
	v := decimal.NewFromInt(int64(*distance)).Div(liters)
	return &v
}
