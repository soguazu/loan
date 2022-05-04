package server

import (
	"core_business/internals/core/services"
	"core_business/internals/handlers"
	"core_business/internals/repositories"
	"core_business/pkg/config"
	"core_business/pkg/logger"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
		ginRoutes = NewGinRouter(gin.Default())

		addressRepository = repositories.NewAddressRepository(DBConnection)
		addressService    = services.NewAddressService(addressRepository, logging)
		addressHandler    = handlers.NewAddressHandler(addressService, logging, "Address")

		businessHeadRepository = repositories.NewBusinessHeadRepository(DBConnection)
		businessHeadService    = services.NewBusinessHeadService(businessHeadRepository, logging)
		businessHeadHandler    = handlers.NewBusinessHeadHandler(businessHeadService, logging, "Business head")

		businessPartnerRepository = repositories.NewBusinessPartnerRepository(DBConnection)
		businessPartnerService    = services.NewBusinessPartnerService(businessPartnerRepository, logging)
		businessPartnerHandler    = handlers.NewBusinessPartnerHandler(businessPartnerService, logging, "Business partner")

		companyProfileRepository = repositories.NewCompanyProfileRepository(DBConnection)
		companyProfileService    = services.NewCompanyProfileService(companyProfileRepository, logging)
		companyProfileHandler    = handlers.NewCompanyProfileHandler(companyProfileService, logging, "Company profile")

		walletRepository = repositories.NewWalletRepository(DBConnection)
		walletService    = services.NewWalletService(walletRepository, DBConnection, logging)
		walletHandler    = handlers.NewWalletHandler(walletService, logging, "Wallet")

		expenseCategoryRepository = repositories.NewExpenseCategoryRepository(DBConnection)
		expenseCategoryService    = services.NewExpenseCategoryService(expenseCategoryRepository, logging)
		expenseCategoryHandler    = handlers.NewExpenseCategoryHandler(expenseCategoryService, logging, "Expense category")

		creditLimitRequestRepository = repositories.NewCreditLimitRequestRepository(DBConnection)

		companyRepository = repositories.NewCompanyRepository(DBConnection)
		companyService    = services.NewCompanyService(companyRepository, companyProfileRepository, walletRepository, creditLimitRequestRepository, logging)
		companyHandler    = handlers.NewCompanyHandler(companyService, logging, "Company")
	)

	v1 := ginRoutes.GROUP("v1")
	company := v1.Group("/company")
	company.GET("/:id", companyHandler.GetCompanyByID)
	company.GET("/", companyHandler.GetAllCompany)
	company.POST("/", companyHandler.CreateCompany)
	company.DELETE("/:id", companyHandler.DeleteCompany)
	company.PATCH("/:id", companyHandler.UpdateCompany)
	company.PATCH("/:id/under_writing", companyHandler.UnderWriting)
	company.PATCH("/:id/request_credit_limit_upgrade", companyHandler.RequestCreditLimitIncrease)
	company.PATCH("/:id/update_credit_limit", companyHandler.UpdateRequestCreditLimitIncrease)

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

	companyProfile := v1.Group("/company_profile")
	companyProfile.GET("/:id", companyProfileHandler.GetCompanyProfileByID)
	companyProfile.GET("/", companyProfileHandler.GetAllCompanyProfile)
	companyProfile.GET("/company/:id", companyProfileHandler.GetCompanyProfileByCompanyID)
	companyProfile.POST("/", companyProfileHandler.CreateCompanyProfile)
	companyProfile.DELETE("/:id", companyProfileHandler.DeleteCompanyProfile)
	companyProfile.PATCH("/:id", companyProfileHandler.UpdateCompanyProfile)

	wallet := v1.Group("/wallet")
	wallet.GET("/:id", walletHandler.GetWalletByID)
	wallet.POST("/webhook", walletHandler.CreateWallet)
	wallet.DELETE("/:id", walletHandler.DeleteWallet)
	wallet.PATCH("/:id", walletHandler.UpdateWallet)

	expenseCategory := v1.Group("/expense_category")
	expenseCategory.GET("/", expenseCategoryHandler.GetAllExpenseCategory)
	expenseCategory.GET("/:id", expenseCategoryHandler.GetExpenseCategoryByID)
	expenseCategory.POST("/", expenseCategoryHandler.CreateExpenseCategory)
	expenseCategory.DELETE("/:id", expenseCategoryHandler.DeleteExpenseCategory)
	expenseCategory.PATCH("/:id", expenseCategoryHandler.UpdateExpenseCategory)

	err := ginRoutes.SERVE()

	if err != nil {
		return
	}

}
