package ports

import (
	"core_business/internals/core/domain"
	"gorm.io/gorm"
)

// IPANRepository defines the interface for pan repository
type IPANRepository interface {
	GetByID(id string) (*domain.PAN, error)
	GetFirstOne() (*domain.PAN, error)
	BatchInsert(pans []domain.PAN) error
	Delete(id string) error
	DeleteAll() error
	WithTx(tx *gorm.DB) IPANRepository
}
