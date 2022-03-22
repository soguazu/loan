package repositories

import (
	"fmt"
	"github.com/soguazu/core_business/internals/core/domain"
	"github.com/soguazu/core_business/internals/core/ports"
	"github.com/soguazu/core_business/pkg/utils"
	"gorm.io/gorm"
)

type companyRepository struct {
	db *gorm.DB
}

// NewCompanyRepository creates a new instance company repository
func NewCompanyRepository(db *gorm.DB) ports.ICompanyRepository {
	return &companyRepository{
		db: db,
	}
}

func (c *companyRepository) GetByID(id string) (*domain.Company, error) {
	var company domain.Company
	if err := c.db.Where("id = ?", id).First(&company).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (c *companyRepository) GetBy(filter interface{}) ([]domain.Company, error) {
	var company []domain.Company
	if err := c.db.Model(&domain.Company{}).Find(&company, filter).Error; err != nil {
		return nil, err
	}
	return company, nil
}

func (c *companyRepository) Get(pagination *utils.Pagination) (*utils.Pagination, error) {
	var companies []domain.Company
	if err := c.db.Scopes(utils.Paginate(companies, pagination, c.db)).Find(&companies).Error; err != nil {
		fmt.Println(err.Error(), "repo")
		return nil, err
	}
	pagination.Rows = companies
	return pagination, nil
}

func (c *companyRepository) Persist(company *domain.Company) (*domain.Company, error) {
	if company.ID.String() != "" {
		if err := c.db.Save(company).Error; err != nil {
			return nil, err
		}
		return company, nil
	}
	if err := c.db.Create(&company).Error; err != nil {
		return nil, err
	}
	return company, nil
}

func (c *companyRepository) Delete(id string) error {
	if err := c.db.Where("id = ?", id).Delete(&domain.Company{}).Error; err != nil {
		return err
	}
	return nil
}

func (c *companyRepository) DeleteAll() error {
	if err := c.db.Exec("DELETE FROM companies").Error; err != nil {
		return err
	}
	return nil
}

func (c *companyRepository) WithTx(tx *gorm.DB) ports.ICompanyRepository {
	return NewCompanyRepository(tx)
}
