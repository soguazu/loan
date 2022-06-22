package repositories

import (
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"gorm.io/gorm"
)

type panRepository struct {
	db *gorm.DB
}

// NewPANRepository creates a new instance pan repository
func NewPANRepository(db *gorm.DB) ports.IPANRepository {
	return &panRepository{
		db: db,
	}
}

func (c *panRepository) GetByID(id string) (*domain.PAN, error) {
	var pan domain.PAN
	if err := c.db.Where("id = ?", id).
		First(&pan).Error; err != nil {
		return nil, err
	}
	return &pan, nil
}

func (c *panRepository) GetFirstOne() (*domain.PAN, error) {
	var pan domain.PAN
	if err := c.db.Where("status = ?", true).
		First(&pan).Error; err != nil {
		return nil, err
	}
	return &pan, nil
}

func (c *panRepository) Delete(id string) error {
	if err := c.db.Where("id = ?", id).Delete(&domain.PAN{}).Error; err != nil {
		return err
	}
	return nil
}

func (c *panRepository) DeleteAll() error {
	if err := c.db.Exec("DELETE FROM panss").Error; err != nil {
		return err
	}
	return nil
}

func (c *panRepository) BatchInsert(pans []domain.PAN) error {
	if err := c.db.Create(&pans).Error; err != nil {
		return err
	}
	return nil
}

func (c *panRepository) WithTx(tx *gorm.DB) ports.IPANRepository {
	return NewPANRepository(tx)
}
