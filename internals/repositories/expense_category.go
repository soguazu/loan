package repositories

import (
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	"gorm.io/gorm"
)

type expenseCategoryRepository struct {
	db *gorm.DB
}

// NewExpenseCategoryRepository creates a new instance expense category repository
func NewExpenseCategoryRepository(db *gorm.DB) ports.IExpenseCategoryRepository {
	return &expenseCategoryRepository{
		db: db,
	}
}

func (ec *expenseCategoryRepository) GetByID(id string) (*domain.ExpenseCategory, error) {
	var expenseCategory domain.ExpenseCategory
	if err := ec.db.Where("id = ?", id).First(&expenseCategory).Error; err != nil {
		return nil, err
	}
	return &expenseCategory, nil
}

func (ec *expenseCategoryRepository) Get(pagination *utils.Pagination) (*utils.Pagination, error) {
	var expenseCategory []domain.ExpenseCategory
	if err := ec.db.Scopes(utils.Paginate(expenseCategory, pagination, ec.db)).Find(&expenseCategory).Error; err != nil {
		return nil, err
	}
	pagination.Rows = expenseCategory
	return pagination, nil
}

func (ec *expenseCategoryRepository) Persist(expenseCategory *domain.ExpenseCategory) error {
	if expenseCategory.ID.String() != "" {
		if err := ec.db.Save(expenseCategory).Error; err != nil {
			return err
		}
		return nil
	}
	if err := ec.db.Create(&expenseCategory).Error; err != nil {
		return err
	}
	return nil
}

func (ec *expenseCategoryRepository) Delete(id string) error {
	if err := ec.db.Where("id = ?", id).Delete(&domain.ExpenseCategory{}).Error; err != nil {
		return err
	}
	return nil
}

func (ec *expenseCategoryRepository) DeleteAll() error {
	if err := ec.db.Exec("DELETE FROM expense_categories").Error; err != nil {
		return err
	}
	return nil
}

func (ec *expenseCategoryRepository) WithTx(tx *gorm.DB) ports.IExpenseCategoryRepository {
	return NewExpenseCategoryRepository(tx)
}
