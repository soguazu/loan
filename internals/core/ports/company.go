package ports

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ICompanyRepository defines the interface for company repository
type ICompanyRepository interface {
	GetByID(id string) (*domain.Company, error)
	GetBy(filter interface{}) ([]domain.Company, error)
	Get(pagination *utils.Pagination) (*utils.Pagination, error)
	Persist(company *domain.Company) error
	Delete(id string) error
	DeleteAll() error
	WithTx(tx *gorm.DB) ICompanyRepository
}

// ICompanyService defines the interface for company service
type ICompanyService interface {
	GetCompanyByID(id string) (*domain.Company, error)
	GetAllCompany(pagination *utils.Pagination) (*utils.Pagination, error)
	CreateCompany(company *domain.Company) error
	UpdateCompany(params common.GetByIDRequest, company common.UpdateCompanyRequest) (*domain.Company, error)
	DeleteCompany(id string) error
	UnderWriting(id string) (*common.UnderWritingResponse, error)
	RequestCreditLimitIncrease(id string, body *domain.CreditIncrease) error
	UpdateRequestCreditLimitIncrease(params common.GetByIDRequest, body common.UpdateCreditLimitIncreaseRequest) error
}

// ICompanyHandler defines the interface for company handler
type ICompanyHandler interface {
	GetCompanyByID(c *gin.Context)
	GetAllCompany(c *gin.Context)
	CreateCompany(c *gin.Context)
	DeleteCompany(c *gin.Context)
	UpdateCompany(c *gin.Context)
	UnderWriting(c *gin.Context)
	RequestCreditLimitIncrease(c *gin.Context)
	UpdateRequestCreditLimitIncrease(c *gin.Context)
}
