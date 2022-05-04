package utils

import (
	"strconv"
)

// ToMinorUnit converts to kobo
func ToMinorUnit(amount float64) int64 {
	Precision := 2
	bottom := "100"

	if Precision > 2 {
		for i := 0; i < Precision; i++ {
			bottom += "0"
		}
	}

	denominator, _ := strconv.ParseFloat(bottom, 64)
	Amount := amount * denominator
	return int64(Amount)
}

// ToMajorUnit converts to naira
func ToMajorUnit(amount float64) float64 {
	Precision := 2
	bottom := "100"

	if Precision >= 2 {
		for i := 0; i < Precision; i++ {
			bottom += "0"
		}
	}

	denominator, _ := strconv.ParseFloat(bottom, 64)
	Amount := amount / denominator
	return Amount
}
