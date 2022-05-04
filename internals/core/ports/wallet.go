package ports

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IWalletRepository defines the interface for wallet repository
type IWalletRepository interface {
	GetByID(id string) (*domain.Wallet, error)
	GetByIDForUpdate(id string) (*domain.Wallet, error)
	GetBy(filter interface{}) ([]domain.Wallet, error)
	Persist(wallet *domain.Wallet) error
	Delete(id string) error
	DeleteAll() error
	WithTx(tx *gorm.DB) IWalletRepository
}

// IWalletService defines the interface for wallet service
type IWalletService interface {
	GetWalletByID(id string) (*domain.Wallet, error)
	CreateWallet(wallet *domain.Wallet) error
	UpdateWallet(id string, body common.UpdateWalletRequest) (*domain.Wallet, error)
	DeleteWallet(id string) error
}

// IWalletHandler defines the interface for wallet handler
type IWalletHandler interface {
	GetWalletByID(c *gin.Context)
	CreateWallet(c *gin.Context)
	DeleteWallet(c *gin.Context)
	UpdateWallet(c *gin.Context)
}
