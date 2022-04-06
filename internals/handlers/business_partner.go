package handlers

import (
	"core_business/internals/common"
	"core_business/internals/common/types"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type businessPartnerHandler struct {
	BusinessPartnerService ports.IBusinessPartnerService
	logger                 *log.Logger
	handlerName            string
}

// NewBusinessPartnerHandler function creates a new instance for business partner handler
func NewBusinessPartnerHandler(bp ports.IBusinessPartnerService, l *log.Logger, n string) ports.IBusinessPartnerHandler {
	return &businessPartnerHandler{
		BusinessPartnerService: bp,
		logger:                 l,
		handlerName:            n,
	}
}

// GetBusinessPartnerByID godoc
// @Summary      Get a business partner
// @Description  get business partner by ID
// @Tags         business_partner
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "BusinessPartner ID"
// @Success      200  {object}  common.BusinessPartnerDataResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /business_partner/{id} [get]
func (bp *businessPartnerHandler) GetBusinessPartnerByID(c *gin.Context) {
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		bp.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	businessPartner, err := bp.BusinessPartnerService.GetBusinessPartnerByID(params.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			bp.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		bp.logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(businessPartner, message.GetResponseMessage(bp.handlerName, types.OKAY)))
}

// GetBusinessPartnerByCompanyID godoc
// @Summary      Get a business head
// @Description  get business head by ID
// @Tags         business_head
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Company ID"
// @Success      200  {object}  common.FilterBusinessPartnerDataResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /business_head/company/{id} [get]
func (bp *businessPartnerHandler) GetBusinessPartnerByCompanyID(c *gin.Context) {
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		bp.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	companyID, err := uuid.FromString(params.ID)

	if err != nil {
		bp.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	filter := domain.BusinessPartner{
		Company: companyID,
	}

	businessPartner, err := bp.BusinessPartnerService.GetBusinessPartnerBy(filter)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			bp.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		bp.logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(businessPartner, message.GetResponseMessage(bp.handlerName, types.OKAY)))
}

// GetAllBusinessPartner godoc
// @Summary      Get all business partner
// @Description  gets all business partner
// @Tags         business_partner
// @Accept       json
// @Produce      json
// @Param        limit   query  int  false  "Page size"
// @Param        page   query  int  false  "Page no"
// @Param        sort   query  string  false  "Sort by"
// @Success      200  {object}  common.GetAllResponse
// @Failure      500  {object}  common.Error
// @Router       /business_partner [get]
func (bp *businessPartnerHandler) GetAllBusinessPartner(c *gin.Context) {
	var query utils.Pagination
	if err := c.ShouldBindQuery(&query); err != nil {
		bp.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	businessPartner, err := bp.BusinessPartnerService.GetAllBusinessPartner(&query)
	if err != nil {
		bp.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(businessPartner, message.GetResponseMessage(bp.handlerName, types.OKAY)))
}

// CreateBusinessPartner godoc
// @Summary      Create Business Partner
// @Description  creates a business partner
// @Tags         business_partner
// @Accept       json
// @Produce      json
// @Param company body common.CreateBusinessPartnerRequest true "Add business partner"
// @Success      200  {object}  common.BusinessPartnerDataResponse
// @Failure      400  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /business_partner [post]
func (bp *businessPartnerHandler) CreateBusinessPartner(c *gin.Context) {
	var body common.CreateBusinessPartnerRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		bp.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	businessPartner := &domain.BusinessPartner{
		Company: body.Company,
		Name:    body.Name,
		Phone:   body.Phone,
	}
	err := bp.BusinessPartnerService.CreateBusinessPartner(businessPartner)

	if err != nil {
		bp.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.ReturnSuccessResult(businessPartner, message.GetResponseMessage(bp.handlerName, types.CREATED)))
}

// DeleteBusinessPartner godoc
// @Summary      Delete a business partner by ID
// @Description  deletes business partner by id
// @Tags         business_partner
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Business Partner ID"
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /business_partner/{id} [delete]
func (bp *businessPartnerHandler) DeleteBusinessPartner(c *gin.Context) {
	var query common.GetByIDRequest
	if err := c.ShouldBindUri(&query); err != nil {
		bp.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	err := bp.BusinessPartnerService.DeleteBusinessPartner(query.ID)
	if err != nil {
		bp.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusNoContent, result.ReturnSuccessMessage(types.DELETED))
}

// UpdateBusinessPartner godoc
// @Summary      Update a business partner by ID
// @Description  updates business partner by id
// @Tags         business_partner
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Business Partner ID"
// @Param company body common.UpdateBusinessPartnerRequest true "Add business partner"
// @Success      200  {object}  common.BusinessPartnerDataResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /business_partner/{id} [patch]
func (bp *businessPartnerHandler) UpdateBusinessPartner(c *gin.Context) {
	var body common.UpdateBusinessPartnerRequest
	var params common.GetByIDRequest

	if err := c.ShouldBindUri(&params); err != nil {
		bp.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		bp.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	businessHead, err := bp.BusinessPartnerService.UpdateBusinessPartner(params.ID, body)
	if err != nil {
		bp.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(businessHead, message.GetResponseMessage(bp.handlerName, types.UPDATED)))
}
