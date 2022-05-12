package repositories

import (
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

// NewCustomerRepository creates a new instance customer repository
func NewCustomerRepository(db *gorm.DB) ports.ICustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (c *customerRepository) GetByID(id string) (*domain.Customer, error) {
	var customer domain.Customer
	if err := c.db.Where("id = ?", id).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *customerRepository) GetBy(filter interface{}) (*domain.Customer, error) {
	var customer *domain.Customer
	if err := c.db.Model(&domain.Customer{}).Find(&customer, filter).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *customerRepository) Get(pagination *utils.Pagination) (*utils.Pagination, error) {
	var customers []domain.Customer
	if err := c.db.Scopes(utils.Paginate(customers, pagination, c.db)).Find(&customers).Error; err != nil {
		return nil, err
	}
	pagination.Rows = customers
	return pagination, nil
}

func (c *customerRepository) Persist(customer *domain.Customer) error {
	if customer.ID.String() != "" {
		if err := c.db.Save(customer).Error; err != nil {
			return err
		}
		return nil
	}
	if err := c.db.Create(&customer).Error; err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) Delete(id string) error {
	if err := c.db.Where("id = ?", id).Delete(&domain.Customer{}).Error; err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) DeleteAll() error {
	if err := c.db.Exec("DELETE FROM customers").Error; err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) WithTx(tx *gorm.DB) ports.ICustomerRepository {
	return NewCustomerRepository(tx)
}
