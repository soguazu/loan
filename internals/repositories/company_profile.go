package repositories

import (
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	"gorm.io/gorm"
)

type companyProfileRepository struct {
	db *gorm.DB
}

// NewCompanyProfileRepository creates a new instance company profile repository
func NewCompanyProfileRepository(db *gorm.DB) ports.ICompanyProfileRepository {
	return &companyProfileRepository{
		db: db,
	}
}

func (cp *companyProfileRepository) GetByID(id string) (*domain.CompanyProfile, error) {
	var companyProfile domain.CompanyProfile
	if err := cp.db.Where("id = ?", id).First(&companyProfile).Error; err != nil {
		return nil, err
	}
	return &companyProfile, nil
}

func (cp *companyProfileRepository) GetBy(filter interface{}) ([]domain.CompanyProfile, error) {
	var companyProfile []domain.CompanyProfile
	if err := cp.db.Model(&domain.CompanyProfile{}).Find(&companyProfile, filter).Error; err != nil {
		return nil, err
	}
	return companyProfile, nil
}

func (cp *companyProfileRepository) Get(pagination *utils.Pagination) (*utils.Pagination, error) {
	var companyProfile []domain.CompanyProfile
	if err := cp.db.Scopes(utils.Paginate(companyProfile, pagination, cp.db)).Find(&companyProfile).Error; err != nil {
		return nil, err
	}
	pagination.Rows = companyProfile
	return pagination, nil
}

func (cp *companyProfileRepository) Persist(companyProfile *domain.CompanyProfile) error {
	if companyProfile.ID.String() != "" {
		if err := cp.db.Save(companyProfile).Error; err != nil {
			return err
		}
		return nil
	}
	if err := cp.db.Create(&companyProfile).Error; err != nil {
		return err
	}
	return nil
}

func (cp *companyProfileRepository) Delete(id string) error {
	if err := cp.db.Where("id = ?", id).Delete(&domain.CompanyProfile{}).Error; err != nil {
		return err
	}
	return nil
}

func (cp *companyProfileRepository) DeleteAll() error {
	if err := cp.db.Exec("DELETE FROM company_profiles").Error; err != nil {
		return err
	}
	return nil
}

func (cp *companyProfileRepository) WithTx(tx *gorm.DB) ports.ICompanyProfileRepository {
	return NewCompanyProfileRepository(tx)
}
