package domain

import (
	"github.com/satori/go.uuid"
)

// Card model
type Card struct {
	Base
	Company           uuid.UUID     `json:"company,omitempty" gorm:"column:company"`
	Wallet            uuid.UUID     `json:"wallet,omitempty"`
	PartnerCustomerID string        `json:"customerId"`
	Type              string        `json:"type" gorm:"default:'virtual'"`
	Brand             string        `json:"brand"`
	Number            string        `json:"number"`
	Currency          string        `json:"currency" gorm:"default:'NG'"`
	Status            string        `json:"status" gorm:"default:'active'"`
	Lock              bool          `json:"lock" gorm:"default:false"`
	PartnerID         string        `json:"-" gorm:"index;not null; unique"`
	Partner           string        `json:"-" gorm:"index;not null;default:'sudo'"`
	CardAuth          string        `json:"card_auth" gorm:"size:4;not null;default:1234"`
	PartnerRef        string        `json:"-" gorm:"index;unique"`
	Summary           string        `json:"summary"`
	Transactions      []Transaction `json:"transaction,omitempty" gorm:"ForeignKey:Card;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Customer          uuid.UUID     `json:"customer" gorm:"column:customer"`
	SpendingControls  `json:"spendingControls"`
}

// SpendingControls set spending limit
type SpendingControls struct {
	Channels       `json:"channels"`
	SpendingLimits `json:"spendingLimits"`
}

// Channels available means of transaction
type Channels struct {
	Atm bool `json:"atm"`
	Web bool `json:"web"`
	Pos bool `json:"pos"`
}

//SpendingLimits set transaction limit on intervals
type SpendingLimits struct {
	Amount   int64  `json:"amount"`
	Interval string `json:"interval"`
}
