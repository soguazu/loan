package domain

import (
	"github.com/satori/go.uuid"
)

// Wallet model
type Wallet struct {
	Base
	Company         uuid.UUID `json:"company" gorm:"not null;index"`
	CreditLimit     int64     `json:"credit_limit" gorm:"index;not null"`
	PreviousBalance int64     `json:"previous_balance" gorm:"default:0;not null"`
	CurrentBalance  int64     `json:"current_balance" gorm:"default:0;not null"`
	Payment         int64     `json:"payment"`
	AccountID       string    `json:"account_id" gorm:"not null;index"`
	CustomerID      string    `json:"customerId" gorm:"not null;index"`
	Status          bool      `json:"status" gorm:"not null;index"`
}
