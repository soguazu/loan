package ports

import (
	"core_business/internals/core/domain"
	"gorm.io/gorm"
)

// ICreditLimitRequestRepository defines the interface for credit limit request repository
type ICreditLimitRequestRepository interface {
	GetByID(id string) (*domain.CreditIncrease, error)
	GetBy(filter interface{}) (*domain.CreditIncrease, error)
	Persist(creditLimitRequest *domain.CreditIncrease) error
	Delete(id string) error
	DeleteAll() error
	WithTx(tx *gorm.DB) ICreditLimitRequestRepository
}
