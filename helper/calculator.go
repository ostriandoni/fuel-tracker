package helper

import "github.com/shopspring/decimal"

func CalculateDistance(kmStart, kmEnd int) int {
	return kmEnd - kmStart
}

func CalculatePaidPerKm(totalPaid decimal.Decimal, distance int) decimal.Decimal {
	if distance == 0 {
		return decimal.Zero
	}
	return totalPaid.Div(decimal.NewFromInt(int64(distance)))
}

func CalculateUsageKmL(distance int, liters decimal.Decimal) decimal.Decimal {
	if liters.IsZero() {
		return decimal.Zero
	}
	return decimal.NewFromInt(int64(distance)).Div(liters)
}
