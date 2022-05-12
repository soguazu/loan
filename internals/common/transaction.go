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

type CreateTransactionRequest struct {
	Environment string `json:"environment"`
	Business    string `json:"business"`
	Data        struct {
		Object struct {
			Id       string `json:"_id"`
			Business string `json:"business"`
			Customer struct {
				Id         string `json:"_id"`
				Business   string `json:"business"`
				Type       string `json:"type"`
				Name       string `json:"name"`
				Status     string `json:"status"`
				Individual struct {
					FirstName string `json:"firstName"`
					LastName  string `json:"lastName"`
					Id        string `json:"_id"`
				} `json:"individual"`
				BillingAddress struct {
					Line1      string `json:"line1"`
					Line2      string `json:"line2"`
					City       string `json:"city"`
					State      string `json:"state"`
					Country    string `json:"country"`
					PostalCode string `json:"postalCode"`
					Id         string `json:"_id"`
				} `json:"billingAddress"`
				IsDeleted    bool      `json:"isDeleted"`
				CreatedAt    time.Time `json:"createdAt"`
				UpdatedAt    time.Time `json:"updatedAt"`
				V            int       `json:"__v"`
				EmailAddress string    `json:"emailAddress"`
				PhoneNumber  string    `json:"phoneNumber"`
			} `json:"customer"`
			Account struct {
				Id                string    `json:"_id"`
				Business          string    `json:"business"`
				Type              string    `json:"type"`
				Currency          string    `json:"currency"`
				AccountName       string    `json:"accountName"`
				BankCode          string    `json:"bankCode"`
				AccountType       string    `json:"accountType"`
				AccountNumber     string    `json:"accountNumber"`
				CurrentBalance    int       `json:"currentBalance"`
				AvailableBalance  int       `json:"availableBalance"`
				Provider          string    `json:"provider"`
				ProviderReference string    `json:"providerReference"`
				ReferenceCode     string    `json:"referenceCode"`
				IsDefault         bool      `json:"isDefault"`
				IsDeleted         bool      `json:"isDeleted"`
				CreatedAt         time.Time `json:"createdAt"`
				UpdatedAt         time.Time `json:"updatedAt"`
				V                 int       `json:"__v"`
			} `json:"account"`
			Card struct {
				Id            string `json:"_id"`
				Business      string `json:"business"`
				Customer      string `json:"customer"`
				Account       string `json:"account"`
				FundingSource struct {
					Id         string `json:"_id"`
					Business   string `json:"business"`
					Type       string `json:"type"`
					Status     string `json:"status"`
					JitGateway struct {
						Url                 string `json:"url"`
						AuthorizationHeader string `json:"authorizationHeader"`
						AuthorizeByDefault  bool   `json:"authorizeByDefault"`
						Id                  string `json:"_id"`
					} `json:"jitGateway"`
					IsDefault bool      `json:"isDefault"`
					IsDeleted bool      `json:"isDeleted"`
					CreatedAt time.Time `json:"createdAt"`
					UpdatedAt time.Time `json:"updatedAt"`
					V         int       `json:"__v"`
				} `json:"fundingSource"`
				Type          string    `json:"type"`
				Brand         string    `json:"brand"`
				Currency      string    `json:"currency"`
				MaskedPan     string    `json:"maskedPan"`
				ExpiryMonth   string    `json:"expiryMonth"`
				ExpiryYear    string    `json:"expiryYear"`
				Status        string    `json:"status"`
				Is2FAEnrolled bool      `json:"is2FAEnrolled"`
				IsDeleted     bool      `json:"isDeleted"`
				CreatedAt     time.Time `json:"createdAt"`
				UpdatedAt     time.Time `json:"updatedAt"`
				V             int       `json:"__v"`
			} `json:"card"`
			Amount              int           `json:"amount"`
			Fee                 int           `json:"fee"`
			Vat                 int           `json:"vat"`
			Approved            bool          `json:"approved"`
			Currency            string        `json:"currency"`
			Status              string        `json:"status"`
			AuthorizationMethod string        `json:"authorizationMethod"`
			BalanceTransactions []interface{} `json:"balanceTransactions"`
			MerchantAmount      int           `json:"merchantAmount"`
			MerchantCurrency    string        `json:"merchantCurrency"`
			Merchant            struct {
				Category   string `json:"category"`
				Name       string `json:"name"`
				MerchantId string `json:"merchantId"`
				City       string `json:"city"`
				State      string `json:"state"`
				Country    string `json:"country"`
				PostalCode string `json:"postalCode"`
				Id         string `json:"_id"`
			} `json:"merchant"`
			Terminal struct {
				Rrn                          string `json:"rrn"`
				Stan                         string `json:"stan"`
				TerminalId                   string `json:"terminalId"`
				TerminalOperatingEnvironment string `json:"terminalOperatingEnvironment"`
				TerminalAttendance           string `json:"terminalAttendance"`
				TerminalType                 string `json:"terminalType"`
				PanEntryMode                 string `json:"panEntryMode"`
				PinEntryMode                 string `json:"pinEntryMode"`
				CardHolderPresence           bool   `json:"cardHolderPresence"`
				CardPresence                 bool   `json:"cardPresence"`
				Id                           string `json:"_id"`
			} `json:"terminal"`
			TransactionMetadata struct {
				Channel   string `json:"channel"`
				Type      string `json:"type"`
				Reference string `json:"reference"`
				Id        string `json:"_id"`
			} `json:"transactionMetadata"`
			PendingRequest struct {
				Amount           int    `json:"amount"`
				Currency         string `json:"currency"`
				MerchantAmount   int    `json:"merchantAmount"`
				MerchantCurrency string `json:"merchantCurrency"`
				Id               string `json:"_id"`
			} `json:"pendingRequest"`
			RequestHistory []interface{} `json:"requestHistory"`
			Verification   struct {
				BillingAddressLine1      string `json:"billingAddressLine1"`
				BillingAddressPostalCode string `json:"billingAddressPostalCode"`
				Cvv                      string `json:"cvv"`
				Expiry                   string `json:"expiry"`
				Pin                      string `json:"pin"`
				ThreeDSecure             string `json:"threeDSecure"`
				SafeToken                string `json:"safeToken"`
				Authentication           string `json:"authentication"`
				Id                       string `json:"_id"`
			} `json:"verification"`
			IsDeleted  bool      `json:"isDeleted"`
			CreatedAt  time.Time `json:"createdAt"`
			UpdatedAt  time.Time `json:"updatedAt"`
			FeeDetails []struct {
				Contract    string `json:"contract"`
				Currency    string `json:"currency"`
				Amount      int    `json:"amount"`
				Description string `json:"description"`
				Id          string `json:"_id"`
			} `json:"feeDetails"`
			V int `json:"__v"`
		} `json:"object"`
		Id      string `json:"_id"`
		Changes string `json:"changes"`
	} `json:"data"`
	Type            string `json:"type"`
	PendingWebhook  bool   `json:"pendingWebhook"`
	WebhookArchived bool   `json:"webhookArchived"`
	CreatedAt       int    `json:"createdAt"`
	Id              string `json:"_id"`
}

type LogTransactionRequest struct {
	Wallet            uuid.UUID
	Company           uuid.UUID
	PartnerCustomerID string
	Customer          uuid.UUID
	Charges           []float64
	Transaction       CreateTransactionRequest
}
