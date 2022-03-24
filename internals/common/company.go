package common

import (
	uuid "github.com/satori/go.uuid"
	"github.com/soguazu/core_business/internals/core/domain"
	"github.com/soguazu/core_business/pkg/utils"
)

// CreateCompanyRequest DTO to create company
type CreateCompanyRequest struct {
	Owner         uuid.UUID `json:"owner" binding:"required"`
	Name          string    `json:"name" binding:"required"`
	Website       string    `json:"website"`
	Type          string    `json:"type" binding:"required"`
	FundingSource string    `json:"funding_source"`
	NoOfEmployee  int32     `json:"no_of_employee"`
}

// GetCompanyResponse DTO
type GetCompanyResponse struct {
	ID            uuid.UUID `json:"id" binding:"required"`
	Owner         uuid.UUID `json:"owner" binding:"required"`
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
	Owner uuid.UUID `json:"owner,omitempty" form:"owner"`
	Name  string    `json:"name,omitempty" form:"name"`
	Type  string    `json:"type,omitempty" form:"type"`
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

// PassedTT random data for unit testing
var PassedTT = []PassedCompanyTable{
	{
		TestName: "All columns are complete",
		Company: CreateCompanyRequest{
			Owner:         (&utils.Faker{}).RandomUUID(),
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
			Owner:   (&utils.Faker{}).RandomUUID(),
			Name:    (&utils.Faker{}).RandomName(),
			Website: (&utils.Faker{}).RandomWebsite(),
			Type:    (&utils.Faker{}).RandomType(),
		},
	},
	{
		TestName: "With no NoOfEmployee",
		Company: CreateCompanyRequest{
			Owner:         (&utils.Faker{}).RandomUUID(),
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
			Owner:         (&utils.Faker{}).RandomUUID(),
			Name:          (&utils.Faker{}).RandomName(),
			Website:       (&utils.Faker{}).RandomWebsite(),
			FundingSource: (&utils.Faker{}).RandomFundSource(),
			NoOfEmployee:  (&utils.Faker{}).RandomNoOfEmployee(),
		},
	},
	{
		TestName: "NoOfEmployee wasn't passed",
		Company: CreateCompanyRequest{
			Owner:         (&utils.Faker{}).RandomUUID(),
			Name:          (&utils.Faker{}).RandomName(),
			Website:       (&utils.Faker{}).RandomWebsite(),
			FundingSource: (&utils.Faker{}).RandomFundSource(),
		},
	},
}
