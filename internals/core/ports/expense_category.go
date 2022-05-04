package ports

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IExpenseCategoryRepository defines the interface for category repository
type IExpenseCategoryRepository interface {
	GetByID(id string) (*domain.ExpenseCategory, error)
	Get(pagination *utils.Pagination) (*utils.Pagination, error)
	Persist(expenseCategory *domain.ExpenseCategory) error
	Delete(id string) error
	DeleteAll() error
	WithTx(tx *gorm.DB) IExpenseCategoryRepository
}

// IExpenseCategoryService defines the interface for category service
type IExpenseCategoryService interface {
	GetExpenseCategoryByID(id string) (*domain.ExpenseCategory, error)
	GetAllExpenseCategory(pagination *utils.Pagination) (*utils.Pagination, error)
	CreateExpenseCategory(expenseCategory *domain.ExpenseCategory) error
	UpdateExpenseCategory(id string, company common.UpdateExpenseCategoryRequest) (*domain.ExpenseCategory, error)
	DeleteExpenseCategory(id string) error
}

// IExpenseCategoryHandler defines the interface for company handler
type IExpenseCategoryHandler interface {
	GetExpenseCategoryByID(c *gin.Context)
	GetAllExpenseCategory(c *gin.Context)
	CreateExpenseCategory(c *gin.Context)
	DeleteExpenseCategory(c *gin.Context)
	UpdateExpenseCategory(c *gin.Context)
}
