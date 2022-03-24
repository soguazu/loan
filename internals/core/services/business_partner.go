package services

import (
	log "github.com/sirupsen/logrus"
	"github.com/soguazu/core_business/internals/common"
	"github.com/soguazu/core_business/internals/core/domain"
	"github.com/soguazu/core_business/internals/core/ports"
	"github.com/soguazu/core_business/pkg/utils"
)

type businessPartnerService struct {
	BusinessPartnerRepository ports.IBusinessPartnerRepository
	logger                    *log.Logger
}

// NewBusinessPartnerService function create a new instance for service
func NewBusinessPartnerService(cr ports.IBusinessPartnerRepository, l *log.Logger) ports.IBusinessPartnerService {
	return &businessPartnerService{
		BusinessPartnerRepository: cr,
		logger:                    l,
	}
}

func (bp *businessPartnerService) GetBusinessPartnerByID(id string) (*domain.BusinessPartner, error) {
	businessPartner, err := bp.BusinessPartnerRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return businessPartner, nil
}

func (bp *businessPartnerService) GetBusinessPartnerBy(filter interface{}) ([]domain.BusinessPartner, error) {
	businessPartners, err := bp.BusinessPartnerRepository.GetBy(filter)
	if err != nil {
		bp.logger.Error(err)
		return nil, err
	}
	return businessPartners, nil
}

func (bp *businessPartnerService) GetAllBusinessPartner(pagination *utils.Pagination) (*utils.Pagination, error) {
	businessHeads, err := bp.BusinessPartnerRepository.Get(pagination)
	if err != nil {
		bp.logger.Error(err)
		return nil, err
	}
	return businessHeads, nil
}

func (bp *businessPartnerService) CreateBusinessPartner(businessPartner *domain.BusinessPartner) error {
	err := bp.BusinessPartnerRepository.Persist(businessPartner)
	if err != nil {
		bp.logger.Error(err)
		return err
	}
	return nil
}

func (bp *businessPartnerService) DeleteBusinessPartner(id string) error {
	err := bp.BusinessPartnerRepository.Delete(id)
	if err != nil {
		bp.logger.Error(err)
		return err
	}
	return nil
}

func (bp *businessPartnerService) UpdateBusinessPartner(id string, body common.UpdateBusinessPartnerRequest) (*domain.BusinessPartner, error) {
	businessPartner, err := bp.BusinessPartnerRepository.GetByID(id)
	if err != nil {
		bp.logger.Error(err)
		return nil, err
	}
	if body.Name != nil {
		businessPartner.Name = *body.Name
	}

	if body.Phone != nil {
		businessPartner.Phone = *body.Phone
	}

	err = bp.BusinessPartnerRepository.Persist(businessPartner)

	if err != nil {
		bp.logger.Error(err)
		return nil, err
	}
	return businessPartner, nil
}
