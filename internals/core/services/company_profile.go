package services

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type companyProfileService struct {
	CompanyProfileRepository ports.ICompanyProfileRepository
	logger                   *log.Logger
}

// NewCompanyProfileService function create a new instance for service
func NewCompanyProfileService(cr ports.ICompanyProfileRepository, l *log.Logger) ports.ICompanyProfileService {
	return &companyProfileService{
		CompanyProfileRepository: cr,
		logger:                   l,
	}
}

func (cp *companyProfileService) GetCompanyProfileByID(id string) (*domain.CompanyProfile, error) {
	companyProfile, err := cp.CompanyProfileRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return companyProfile, nil
}

func (cp *companyProfileService) GetCompanyProfile(filter interface{}) ([]domain.CompanyProfile, error) {
	companyProfiles, err := cp.CompanyProfileRepository.GetBy(filter)
	if err != nil {
		cp.logger.Error(err)
		return nil, err
	}
	return companyProfiles, nil
}

func (cp *companyProfileService) GetAllCompanyProfile(pagination *utils.Pagination) (*utils.Pagination, error) {
	companyProfiles, err := cp.CompanyProfileRepository.Get(pagination)
	if err != nil {
		cp.logger.Error(err)
		return nil, err
	}
	return companyProfiles, nil
}

func (cp *companyProfileService) GetCompanyProfileBy(filter interface{}) ([]domain.CompanyProfile, error) {
	companyProfile, err := cp.CompanyProfileRepository.GetBy(filter)
	if err != nil {
		return nil, err
	}
	return companyProfile, nil
}

func (cp *companyProfileService) CreateCompanyProfile(companyProfile *domain.CompanyProfile) error {
	err := cp.CompanyProfileRepository.Persist(companyProfile)
	if err != nil {
		cp.logger.Error(err)
		return err
	}
	return nil
}

func (cp *companyProfileService) DeleteCompanyProfile(id string) error {
	err := cp.CompanyProfileRepository.Delete(id)
	if err != nil {
		cp.logger.Error(err)
		return err
	}
	return nil
}

func (cp *companyProfileService) UpdateCompanyProfile(id string, body common.UpdateCompanyProfileRequest) (*domain.CompanyProfile, error) {
	companyProfile, err := cp.CompanyProfileRepository.GetByID(id)
	if err != nil {
		cp.logger.Error(err)
		return nil, err
	}
	if body.RCNumber != nil {
		companyProfile.RCNumber = *body.RCNumber
	}

	if body.BusinessTin != nil {
		companyProfile.BusinessTin = *body.BusinessTin
	}

	if body.BusinessType != nil {
		companyProfile.BusinessType = *body.BusinessType
	}

	if body.BusinessActivity != nil {
		companyProfile.BusinessActivity = *body.BusinessActivity
	}

	if body.CACCertificateURL != nil {
		companyProfile.CACCertificateURL = *body.CACCertificateURL
	}

	if body.MermatURL != nil {
		companyProfile.MermatURL = *body.MermatURL
	}

	if body.StatusReportURL != nil {
		companyProfile.StatusReportURL = *body.StatusReportURL
	}

	if body.YearsInOperation != nil {
		companyProfile.YearsInOperation = *body.YearsInOperation
	}

	if body.IncorporationState != nil {
		companyProfile.IncorporationState = *body.IncorporationState
	}

	if body.IncorporationYear != nil {
		companyProfile.IncorporationYear = *body.IncorporationYear
	}

	err = cp.CompanyProfileRepository.Persist(companyProfile)

	if err != nil {
		cp.logger.Error(err)
		return nil, err
	}
	return companyProfile, nil
}
