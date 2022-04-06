package ports

import (
	"github.com/gin-gonic/gin"
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/pkg/utils"
	"gorm.io/gorm"
)

// IBusinessHeadRepository defines the interface for business head repository
type IBusinessHeadRepository interface {
	GetByID(id string) (*domain.BusinessHead, error)
	GetBy(filter interface{}) ([]domain.BusinessHead, error)
	Get(pagination *utils.Pagination) (*utils.Pagination, error)
	Persist(businessHead *domain.BusinessHead) error
	Delete(id string) error
	DeleteAll() error
	WithTx(tx *gorm.DB) IBusinessHeadRepository
}

// IBusinessHeadService defines the interface for business head service
type IBusinessHeadService interface {
	GetBusinessHeadByID(id string) (*domain.BusinessHead, error)
	GetBusinessHeadBy(filter interface{}) ([]domain.BusinessHead, error)
	GetAllBusinessHead(pagination *utils.Pagination) (*utils.Pagination, error)
	CreateBusinessHead(address *domain.BusinessHead) error
	UpdateBusinessHead(params string, body common.UpdateBusinessHeadRequest) (*domain.BusinessHead, error)
	DeleteBusinessHead(id string) error
}

// IBusinessHeadHandler defines the interface for address handler
type IBusinessHeadHandler interface {
	GetBusinessHeadByID(c *gin.Context)
	GetAllBusinessHead(c *gin.Context)
	GetBusinessHeadByCompanyID(c *gin.Context)
	CreateBusinessHead(c *gin.Context)
	DeleteBusinessHead(c *gin.Context)
	UpdateBusinessHead(c *gin.Context)
}
