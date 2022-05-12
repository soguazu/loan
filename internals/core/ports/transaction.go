package ports

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ITransactionRepository defines the interface for transaction repository
type ITransactionRepository interface {
	GetByID(id string) (*domain.Transaction, error)
	GetTransactionByCompanyID(id string, pagination *utils.Pagination) (*utils.Pagination, error)
	GetTransactionByCardID(id string, pagination *utils.Pagination) (*utils.Pagination, error)
	Get(pagination *utils.Pagination) (*utils.Pagination, error)
	Persist(transaction *domain.Transaction) error
	Delete(id string) error
	DeleteAll() error
	WithTx(tx *gorm.DB) ITransactionRepository
}

// ITransactionService defines the interface for transaction service
type ITransactionService interface {
	GetTransactionByID(id string) (*domain.Transaction, error)
	GetTransactionByCompanyID(id string, pagination *utils.Pagination) (*utils.Pagination, error)
	GetTransactionByCardID(id string, pagination *utils.Pagination) (*utils.Pagination, error)
	GetAllTransaction(pagination *utils.Pagination) (*utils.Pagination, error)
	CreateTransaction(transaction *common.CreateTransactionRequest) error
	UpdateTransaction(id string, body common.UpdateTransactionRequest) (*domain.Transaction, error)
	DeleteTransaction(id string) error
	LockTransaction(id string) (*domain.Transaction, error)
}

// ITransactionHandler defines the interface for transaction handler
type ITransactionHandler interface {
	GetTransactionByID(c *gin.Context)
	GetTransactionByCompanyID(c *gin.Context)
	GetTransactionByCardID(c *gin.Context)
	GetAllTransaction(c *gin.Context)
	CreateTransaction(c *gin.Context)
	UpdateTransaction(c *gin.Context)
	DeleteTransaction(c *gin.Context)
	LockTransaction(c *gin.Context)
}
