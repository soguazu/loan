package utils

import (
	"strconv"
)

// CurrencyUnit data
type CurrencyUnit struct {
	Amount    float64
	Precision int
}

func (c *CurrencyUnit) ToMinorUnit(amount float64) int64 {
	c.Precision = 2
	bottom := "100"

	if c.Precision >= 2 {
		for i := 0; i < c.Precision; i++ {
			bottom += "0"
		}
	}

	denominator, _ := strconv.ParseFloat(bottom, 64)
	c.Amount = amount * denominator
	return int64(c.Amount)
}

func (c *CurrencyUnit) ToMajorUnit(amount float64) int64 {
	c.Precision = 2
	bottom := "100"

	if c.Precision >= 2 {
		for i := 0; i < c.Precision; i++ {
			bottom += "0"
		}
	}

	denominator, _ := strconv.ParseFloat(bottom, 64)
	c.Amount = amount / denominator
	return int64(c.Amount)
}
