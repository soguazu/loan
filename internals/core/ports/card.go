package ports

import (
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ICardRepository defines the interface for card repository
type ICardRepository interface {
	GetByID(id string) (*domain.Card, error)
	GetCardByCompanyID(id string, pagination *utils.Pagination) (*utils.Pagination, error)
	Get(pagination *utils.Pagination) (*utils.Pagination, error)
	Persist(card *domain.Card) error
	Delete(id string) error
	DeleteAll() error
	WithTx(tx *gorm.DB) ICardRepository
}

// ICardService defines the interface for card service
type ICardService interface {
	GetCardByID(id string) (*domain.Card, error)
	GetCardByCompanyID(id string, pagination *utils.Pagination) (*utils.Pagination, error)
	GetAllCard(pagination *utils.Pagination) (*utils.Pagination, error)
	CreateCard(address *domain.Card) error
	UpdateCard(params string, body common.UpdateAddressRequest) (*domain.Card, error)
	DeleteCard(id string) error
	LockCard(id string) (*domain.Card, error)
	CancelCard(id string) (*domain.Card, error)
	ChangeCardPin(id string) (*domain.Card, error)
}

// ICardHandler defines the interface for card handler
type ICardHandler interface {
	GetCardByID(c *gin.Context)
	GetCardByCompanyID(c *gin.Context)
	GetAllCard(c *gin.Context)
	CreateCard(c *gin.Context)
	UpdateCard(c *gin.Context)
	CancelCard(c *gin.Context)
	ChangeCardPin(c *gin.Context)
	DeleteCard(c *gin.Context)
	LockCard(c *gin.Context)
}
