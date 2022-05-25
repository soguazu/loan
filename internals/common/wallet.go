package common

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// UpdateWalletRequest DTO to update wallet
type UpdateWalletRequest struct {
	CreditLimit     *int64  `json:"credit_limit,omitempty"`
	PreviousBalance *int64  `json:"previous_balance,omitempty"`
	CurrentSpending *int64  `json:"current_spending,omitempty"`
	Payment         *int64  `json:"payment,omitempty"`
	Type            *string `json:"type,omitempty"`
}

// GetWalletResponse DTO
type GetWalletResponse struct {
	ID              uuid.UUID `json:"id,omitempty"`
	Company         uuid.UUID `json:"company,omitempty"`
	CreditLimit     int64     `json:"credit_limit"`
	PreviousBalance int64     `json:"previous_balance"`
	CurrentSpending int64     `json:"current_spending"`
	CashBackPayment int64     `json:"cash_back_payment"`
	AccountID       string    `json:"account_id"`
	CustomerID      string    `json:"customerId"`
	Status          bool      `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

// CreateWalletResponse DTO
type CreateWalletResponse struct {
	ID              uuid.UUID `json:"id,omitempty"`
	Company         uuid.UUID `json:"company,omitempty"`
	CreditLimit     int64     `json:"credit_limit"`
	PreviousBalance int64     `json:"previous_balance"`
	CurrentSpending int64     `json:"current_spending"`
	CashBackPayment int64     `json:"cash_back_payment"`
	AccountID       string    `json:"account_id"`
	CustomerID      string    `json:"customerId"`
	Status          bool      `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

// CreateWalletDataResponse returns wallet response
type CreateWalletDataResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Data    CreateWalletResponse `json:"data"`
}

// GetWalletDataResponse returns wallet response
type GetWalletDataResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    GetWalletResponse `json:"data"`
}

// CreateWalletRequest DTO to create wallet
type CreateWalletRequest struct {
	Company    uuid.UUID `json:"company" binding:"required"`
	AccountID  string    `json:"accountId" binding:"required"`
	CustomerID string    `json:"customerId" binding:"required"`
}
