package domain

import (
	"github.com/satori/go.uuid"
)

// CreditIncrease model
type CreditIncrease struct {
	Base
	Company            uuid.UUID `json:"company" gorm:"not null;index;column:company"`
	Wallet             uuid.UUID `json:"wallet" gorm:"not null;index"`
	Owner              string    `json:"user" gorm:"not null;index"`
	CreditLimit        int64     `json:"credit_limit" gorm:"index;not null"`
	DesiredCreditLimit int64     `json:"desired_credit_limit" gorm:"index;not null"`
	Reason             string    `json:"reason" gorm:"not null;index"`
	Status             bool      `json:"status" gorm:"not null;index;default:false"`
}
