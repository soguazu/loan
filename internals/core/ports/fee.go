package ports

import (
	"core_business/internals/core/domain"
	"gorm.io/gorm"
)

// IFeeRepository defines the interface for fee repository
type IFeeRepository interface {
	GetByID(id string) (*domain.Fee, error)
	GetByIdentifier(identifier string) (*domain.Fee, error)
	Delete(id string) error
	DeleteAll() error
	WithTx(tx *gorm.DB) IFeeRepository
}
