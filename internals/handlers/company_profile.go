package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"core_business/internals/common"
	"core_business/internals/common/types"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	"gorm.io/gorm"
	"net/http"
)

type companyProfileHandler struct {
	CompanyProfileService ports.ICompanyProfileService
	logger                *log.Logger
	handlerName           string
}

// NewCompanyProfileHandler function creates a new instance for company profile handler
func NewCompanyProfileHandler(cs ports.ICompanyProfileService, l *log.Logger, n string) ports.ICompanyProfileHandler {
	return &companyProfileHandler{
		CompanyProfileService: cs,
		logger:                l,
		handlerName:           n,
	}
}

// GetCompanyProfileByID godoc
// @Summary      Get a company profile
// @Description  get company profile by ID
// @Tags         company_profile
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "CompanyProfile ID"
// @Success      200  {object}  common.CompanyProfileDataResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /company_profile/{id} [get]
func (ch *companyProfileHandler) GetCompanyProfileByID(c *gin.Context) {
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	companyProfile, err := ch.CompanyProfileService.GetCompanyProfileByID(params.ID)
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

	c.JSON(http.StatusOK, result.ReturnSuccessResult(companyProfile, message.GetResponseMessage(ch.handlerName, types.OKAY)))
}

// GetCompanyProfileByCompanyID godoc
// @Summary      Get a company profile
// @Description  get company profile by ID
// @Tags         company_profile
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Company ID"
// @Success      200  {object}  common.FilterCompanyProfileDataResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /company_profile/company/{id} [get]
func (ch *companyProfileHandler) GetCompanyProfileByCompanyID(c *gin.Context) {
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	companyID, err := uuid.FromString(params.ID)

	if err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	filter := domain.CompanyProfile{
		Company: companyID,
	}

	companyProfile, err := ch.CompanyProfileService.GetCompanyProfileBy(filter)
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

	c.JSON(http.StatusOK, result.ReturnSuccessResult(companyProfile, message.GetResponseMessage(ch.handlerName, types.OKAY)))
}

// GetAllCompanyProfile godoc
// @Summary      Get all company profile
// @Description  gets all company profile
// @Tags         company_profile
// @Accept       json
// @Produce      json
// @Param        limit   query  int  false  "Page size"
// @Param        page   query  int  false  "Page no"
// @Param        sort   query  string  false  "Sort by"
// @Success      200  {object}  common.GetAllResponse
// @Failure      500  {object}  common.Error
// @Router       /company_profile [get]
func (ch *companyProfileHandler) GetAllCompanyProfile(c *gin.Context) {
	var query utils.Pagination
	if err := c.ShouldBindQuery(&query); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	companyProfile, err := ch.CompanyProfileService.GetAllCompanyProfile(&query)
	if err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(companyProfile, message.GetResponseMessage(ch.handlerName, types.OKAY)))
}

// CreateCompanyProfile godoc
// @Summary      Create Business Head
// @Description  creates a business head
// @Tags         business_head
// @Accept       json
// @Produce      json
// @Param company body common.CreateBusinessHeadRequest true "Add company"
// @Success      200  {object}  common.CreateDataResponse
// @Failure      400  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /company_profile [post]
func (ch *companyProfileHandler) CreateCompanyProfile(c *gin.Context) {
	var body common.CreateCompanyProfileRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	companyProfile := &domain.CompanyProfile{
		Company:            body.Company,
		RCNumber:           body.RCNumber,
		BusinessTin:        body.BusinessTin,
		BusinessType:       body.BusinessType,
		BusinessActivity:   body.BusinessActivity,
		CACCertificateURL:  body.CACCertificateURL,
		MermatURL:          body.MermatURL,
		StatusReportURL:    body.StatusReportURL,
		IncorporationYear:  body.IncorporationYear,
		IncorporationState: body.IncorporationState,
	}
	err := ch.CompanyProfileService.CreateCompanyProfile(companyProfile)

	if err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.ReturnSuccessResult(companyProfile, message.GetResponseMessage(ch.handlerName, types.CREATED)))
}

// DeleteCompanyProfile godoc
// @Summary      Delete a company profile by ID
// @Description  deletes company profile by id
// @Tags         company_profile
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "CompanyProfile ID"
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /company_profile/{id} [delete]
func (ch *companyProfileHandler) DeleteCompanyProfile(c *gin.Context) {
	var query common.GetByIDRequest
	if err := c.ShouldBindUri(&query); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	err := ch.CompanyProfileService.DeleteCompanyProfile(query.ID)
	if err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusNoContent, result.ReturnSuccessMessage(types.DELETED))
}

// UpdateCompanyProfile godoc
// @Summary      Update a company profile by ID
// @Description  updates company profile by id
// @Tags         company_profile
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "CompanyProfile ID"
// @Param company body common.UpdateCompanyProfileRequest true "Add CompanyProfile"
// @Success      200  {object}  common.GetBusinessHeadResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /company_profile/{id} [patch]
func (ch *companyProfileHandler) UpdateCompanyProfile(c *gin.Context) {
	var body common.UpdateCompanyProfileRequest
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

	companyProfile, err := ch.CompanyProfileService.UpdateCompanyProfile(params.ID, body)
	if err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(companyProfile, message.GetResponseMessage(ch.handlerName, types.UPDATED)))
}
