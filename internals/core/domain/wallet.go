package domain

import (
	"github.com/satori/go.uuid"
)

// Wallet model
type Wallet struct {
	Base
	Company         uuid.UUID `json:"company" gorm:"not null;index;column:company"`
	CreditLimit     int64     `json:"credit_limit" gorm:"index;not null"`
	PreviousBalance int64     `json:"previous_balance" gorm:"default:0;not null"`
	CurrentSpending int64     `json:"current_spending" gorm:"default:0;not null"`
	AvailableCredit int64     `json:"available_credit" gorm:"default:0;not null"`
	TotalBalance    int64     `json:"total_balance" gorm:"default:0;not null"`
	CashBackPayment int64     `json:"cash_back_payment"`
	AccountID       string    `json:"account_id" gorm:"not null;index"`
	CustomerID      string    `json:"customerId" gorm:"not null;index"`
	Status          bool      `json:"status" gorm:"not null;index"`
}
