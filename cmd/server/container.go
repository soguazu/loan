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
		companyHandler    = handlers.NewCompanyHandler(companyService, logging)
	)

	v1 := ginRoutes.GROUP("v1")
	company := v1.Group("/company")
	company.GET("/:id", companyHandler.GetCompanyByID)
	company.GET("/", companyHandler.GetAllCompany)
	company.POST("/", companyHandler.CreateCompany)
	company.DELETE("/:id", companyHandler.DeleteCompany)
	company.PATCH("/:id", companyHandler.UpdateCompany)

	err := ginRoutes.SERVE()

	if err != nil {
		return
	}

}
