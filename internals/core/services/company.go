package services

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/soguazu/core_business/internals/common"
	"github.com/soguazu/core_business/internals/core/domain"
	"github.com/soguazu/core_business/internals/core/ports"
	"github.com/soguazu/core_business/pkg/utils"
)

type companyService struct {
	CompanyRepository ports.ICompanyRepository
	logger            *log.Logger
}

// NewCompanyService function create a new instance for service
func NewCompanyService(cr ports.ICompanyRepository, l *log.Logger) ports.ICompanyService {
	return &companyService{
		CompanyRepository: cr,
		logger:            l,
	}
}

func (c *companyService) GetCompanyByID(id string) (*domain.Company, error) {
	company, err := c.CompanyRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (c *companyService) GetCompany(filter interface{}) ([]domain.Company, error) {
	companies, err := c.CompanyRepository.GetBy(filter)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}
	return companies, nil
}

func (c *companyService) GetAllCompany(pagination *utils.Pagination) (*utils.Pagination, error) {
	companies, err := c.CompanyRepository.Get(pagination)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}
	return companies, nil
}

func (c *companyService) CreateCompany(company *domain.Company) error {
	var entity []domain.Company
	entity, err := c.CompanyRepository.GetBy(domain.Company{Owner: company.Owner, Name: company.Name})
	if err != nil {
		return err
	}

	if len(entity) > 0 {
		return errors.New("already exist")
	}

	err = c.CompanyRepository.Persist(company)

	if err != nil {
		c.logger.Error(err)
		return err
	}

	return nil
}

func (c *companyService) DeleteCompany(id string) error {
	err := c.CompanyRepository.Delete(id)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	return nil
}

func (c *companyService) UpdateCompany(params common.GetByIDRequest, body common.UpdateCompanyRequest) (*domain.Company, error) {
	company, err := c.CompanyRepository.GetByID(params.ID)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}
	if body.Type != nil {
		company.Type = *body.Type
	}

	if body.Website != nil {
		company.Website = *body.Website
	}

	if body.Name != nil {
		company.Name = *body.Name
	}

	if body.NoOfEmployee != nil {
		company.NoOfEmployee = *body.NoOfEmployee
	}

	if body.FundingSource != nil {
		company.FundingSource = *body.FundingSource
	}

	err = c.CompanyRepository.Persist(company)

	if err != nil {
		c.logger.Error(err)
		return nil, err
	}
	return company, nil
}
