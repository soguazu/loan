package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/soguazu/core_business/internals/common"
	"github.com/soguazu/core_business/internals/common/types"
	"github.com/soguazu/core_business/internals/core/domain"
	"github.com/soguazu/core_business/internals/core/ports"
	"github.com/soguazu/core_business/pkg/utils"
	"gorm.io/gorm"
	"net/http"
)

type companyHandler struct {
	CompanyService ports.ICompanyService
	logger         *log.Logger
	handlerName    string
}

var (
	result  utils.Result
	message types.Messages
)

// NewCompanyHandler function creates a new instance for company handler
func NewCompanyHandler(cs ports.ICompanyService, l *log.Logger, n string) ports.ICompanyHandler {
	return &companyHandler{
		CompanyService: cs,
		logger:         l,
		handlerName:    n,
	}
}

// GetCompanyByID godoc
// @Summary      Get a company
// @Description  get company by ID
// @Tags         company
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Company ID"
// @Success      200  {object}  common.GetCompanyResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /company/{id} [get]
func (ch *companyHandler) GetCompanyByID(c *gin.Context) {
	var params common.GetCompanyByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	company, err := ch.CompanyService.GetCompanyByID(params.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ch.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		ch.logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(company, message.GetResponseMessage(ch.handlerName, types.OKAY)))
}

// GetAllCompany godoc
// @Summary      Get all company
// @Description  gets all company
// @Tags         company
// @Accept       json
// @Produce      json
// @Param        limit   query  int  false  "Page size"
// @Param        page   query  int  false  "Page no"
// @Param        sort   query  string  false  "Sort by"
// @Success      200  {object}  common.GetAllResponse
// @Failure      500  {object}  common.Error
// @Router       /company [get]
func (ch *companyHandler) GetAllCompany(c *gin.Context) {
	var query utils.Pagination
	if err := c.ShouldBindQuery(&query); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	companies, err := ch.CompanyService.GetAllCompany(&query)
	if err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(companies, message.GetResponseMessage(ch.handlerName, types.OKAY)))
}

// CreateCompany godoc
// @Summary      Create company
// @Description  creates a company
// @Tags         company
// @Accept       json
// @Produce      json
// @Param company body common.CreateCompanyRequest true "Add company"
// @Success      200  {object}  common.GetCompanyResponse
// @Failure      400  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /company [post]
func (ch *companyHandler) CreateCompany(c *gin.Context) {
	var body common.CreateCompanyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		ch.logger.Error(err, "###")
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	company := &domain.Company{
		Owner:         body.Owner,
		Name:          body.Name,
		Website:       body.Website,
		Type:          body.Type,
		FundingSource: body.FundingSource,
		NoOfEmployee:  body.NoOfEmployee,
	}

	err := ch.CompanyService.CreateCompany(company)
	if err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.ReturnSuccessResult(company, message.GetResponseMessage(ch.handlerName, types.CREATED)))
}

// DeleteCompany godoc
// @Summary      Delete a company by ID
// @Description  delete company by id
// @Tags         company
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Company ID"
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /company/{id} [delete]
func (ch *companyHandler) DeleteCompany(c *gin.Context) {
	var query common.GetCompanyByIDRequest
	if err := c.ShouldBindUri(&query); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	err := ch.CompanyService.DeleteCompany(query.ID)
	if err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusNoContent, result.ReturnSuccessMessage(types.DELETED))
}

// UpdateCompany godoc
// @Summary      Update a company by ID
// @Description  update company by id
// @Tags         company
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Company ID"
// @Param company body common.UpdateCompanyRequest true "Add company"
// @Success      200  {object}  common.GetCompanyResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /company/{id} [patch]
func (ch *companyHandler) UpdateCompany(c *gin.Context) {
	var body common.UpdateCompanyRequest
	var params common.GetByIDRequest

	if err := c.ShouldBindUri(&params); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	company, err := ch.CompanyService.UpdateCompany(params, body)
	if err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(company, message.GetResponseMessage(ch.handlerName, types.UPDATED)))
}
