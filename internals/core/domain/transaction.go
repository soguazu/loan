package domain

import (
	"github.com/satori/go.uuid"
)

// TransactionChannel channel used for withdrawal pos, web, atm, card etc
type TransactionChannel string

// TransactionStatus Pending, Success, Failed
type TransactionStatus string

// TransactionEntry debit or credit
type TransactionEntry string

// TransactionType withdrawal, cashback, interest, shipping, cards, fee, refund
type TransactionType string

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

// Transaction model
type Transaction struct {
	Base
	Company           uuid.UUID          `json:"company,omitempty" gorm:"column:company"`
	Wallet            uuid.UUID          `json:"wallet,omitempty" gorm:"column:wallet"`
	Card              uuid.UUID          `json:"card,omitempty" gorm:"column:card"`
	PartnerCardID     string             `json:"partner_card_id" gorm:"not null"`
	Customer          uuid.UUID          `json:"customer,omitempty" gorm:"column:customer"`
	PartnerCustomerID string             `json:"partner_customer_id"`
	Debit             float64            `json:"debit"`
	Credit            float64            `json:"credit"`
	Note              string             `json:"note" gorm:"not null"`
	ReferenceID       string             `json:"reference_id"`
	PartnerFee        float64            `json:"partner_fee"`
	Fee               []uuid.UUID        `json:"fee" gorm:"type:text;column:fee"`
	Status            TransactionStatus  `json:"status" gorm:"index;not null;"`               // Pending, Success, Failed
	Entry             TransactionEntry   `json:"entry" gorm:"index;not null;default:'debit'"` // debit or credit
	Channel           TransactionChannel `json:"channel" gorm:"index;not null;"`              // channel used for withdrawal pos, web, atm, card etc
	Reason            string             `json:"reason"`
	Type              TransactionType    `json:"type" gorm:"not null"` // withdrawal, cashback, interest, shipping, cards, fee, refund
	CardType          CardType           `json:"card_type"`
	ParentID          string             `json:"parent_id"` //Parent id for the refund
	Lock              bool               `json:"lock" gorm:"default:false"`
	Receipt           string             `json:"receipt"`
	ExpenseCategory   string             `json:"expense_category"`
}
