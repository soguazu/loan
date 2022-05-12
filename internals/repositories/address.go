package repositories

import (
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	"gorm.io/gorm"
)

type addressRepository struct {
	db *gorm.DB
}

// NewAddressRepository creates a new instance address repository
func NewAddressRepository(db *gorm.DB) ports.IAddressRepository {
	return &addressRepository{
		db: db,
	}
}

func (a *addressRepository) GetByID(id string) (*domain.Address, error) {
	var address domain.Address
	if err := a.db.Where("id = ?", id).
		Preload("Company").
		First(&address).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

func (a *addressRepository) GetBy(filter interface{}) ([]domain.Address, error) {
	var address []domain.Address
	if err := a.db.Model(&domain.Address{}).Find(&address, filter).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (a *addressRepository) Get(pagination *utils.Pagination) (*utils.Pagination, error) {
	var addresses []domain.Address
	if err := a.db.Scopes(utils.Paginate(addresses, pagination, a.db)).Find(&addresses).Error; err != nil {
		return nil, err
	}
	pagination.Rows = addresses
	return pagination, nil
}

func (a *addressRepository) Persist(address *domain.Address) error {
	if address.ID.String() != "" {
		if err := a.db.Save(address).Error; err != nil {
			return err
		}
		return nil
	}
	if err := a.db.Create(&address).Error; err != nil {
		return err
	}
	return nil
}

func (a *addressRepository) Delete(id string) error {
	if err := a.db.Where("id = ?", id).Delete(&domain.Address{}).Error; err != nil {
		return err
	}
	return nil
}

func (a *addressRepository) DeleteAll() error {
	if err := a.db.Exec("DELETE FROM addresses").Error; err != nil {
		return err
	}
	return nil
}

func (a *addressRepository) WithTx(tx *gorm.DB) ports.IAddressRepository {
	return NewAddressRepository(tx)
}
