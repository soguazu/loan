package domain

import (
	"github.com/satori/go.uuid"
)

// Customer model
type Customer struct {
	Base
	Company            uuid.UUID     `json:"company,omitempty" gorm:"column:company"`
	Wallet             uuid.UUID     `json:",omitempty"`
	Address            string        `json:"address" gorm:"index;not null"`
	ApartmentUnitFloor int32         `json:"apartment_unit_floor"`
	City               string        `json:"city" gorm:"not null"`
	State              string        `json:"state" gorm:"not null"`
	Country            string        `json:"country" gorm:"not null"`
	UtilityBill        string        `json:"utility_bill" gorm:"not null"`
	PostalCode         string        `json:"postalCode" gorm:"default:'100001'"`
	Transactions       []Transaction `json:"transaction" gorm:"ForeignKey:Customer;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Card               []Card        `json:"card,omitempty" gorm:"ForeignKey:Customer;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
