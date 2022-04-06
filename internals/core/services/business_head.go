package services

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type businessHeadService struct {
	BusinessHeadRepository ports.IBusinessHeadRepository
	logger                 *log.Logger
}

// NewBusinessHeadService function create a new instance for service
func NewBusinessHeadService(cr ports.IBusinessHeadRepository, l *log.Logger) ports.IBusinessHeadService {
	return &businessHeadService{
		BusinessHeadRepository: cr,
		logger:                 l,
	}
}

func (bs *businessHeadService) GetBusinessHeadByID(id string) (*domain.BusinessHead, error) {
	businessHead, err := bs.BusinessHeadRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return businessHead, nil
}

func (bs *businessHeadService) GetBusinessHead(filter interface{}) ([]domain.BusinessHead, error) {
	businessHeads, err := bs.BusinessHeadRepository.GetBy(filter)
	if err != nil {
		bs.logger.Error(err)
		return nil, err
	}
	return businessHeads, nil
}

func (bs *businessHeadService) GetAllBusinessHead(pagination *utils.Pagination) (*utils.Pagination, error) {
	businessHeads, err := bs.BusinessHeadRepository.Get(pagination)
	if err != nil {
		bs.logger.Error(err)
		return nil, err
	}
	return businessHeads, nil
}

func (bs *businessHeadService) GetBusinessHeadBy(filter interface{}) ([]domain.BusinessHead, error) {
	businessHead, err := bs.BusinessHeadRepository.GetBy(filter)
	if err != nil {
		return nil, err
	}
	return businessHead, nil
}

func (bs *businessHeadService) CreateBusinessHead(businessHead *domain.BusinessHead) error {
	err := bs.BusinessHeadRepository.Persist(businessHead)
	if err != nil {
		bs.logger.Error(err)
		return err
	}
	return nil
}

func (bs *businessHeadService) DeleteBusinessHead(id string) error {
	err := bs.BusinessHeadRepository.Delete(id)
	if err != nil {
		bs.logger.Error(err)
		return err
	}
	return nil
}

func (bs *businessHeadService) UpdateBusinessHead(id string, body common.UpdateBusinessHeadRequest) (*domain.BusinessHead, error) {
	businessHead, err := bs.BusinessHeadRepository.GetByID(id)
	if err != nil {
		bs.logger.Error(err)
		return nil, err
	}
	if body.JobTitle != nil {
		businessHead.JobTitle = *body.JobTitle
	}

	if body.Phone != nil {
		businessHead.Phone = *body.Phone
	}

	if body.IdentificationType != nil {
		businessHead.IdentificationType = *body.IdentificationType
	}

	if body.IdentificationNumber != nil {
		businessHead.IdentificationNumber = *body.IdentificationNumber
	}

	if body.IdentificationImageURL != nil {
		businessHead.IdentificationImageURL = *body.IdentificationImageURL
	}

	if body.CompanyIDUrl != nil {
		businessHead.CompanyIDUrl = *body.CompanyIDUrl
	}

	err = bs.BusinessHeadRepository.Persist(businessHead)

	if err != nil {
		bs.logger.Error(err)
		return nil, err
	}
	return businessHead, nil
}
