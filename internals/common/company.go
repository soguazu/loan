package common

import (
	"core_business/internals/core/domain"
	"core_business/pkg/utils"
	uuid "github.com/satori/go.uuid"
)

// CreateCompanyRequest DTO to create company
type CreateCompanyRequest struct {
	Company       uuid.UUID `json:"company"`
	Owner         string    `json:"owner" binding:"required"`
	Name          string    `json:"name" binding:"required"`
	Website       string    `json:"website"`
	Type          string    `json:"type" binding:"required"`
	FundingSource string    `json:"funding_source"`
	NoOfEmployee  int32     `json:"no_of_employee"`
}

// GetCompanyResponse DTO
type GetCompanyResponse struct {
	ID            uuid.UUID `json:"id" binding:"required"`
	Owner         string    `json:"owner" binding:"required"`
	Name          string    `json:"name" binding:"required"`
	Website       string    `json:"website"`
	Type          string    `json:"type"`
	FundingSource string    `json:"funding_source"`
	NoOfEmployee  int32     `json:"no_of_employee"`
}

// GetCompanyByIDRequest DTO to get company by id
type GetCompanyByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

// GetAllCompanyRequest DTO to get all company
type GetAllCompanyRequest struct {
	ParamID  int32 `form:"page_id;default=1" binding:"min=1"`
	PageSize int32 `form:"page_size;default=5" binding:"min=5"`
}

// GetCompany DTO to filter company
type GetCompany struct {
	Owner string `json:"owner,omitempty" form:"owner"`
	Name  string `json:"name,omitempty" form:"name"`
	Type  string `json:"type,omitempty" form:"type"`
}

// UpdateCompanyRequest DTO to update company
type UpdateCompanyRequest struct {
	Name          *string `json:"name,omitempty"`
	Type          *string `json:"type,omitempty"`
	Website       *string `json:"website,omitempty"`
	FundingSource *string `json:"funding_source,omitempty"`
	NoOfEmployee  *int32  `json:"no_of_employee,omitempty"`
}

// CreateCompanyResponse DTO get all companies
type CreateCompanyResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    domain.Company `json:"data"`
}

// Error struct
type Error struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Data to return generic data
type Data struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"error,omitempty"`
}

// Message struct
type Message struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// PassedCompanyTable for unit testing
type PassedCompanyTable struct {
	Company  CreateCompanyRequest
	TestName string
}

// CashBalanceRequest struct
type CashBalanceRequest struct {
	From      string `json:"from"`
	To        string `json:"to"`
	AccountID string `json:"account_id"`
	Page      int32  `json:"page"`
	Limit     int32  `json:"limit"`
}

// CashBalanceResponse struct
type CashBalanceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Data []struct {
			Transactions []struct {
				BankBalance float64 `json:"bank_balance"`
			} `json:"transactions"`
		} `json:"data"`
	} `json:"data"`
}

// KYCCheckRequest struct
type KYCCheckRequest struct {
	RCNumber    string `json:"rc_number"`
	CompanyName string `json:"company_name"`
}

// KYCCheckResponse struct
type KYCCheckResponse struct {
	Data KYCCheckPayload `json:"data"`
}

type KYCCheckPayload struct {
	Details KYCCheckNestedPayload `json:"details"`
}

type KYCCheckNestedPayload struct {
	RCNumber    string `json:"rc_number"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Date_Reg    string `json:"date_reg"`
}

// TransactionsCheckRequest struct
type TransactionsCheckRequest struct {
	From       string `json:"from"`
	To         string `json:"to"`
	CustomerID string `json:"customer_id"`
	Page       int32  `json:"page"`
	Limit      int32  `json:"limit"`
}

// TransactionCheckResponse struct
type TransactionCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Transaction []struct {
			Credit float64 `json:"credit"`
			Debit  float64 `json:"debit"`
		} `json:"transaction"`
	} `json:"data"`
}

// UnderWritingResponse struct
type UnderWritingResponse struct {
	CreditLimit float64 `json:"credit_limit"`
	TotalPoint  int32   `json:"total_point"`
}

// PassedTT random data for unit testing
var PassedTT = []PassedCompanyTable{
	{
		TestName: "All columns are complete",
		Company: CreateCompanyRequest{
			Company:       (&utils.Faker{}).RandomUUID(),
			Owner:         (&utils.Faker{}).RandomObjectID(),
			Name:          (&utils.Faker{}).RandomName(),
			Website:       (&utils.Faker{}).RandomWebsite(),
			Type:          (&utils.Faker{}).RandomType(),
			FundingSource: (&utils.Faker{}).RandomFundSource(),
			NoOfEmployee:  (&utils.Faker{}).RandomNoOfEmployee(),
		},
	},
	{
		TestName: "Except FundingSource and NoEmployee",
		Company: CreateCompanyRequest{
			Company: (&utils.Faker{}).RandomUUID(),
			Owner:   (&utils.Faker{}).RandomObjectID(),
			Name:    (&utils.Faker{}).RandomName(),
			Website: (&utils.Faker{}).RandomWebsite(),
			Type:    (&utils.Faker{}).RandomType(),
		},
	},
	{
		TestName: "With no NoOfEmployee",
		Company: CreateCompanyRequest{
			Company:       (&utils.Faker{}).RandomUUID(),
			Owner:         (&utils.Faker{}).RandomObjectID(),
			Name:          (&utils.Faker{}).RandomName(),
			Website:       (&utils.Faker{}).RandomWebsite(),
			Type:          (&utils.Faker{}).RandomType(),
			FundingSource: (&utils.Faker{}).RandomFundSource(),
			NoOfEmployee:  (&utils.Faker{}).RandomNoOfEmployee(),
		},
	},
}

// FailedTT for unit testing
var FailedTT = []PassedCompanyTable{
	{
		TestName: "No field was passed",
		Company:  CreateCompanyRequest{},
	},
	{
		TestName: "Owner wasn't passed",
		Company: CreateCompanyRequest{
			Name:          (&utils.Faker{}).RandomName(),
			Website:       (&utils.Faker{}).RandomWebsite(),
			Type:          (&utils.Faker{}).RandomType(),
			FundingSource: (&utils.Faker{}).RandomFundSource(),
			NoOfEmployee:  (&utils.Faker{}).RandomNoOfEmployee(),
		},
	},
	{
		TestName: "Type wasn't passed",
		Company: CreateCompanyRequest{
			Owner:         (&utils.Faker{}).RandomObjectID(),
			Name:          (&utils.Faker{}).RandomName(),
			Website:       (&utils.Faker{}).RandomWebsite(),
			FundingSource: (&utils.Faker{}).RandomFundSource(),
			NoOfEmployee:  (&utils.Faker{}).RandomNoOfEmployee(),
		},
	},
	{
		TestName: "NoOfEmployee wasn't passed",
		Company: CreateCompanyRequest{
			Owner:         (&utils.Faker{}).RandomObjectID(),
			Name:          (&utils.Faker{}).RandomName(),
			Website:       (&utils.Faker{}).RandomWebsite(),
			FundingSource: (&utils.Faker{}).RandomFundSource(),
		},
	},
}

// CreateCreditLimitIncreaseRequest DTO
type CreateCreditLimitIncreaseRequest struct {
	DesiredCreditLimit int64  `json:"desired_credit_limit" gorm:"index;not null"`
	Reason             string `json:"reason" gorm:"not null;index"`
	Status             bool   `json:"status" gorm:"not null;index;default:false"`
}

// UpdateCreditLimitIncreaseRequest DTO
type UpdateCreditLimitIncreaseRequest struct {
	DesiredCreditLimit *int64  `json:"desired_credit_limit" gorm:"index;not null"`
	Reason             *string `json:"reason" gorm:"not null;index"`
	Status             *bool   `json:"status" gorm:"not null;index;default:false"`
}

// GetCreditLimitIncrease DTO
type GetCreditLimitIncrease struct {
	ID                 uuid.UUID `json:"id" gorm:"not null;index"`
	Company            uuid.UUID `json:"company" gorm:"not null;index"`
	Wallet             uuid.UUID `json:"wallet" gorm:"not null;index"`
	Owner              uuid.UUID `json:"user" gorm:"not null;index"`
	CreditLimit        int64     `json:"credit_limit" gorm:"index;not null"`
	DesiredCreditLimit int64     `json:"desired_credit_limit" gorm:"index;not null"`
	Reason             string    `json:"reason" gorm:"not null;index"`
	Status             bool      `json:"status" gorm:"not null;index;default:false"`
}
