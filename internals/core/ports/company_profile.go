package ports

import (
	"github.com/gin-gonic/gin"
	"core_business/internals/common"
	"core_business/internals/core/domain"
	"core_business/pkg/utils"
	"gorm.io/gorm"
)

// ICompanyProfileRepository defines the interface for company profile repository
type ICompanyProfileRepository interface {
	GetByID(id string) (*domain.CompanyProfile, error)
	GetBy(filter interface{}) ([]domain.CompanyProfile, error)
	Get(pagination *utils.Pagination) (*utils.Pagination, error)
	Persist(address *domain.CompanyProfile) error
	Delete(id string) error
	DeleteAll() error
	WithTx(tx *gorm.DB) ICompanyProfileRepository
}

// ICompanyProfileService defines the interface for company profile service
type ICompanyProfileService interface {
	GetCompanyProfileByID(id string) (*domain.CompanyProfile, error)
	GetCompanyProfileBy(filter interface{}) ([]domain.CompanyProfile, error)
	GetAllCompanyProfile(pagination *utils.Pagination) (*utils.Pagination, error)
	CreateCompanyProfile(companyProfile *domain.CompanyProfile) error
	UpdateCompanyProfile(params string, body common.UpdateCompanyProfileRequest) (*domain.CompanyProfile, error)
	DeleteCompanyProfile(id string) error
}

// ICompanyProfileHandler defines the interface for company profile handler
type ICompanyProfileHandler interface {
	GetCompanyProfileByID(c *gin.Context)
	GetCompanyProfileByCompanyID(c *gin.Context)
	GetAllCompanyProfile(c *gin.Context)
	CreateCompanyProfile(c *gin.Context)
	DeleteCompanyProfile(c *gin.Context)
	UpdateCompanyProfile(c *gin.Context)
}
