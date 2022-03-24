package domain

import (
	"github.com/satori/go.uuid"
)

// Address model
type Address struct {
	Base
	Company            uuid.UUID `json:"company,omitempty" gorm:"foreignKey:Company;references:ID"`
	Address            string    `json:"address" gorm:"index;not null"`
	ApartmentUnitFloor int32     `json:"apartment_unit_floor"`
	City               string    `json:"city" gorm:"not null"`
	State              string    `json:"state" gorm:"not null"`
	Country            string    `json:"country" gorm:"not null"`
	UtilityBill        string    `json:"utility_bill" gorm:"not null"`
}
