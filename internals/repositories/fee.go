package repositories

import (
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"gorm.io/gorm"
	"strings"
)

type feeRepository struct {
	db *gorm.DB
}

// NewFeeRepository creates a new instance fee repository
func NewFeeRepository(db *gorm.DB) ports.IFeeRepository {
	return &feeRepository{
		db: db,
	}
}

func (f *feeRepository) GetByID(id string) (*domain.Fee, error) {
	var fee domain.Fee
	if err := f.db.Where("id = ?", id).
		First(&fee).Error; err != nil {
		return nil, err
	}
	return &fee, nil
}

func (f *feeRepository) GetByIdentifier(identifier string) (*domain.Fee, error) {
	var fee domain.Fee
	if err := f.db.Where("identifier = ?", strings.ToLower(identifier)).
		First(&fee).Error; err != nil {
		return nil, err
	}
	return &fee, nil
}

func (f *feeRepository) Delete(id string) error {
	if err := f.db.Where("id = ?", id).Delete(&domain.Fee{}).Error; err != nil {
		return err
	}
	return nil
}

func (f *feeRepository) DeleteAll() error {
	if err := f.db.Exec("DELETE FROM fees").Error; err != nil {
		return err
	}
	return nil
}

func (f *feeRepository) WithTx(tx *gorm.DB) ports.IFeeRepository {
	return NewFeeRepository(tx)
}
