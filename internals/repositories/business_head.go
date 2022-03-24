package repositories

import (
	"fmt"
	"github.com/soguazu/core_business/internals/core/domain"
	"github.com/soguazu/core_business/internals/core/ports"
	"github.com/soguazu/core_business/pkg/utils"
	"gorm.io/gorm"
)

type businessHeadRepository struct {
	db *gorm.DB
}

// NewBusinessHeadRepository creates a new instance address repository
func NewBusinessHeadRepository(db *gorm.DB) ports.IBusinessHeadRepository {
	return &businessHeadRepository{
		db: db,
	}
}

func (bh *businessHeadRepository) GetByID(id string) (*domain.BusinessHead, error) {
	var businessHead domain.BusinessHead
	if err := bh.db.Where("id = ?", id).First(&businessHead).Error; err != nil {
		return nil, err
	}
	return &businessHead, nil
}

func (bh *businessHeadRepository) GetBy(filter interface{}) ([]domain.BusinessHead, error) {
	var businessHead []domain.BusinessHead
	if err := bh.db.Model(&domain.BusinessHead{}).Find(&businessHead, filter).Error; err != nil {
		return nil, err
	}
	fmt.Println(businessHead)
	return businessHead, nil
}

func (bh *businessHeadRepository) Get(pagination *utils.Pagination) (*utils.Pagination, error) {
	var businessHead []domain.BusinessHead
	if err := bh.db.Scopes(utils.Paginate(businessHead, pagination, bh.db)).Find(&businessHead).Error; err != nil {
		return nil, err
	}
	pagination.Rows = businessHead
	return pagination, nil
}

func (bh *businessHeadRepository) Persist(businessHead *domain.BusinessHead) error {
	if businessHead.ID.String() != "" {
		if err := bh.db.Save(businessHead).Error; err != nil {
			return err
		}
		return nil
	}
	if err := bh.db.Create(&businessHead).Error; err != nil {
		return err
	}
	return nil
}

func (bh *businessHeadRepository) Delete(id string) error {
	if err := bh.db.Where("id = ?", id).Delete(&domain.BusinessHead{}).Error; err != nil {
		return err
	}
	return nil
}

func (bh *businessHeadRepository) DeleteAll() error {
	if err := bh.db.Exec("DELETE FROM business_heads").Error; err != nil {
		return err
	}
	return nil
}

func (bh *businessHeadRepository) WithTx(tx *gorm.DB) ports.IBusinessHeadRepository {
	return NewBusinessHeadRepository(tx)
}
