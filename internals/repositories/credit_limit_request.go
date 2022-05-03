package repositories

import (
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"gorm.io/gorm"
)

type creditLimitRequestRepository struct {
	db *gorm.DB
}

// NewCreditLimitRequestRepository creates a new instance wallet repository
func NewCreditLimitRequestRepository(db *gorm.DB) ports.ICreditLimitRequestRepository {
	return &creditLimitRequestRepository{
		db: db,
	}
}

func (c *creditLimitRequestRepository) GetByID(id string) (*domain.CreditIncrease, error) {
	var creditLimit domain.CreditIncrease
	if err := c.db.Where("id = ?", id).First(&creditLimit).Error; err != nil {
		return nil, err
	}
	return &creditLimit, nil
}

func (c *creditLimitRequestRepository) GetBy(filter interface{}) (*domain.CreditIncrease, error) {
	var creditLimit []domain.CreditIncrease
	if err := c.db.Model(&domain.CreditIncrease{}).Find(&creditLimit, filter).Error; err != nil {
		return nil, err
	}
	return &creditLimit[0], nil
}

func (c *creditLimitRequestRepository) Persist(creditLimitRequest *domain.CreditIncrease) error {
	if creditLimitRequest.ID.String() != "" {
		if err := c.db.Save(creditLimitRequest).Error; err != nil {
			return err
		}
		return nil
	}
	if err := c.db.Create(&creditLimitRequest).Error; err != nil {
		return err
	}
	return nil
}

func (c *creditLimitRequestRepository) Delete(id string) error {
	if err := c.db.Where("id = ?", id).Delete(&domain.CreditIncrease{}).Error; err != nil {
		return err
	}
	return nil
}

func (c *creditLimitRequestRepository) DeleteAll() error {
	if err := c.db.Exec("DELETE FROM credit_increases").Error; err != nil {
		return err
	}
	return nil
}

func (c *creditLimitRequestRepository) WithTx(tx *gorm.DB) ports.ICreditLimitRequestRepository {
	return NewCreditLimitRequestRepository(tx)
}
