package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/soguazu/core_business/internals/core/services"
	"github.com/soguazu/core_business/internals/handlers"
	"github.com/soguazu/core_business/internals/repositories"
	"github.com/soguazu/core_business/pkg/config"
	"github.com/soguazu/core_business/pkg/logger"
)

// Injection inject all dependencies
func Injection() {
	var logging *log.Logger

	if config.Instance.Env == "development" {
		logging = logger.NewLogger(log.New()).MakeLogger("logs/info", true)
		logging.Info("Log setup complete")
	} else {
		logging = logger.NewLogger(log.New()).Hook()
	}

	var (
		ginRoutes         = NewGinRouter(gin.Default())
		companyRepository = repositories.NewCompanyRepository(DBConnection)
		companyService    = services.NewCompanyService(companyRepository, logging)
		companyHandler    = handlers.NewCompanyHandler(companyService, logging, "Company")

		addressRepository = repositories.NewAddressRepository(DBConnection)
		addressService    = services.NewAddressService(addressRepository, logging)
		addressHandler    = handlers.NewAddressHandler(addressService, logging, "Address")

		businessHeadRepository = repositories.NewBusinessHeadRepository(DBConnection)
		businessHeadService    = services.NewBusinessHeadService(businessHeadRepository, logging)
		businessHeadHandler    = handlers.NewBusinessHeadHandler(businessHeadService, logging, "Business head")

		businessPartnerRepository = repositories.NewBusinessPartnerRepository(DBConnection)
		businessPartnerService    = services.NewBusinessPartnerService(businessPartnerRepository, logging)
		businessPartnerHandler    = handlers.NewBusinessPartnerHandler(businessPartnerService, logging, "Business partner")
	)

	v1 := ginRoutes.GROUP("v1")
	company := v1.Group("/company")
	company.GET("/:id", companyHandler.GetCompanyByID)
	company.GET("/", companyHandler.GetAllCompany)
	company.POST("/", companyHandler.CreateCompany)
	company.DELETE("/:id", companyHandler.DeleteCompany)
	company.PATCH("/:id", companyHandler.UpdateCompany)

	address := v1.Group("/address")
	address.GET("/:id", addressHandler.GetAddressByID)
	address.GET("/company/:id", addressHandler.GetAddressByCompanyID)
	address.GET("/", addressHandler.GetAllAddress)
	address.POST("/", addressHandler.CreateAddress)
	address.DELETE("/:id", addressHandler.DeleteAddress)
	address.PATCH("/:id", addressHandler.UpdateAddress)

	businessHead := v1.Group("/business_head")
	businessHead.GET("/:id", businessHeadHandler.GetBusinessHeadByID)
	businessHead.GET("/", businessHeadHandler.GetAllBusinessHead)
	businessHead.GET("/company/:id", businessHeadHandler.GetBusinessHeadByCompanyID)
	businessHead.POST("/", businessHeadHandler.CreateBusinessHead)
	businessHead.DELETE("/:id", businessHeadHandler.DeleteBusinessHead)
	businessHead.PATCH("/:id", businessHeadHandler.UpdateBusinessHead)

	businessPartner := v1.Group("/business_partner")
	businessPartner.GET("/:id", businessPartnerHandler.GetBusinessPartnerByID)
	businessPartner.GET("/", businessPartnerHandler.GetAllBusinessPartner)
	businessPartner.GET("/company/:id", businessPartnerHandler.GetBusinessPartnerByCompanyID)
	businessPartner.POST("/", businessPartnerHandler.CreateBusinessPartner)
	businessPartner.DELETE("/:id", businessPartnerHandler.DeleteBusinessPartner)
	businessPartner.PATCH("/:id", businessPartnerHandler.UpdateBusinessPartner)

	err := ginRoutes.SERVE()

	if err != nil {
		return
	}

}
