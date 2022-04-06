package ports

import (
	"github.com/gin-gonic/gin"
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/pkg/utils"
	"gorm.io/gorm"
)

// IBusinessPartnerRepository defines the interface for business partner repository
type IBusinessPartnerRepository interface {
	GetByID(id string) (*domain.BusinessPartner, error)
	GetBy(filter interface{}) ([]domain.BusinessPartner, error)
	Get(pagination *utils.Pagination) (*utils.Pagination, error)
	Persist(address *domain.BusinessPartner) error
	Delete(id string) error
	DeleteAll() error
	WithTx(tx *gorm.DB) IBusinessPartnerRepository
}

// IBusinessPartnerService defines the interface for business partner service
type IBusinessPartnerService interface {
	GetBusinessPartnerByID(id string) (*domain.BusinessPartner, error)
	GetBusinessPartnerBy(filter interface{}) ([]domain.BusinessPartner, error)
	GetAllBusinessPartner(pagination *utils.Pagination) (*utils.Pagination, error)
	CreateBusinessPartner(businessPartner *domain.BusinessPartner) error
	UpdateBusinessPartner(params string, body common.UpdateBusinessPartnerRequest) (*domain.BusinessPartner, error)
	DeleteBusinessPartner(id string) error
}

// IBusinessPartnerHandler defines the interface for business partner handler
type IBusinessPartnerHandler interface {
	GetBusinessPartnerByID(c *gin.Context)
	GetBusinessPartnerByCompanyID(c *gin.Context)
	GetAllBusinessPartner(c *gin.Context)
	CreateBusinessPartner(c *gin.Context)
	DeleteBusinessPartner(c *gin.Context)
	UpdateBusinessPartner(c *gin.Context)
}
