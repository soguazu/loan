package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/soguazu/core_business/internals/common"
	"github.com/soguazu/core_business/internals/core/domain"
	"github.com/soguazu/core_business/pkg/utils"
	"gorm.io/gorm"
)

// IAddressRepository defines the interface for address repository
type IAddressRepository interface {
	GetByID(id string) (*domain.Address, error)
	GetBy(filter interface{}) ([]domain.Address, error)
	Get(pagination *utils.Pagination) (*utils.Pagination, error)
	Persist(address *domain.Address) error
	Delete(id string) error
	DeleteAll() error
	WithTx(tx *gorm.DB) IAddressRepository
}

// IAddressService defines the interface for address service
type IAddressService interface {
	GetAddressByID(id string) (*domain.Address, error)
	GetAddressBy(filter interface{}) ([]domain.Address, error)
	GetAllAddress(pagination *utils.Pagination) (*utils.Pagination, error)
	CreateAddress(address *domain.Address) error
	UpdateAddress(params string, body common.UpdateAddressRequest) (*domain.Address, error)
	DeleteAddress(id string) error
}

// IAddressHandler defines the interface for address handler
type IAddressHandler interface {
	GetAddressByID(c *gin.Context)
	GetAddressByCompanyID(c *gin.Context)
	GetAllAddress(c *gin.Context)
	CreateAddress(c *gin.Context)
	DeleteAddress(c *gin.Context)
	UpdateAddress(c *gin.Context)
}
