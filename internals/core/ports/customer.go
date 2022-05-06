package ports

import (
	"core_business/internals/core/domain"
	"core_business/pkg/utils"
	"gorm.io/gorm"
)

// ICustomerRepository defines the interface for customer repository
type ICustomerRepository interface {
	GetByID(id string) (*domain.Customer, error)
	Get(pagination *utils.Pagination) (*utils.Pagination, error)
	Persist(customer *domain.Customer) error
	Delete(id string) error
	DeleteAll() error
	WithTx(tx *gorm.DB) ICustomerRepository
}
