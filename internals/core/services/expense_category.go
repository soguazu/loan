package services

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type expenseCategoryService struct {
	ExpenseCategoryRepository ports.IExpenseCategoryRepository
	logger                    *log.Logger
}

// NewExpenseCategoryService function create a new instance for service
func NewExpenseCategoryService(ecr ports.IExpenseCategoryRepository, l *log.Logger) ports.IExpenseCategoryService {
	return &expenseCategoryService{
		ExpenseCategoryRepository: ecr,
		logger:                    l,
	}
}

func (es *expenseCategoryService) GetExpenseCategoryByID(id string) (*domain.ExpenseCategory, error) {
	expenseCategory, err := es.ExpenseCategoryRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return expenseCategory, nil
}

func (es *expenseCategoryService) GetAllExpenseCategory(pagination *utils.Pagination) (*utils.Pagination, error) {
	companies, err := es.ExpenseCategoryRepository.Get(pagination)
	if err != nil {
		es.logger.Error(err)
		return nil, err
	}
	return companies, nil
}

func (es *expenseCategoryService) CreateExpenseCategory(expenseCategory *domain.ExpenseCategory) error {
	err := es.ExpenseCategoryRepository.Persist(expenseCategory)
	if err != nil {
		es.logger.Error(err)
		return err
	}
	return nil
}

func (es *expenseCategoryService) DeleteExpenseCategory(id string) error {
	err := es.ExpenseCategoryRepository.Delete(id)
	if err != nil {
		es.logger.Error(err)
		return err
	}
	return nil
}

func (es *expenseCategoryService) UpdateExpenseCategory(id string, body common.UpdateExpenseCategoryRequest) (*domain.ExpenseCategory, error) {
	expenseCategory, err := es.ExpenseCategoryRepository.GetByID(id)
	if err != nil {
		es.logger.Error(err)
		return nil, err
	}
	if body.Title != nil {
		expenseCategory.Title = *body.Title
	}

	err = es.ExpenseCategoryRepository.Persist(expenseCategory)

	if err != nil {
		es.logger.Error(err)
		return nil, err
	}
	return expenseCategory, nil
}
