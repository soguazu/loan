package common

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// TransactionChannel channel used for withdrawal pos, web, atm, card etc
type TransactionChannel string

// TransactionStatus Pending, Success, Failed
type TransactionStatus string

// TransactionEntry debit or credit
type TransactionEntry string

// CardType physical or virtual
type CardType string

const (
	WebChannel TransactionChannel = "WEB"
	PosChannel TransactionChannel = "POS"
	AtmChannel TransactionChannel = "ATM"

	PendingStatus   TransactionStatus = "PENDING"
	SuccessStatus   TransactionStatus = "SUCCESS"
	FailedStatus    TransactionStatus = "FAILED"
	CancelledStatus TransactionStatus = "CANCELLED"
	AbandonedStatus TransactionStatus = "ABANDONED"

	DebitEntry  TransactionEntry = "DEBIT"
	CreditEntry TransactionEntry = "CREDIT"

	RefundType       TransactionType = "REFUND"
	WithdrawalType   TransactionType = "WITHDRAWAL"
	CashbackType     TransactionType = "CASHBACK"
	FeeType          TransactionType = "FEE"
	InterestType     TransactionType = "INTEREST"
	CardCreationType TransactionType = "CARD"
	ShippingType     TransactionType = "SHIPPING"

	PhysicalType CardType = "PHYSICAL"
	VirtualType  CardType = "VIRTUAL"
)

// UpdateTransactionRequest DTO to update transaction files
type UpdateTransactionRequest struct {
	Receipt         *string `json:"receipt,omitempty"`
	ExpenseCategory *string `json:"expenseCategory,omitempty"`
}

type GetTransactionResponse struct {
	ID                uuid.UUID          `json:"id"`
	Company           uuid.UUID          `json:"company,omitempty"`
	Wallet            uuid.UUID          `json:"wallet,omitempty"`
	Card              uuid.UUID          `json:"card,omitempty"`
	Customer          uuid.UUID          `json:"customer,omitempty"`
	PartnerCustomerID string             `json:"partner_customer_id"`
	Debit             float64            `json:"debit"`
	Credit            float64            `json:"credit"`
	Note              string             `json:"note"`
	ReferenceID       string             `json:"reference_id"`
	PartnerFee        float64            `json:"partner_fee"`
	Fee               []uuid.UUID        `json:"fee"`
	Status            TransactionStatus  `json:"status"`  // Pending, Success, Failed
	Entry             TransactionEntry   `json:"entry"`   // debit or credit
	Channel           TransactionChannel `json:"channel"` // channel used for withdrawal pos, web, atm, card etc
	Reason            string             `json:"reason"`
	Type              TransactionType    `json:"type"` // withdrawal, cashback, interest, shipping, cards, fee, refund
	CardType          CardType           `json:"card_type"`
	ParentID          string             `json:"parent_id"` //Parent id for the refund
	Lock              bool               `json:"lock"`
	Receipt           string             `json:"receipt"`
	ExpenseCategory   uuid.UUID          `json:"expense_category,omitempty"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}

// GetSingleTransactionResponse DTO get a transaction
type GetSingleTransactionResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    GetTransactionResponse `json:"data"`
}
