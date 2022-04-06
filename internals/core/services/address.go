package services

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type addressService struct {
	AddressRepository ports.IAddressRepository
	logger            *log.Logger
}

// NewAddressService function create a new instance for service
func NewAddressService(cr ports.IAddressRepository, l *log.Logger) ports.IAddressService {
	return &addressService{
		AddressRepository: cr,
		logger:            l,
	}
}

func (as *addressService) GetAddressByID(id string) (*domain.Address, error) {
	company, err := as.AddressRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (as *addressService) GetAddress(filter interface{}) ([]domain.Address, error) {
	companies, err := as.AddressRepository.GetBy(filter)
	if err != nil {
		as.logger.Error(err)
		return nil, err
	}
	return companies, nil
}

func (as *addressService) GetAllAddress(pagination *utils.Pagination) (*utils.Pagination, error) {
	companies, err := as.AddressRepository.Get(pagination)
	if err != nil {
		as.logger.Error(err)
		return nil, err
	}
	return companies, nil
}

func (as *addressService) GetAddressBy(filter interface{}) ([]domain.Address, error) {
	addresses, err := as.AddressRepository.GetBy(filter)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func (as *addressService) CreateAddress(address *domain.Address) error {
	err := as.AddressRepository.Persist(address)
	if err != nil {
		as.logger.Error(err)
		return err
	}
	return nil
}

func (as *addressService) DeleteAddress(id string) error {
	err := as.AddressRepository.Delete(id)
	if err != nil {
		as.logger.Error(err)
		return err
	}
	return nil
}

func (as *addressService) UpdateAddress(id string, body common.UpdateAddressRequest) (*domain.Address, error) {
	address, err := as.AddressRepository.GetByID(id)
	if err != nil {
		as.logger.Error(err)
		return nil, err
	}
	if body.Address != nil {
		address.Address = *body.Address
	}

	if body.ApartmentUnitFloor != nil {
		address.ApartmentUnitFloor = *body.ApartmentUnitFloor
	}

	if body.City != nil {
		address.City = *body.City
	}

	if body.State != nil {
		address.State = *body.State
	}

	if body.Country != nil {
		address.Country = *body.Country
	}

	err = as.AddressRepository.Persist(address)

	if err != nil {
		as.logger.Error(err)
		return nil, err
	}
	return address, nil
}
