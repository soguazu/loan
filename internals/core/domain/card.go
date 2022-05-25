package domain

import (
	"github.com/satori/go.uuid"
)

type Channels struct {
	Atm    bool `json:"atm"`
	Pos    bool `json:"pos"`
	Web    bool `json:"web"`
	Mobile bool `json:"mobile"`
}

type SpendingLimits struct {
	Amount   int    `json:"amount"`
	Interval string `json:"interval"`
}

type SpendingControls struct {
	Channels          Channels       `gorm:"-"`
	AllowedCategories []string       `json:"allowedCategories"`
	BlockedCategories []string       `json:"blockedCategories"`
	SpendingLimits    SpendingLimits `gorm:"-"`
}

// Card model
type Card struct {
	Base
	Company           uuid.UUID        `json:"company,omitempty" gorm:"column:company"`
	Wallet            uuid.UUID        `json:"wallet,omitempty"`
	Name              string           `json:"name" gorm:"not null"`
	PartnerCustomerID string           `json:"customerId"`
	Type              string           `json:"type" gorm:"default:'virtual'"`
	Brand             string           `json:"brand"`
	Number            string           `json:"number"`
	Currency          string           `json:"currency" gorm:"default:'NG'"`
	Status            string           `json:"status" gorm:"default:'active'"`
	Lock              bool             `json:"lock" gorm:"default:false"`
	PartnerCardID     string           `json:"partner_card_id" gorm:"index;not null; unique"`
	Partner           string           `json:"partner" gorm:"index;not null;default:'sudo'"`
	CardAuth          string           `json:"card_auth" gorm:"size:4;default:1234"`
	Summary           string           `json:"summary"`
	Transactions      []Transaction    `json:"transaction,omitempty" gorm:"ForeignKey:Card;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Customer          uuid.UUID        `json:"customer" gorm:"column:customer"`
	SpendingControls  SpendingControls `json:"_" gorm:"-"`
	Business          string           `json:"business"`
	Account           string           `json:"account"`
	MaskedPan         string           `json:"maskedPan" gorm:"index;not null"`
	ExpiryMonth       string           `json:"expiryMonth" gorm:"index;not null"`
	ExpiryYear        string           `json:"expiryYear" gorm:"index;not null"`
}
