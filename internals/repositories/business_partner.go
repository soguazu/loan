package repositories

import (
	"github.com/soguazu/core_business/internals/core/domain"
	"github.com/soguazu/core_business/internals/core/ports"
	"github.com/soguazu/core_business/pkg/utils"
	"gorm.io/gorm"
)

type businessPartnerRepository struct {
	db *gorm.DB
}

// NewBusinessPartnerRepository creates a new instance business partner repository
func NewBusinessPartnerRepository(db *gorm.DB) ports.IBusinessPartnerRepository {
	return &businessPartnerRepository{
		db: db,
	}
}

func (bp *businessPartnerRepository) GetByID(id string) (*domain.BusinessPartner, error) {
	var businessPartner domain.BusinessPartner
	if err := bp.db.Where("id = ?", id).First(&businessPartner).Error; err != nil {
		return nil, err
	}
	return &businessPartner, nil
}

func (bp *businessPartnerRepository) GetBy(filter interface{}) ([]domain.BusinessPartner, error) {
	var businessPartner []domain.BusinessPartner
	if err := bp.db.Model(&domain.BusinessPartner{}).Find(&businessPartner, filter).Error; err != nil {
		return nil, err
	}
	return businessPartner, nil
}

func (bp *businessPartnerRepository) Get(pagination *utils.Pagination) (*utils.Pagination, error) {
	var businessPartners []domain.BusinessPartner
	if err := bp.db.Scopes(utils.Paginate(businessPartners, pagination, bp.db)).Find(&businessPartners).Error; err != nil {
		return nil, err
	}
	pagination.Rows = businessPartners
	return pagination, nil
}

func (bp *businessPartnerRepository) Persist(businessPartner *domain.BusinessPartner) error {
	if businessPartner.ID.String() != "" {
		if err := bp.db.Save(businessPartner).Error; err != nil {
			return err
		}
		return nil
	}
	if err := bp.db.Create(&businessPartner).Error; err != nil {
		return err
	}
	return nil
}

func (bp *businessPartnerRepository) Delete(id string) error {
	if err := bp.db.Where("id = ?", id).Delete(&domain.BusinessPartner{}).Error; err != nil {
		return err
	}
	return nil
}

func (bp *businessPartnerRepository) DeleteAll() error {
	if err := bp.db.Exec("DELETE FROM business_partners").Error; err != nil {
		return err
	}
	return nil
}

func (bp *businessPartnerRepository) WithTx(tx *gorm.DB) ports.IBusinessPartnerRepository {
	return NewBusinessPartnerRepository(tx)
}
