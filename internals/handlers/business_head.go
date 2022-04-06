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

type businessHeadHandler struct {
	BusinessHeadService ports.IBusinessHeadService
	logger              *log.Logger
	handlerName         string
}

// NewBusinessHeadHandler function creates a new instance for business head handler
func NewBusinessHeadHandler(bs ports.IBusinessHeadService, l *log.Logger, n string) ports.IBusinessHeadHandler {
	return &businessHeadHandler{
		BusinessHeadService: bs,
		logger:              l,
		handlerName:         n,
	}
}

// GetBusinessHeadByID godoc
// @Summary      Get a business head
// @Description  get business head by ID
// @Tags         business_head
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "BusinessHead ID"
// @Success      200  {object}  common.BusinessHeadDataResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /business_head/{id} [get]
func (bh *businessHeadHandler) GetBusinessHeadByID(c *gin.Context) {
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		bh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	businessHead, err := bh.BusinessHeadService.GetBusinessHeadByID(params.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			bh.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		bh.logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(businessHead, message.GetResponseMessage(bh.handlerName, types.OKAY)))
}

// GetBusinessHeadByCompanyID godoc
// @Summary      Get a business head
// @Description  get business head by ID
// @Tags         business_head
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Company ID"
// @Success      200  {object}  common.FilterBusinessHeadDataResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /business_head/company/{id} [get]
func (bh *businessHeadHandler) GetBusinessHeadByCompanyID(c *gin.Context) {
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		bh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	companyID, err := uuid.FromString(params.ID)

	if err != nil {
		bh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	filter := domain.BusinessHead{
		Company: companyID,
	}

	businessHead, err := bh.BusinessHeadService.GetBusinessHeadBy(filter)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			bh.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		bh.logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(businessHead, message.GetResponseMessage(bh.handlerName, types.OKAY)))
}

// GetAllBusinessHead godoc
// @Summary      Get all business head
// @Description  gets all business head
// @Tags         business_head
// @Accept       json
// @Produce      json
// @Param        limit   query  int  false  "Page size"
// @Param        page   query  int  false  "Page no"
// @Param        sort   query  string  false  "Sort by"
// @Success      200  {object}  common.GetAllResponse
// @Failure      500  {object}  common.Error
// @Router       /business_head [get]
func (bh *businessHeadHandler) GetAllBusinessHead(c *gin.Context) {
	var query utils.Pagination
	if err := c.ShouldBindQuery(&query); err != nil {
		bh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	businessHead, err := bh.BusinessHeadService.GetAllBusinessHead(&query)
	if err != nil {
		bh.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(businessHead, message.GetResponseMessage(bh.handlerName, types.OKAY)))
}

// CreateBusinessHead godoc
// @Summary      Create Business Head
// @Description  creates a business head
// @Tags         business_head
// @Accept       json
// @Produce      json
// @Param company body common.CreateBusinessHeadRequest true "Add company"
// @Success      200  {object}  common.CreateDataResponse
// @Failure      400  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /business_head [post]
func (bh *businessHeadHandler) CreateBusinessHead(c *gin.Context) {
	var body common.CreateBusinessHeadRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		bh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	businessHead := &domain.BusinessHead{
		Company:                body.Company,
		JobTitle:               body.JobTitle,
		Phone:                  body.Phone,
		IdentificationType:     body.IdentificationType,
		IdentificationNumber:   body.IdentificationNumber,
		IdentificationImageURL: body.IdentificationImageURL,
		CompanyIDUrl:           body.CompanyIDUrl,
	}
	err := bh.BusinessHeadService.CreateBusinessHead(businessHead)

	if err != nil {
		bh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.ReturnSuccessResult(businessHead, message.GetResponseMessage(bh.handlerName, types.CREATED)))
}

// DeleteBusinessHead godoc
// @Summary      Delete a business head by ID
// @Description  deletes business head by id
// @Tags         business_head
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Business Head ID"
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /business_head/{id} [delete]
func (bh *businessHeadHandler) DeleteBusinessHead(c *gin.Context) {
	var query common.GetByIDRequest
	if err := c.ShouldBindUri(&query); err != nil {
		bh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	err := bh.BusinessHeadService.DeleteBusinessHead(query.ID)
	if err != nil {
		bh.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusNoContent, result.ReturnSuccessMessage(types.DELETED))
}

// UpdateBusinessHead godoc
// @Summary      Update a business head by ID
// @Description  updates business head by id
// @Tags         business_head
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Business Head ID"
// @Param company body common.UpdateBusinessHeadRequest true "Add address"
// @Success      200  {object}  common.GetBusinessHeadResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /business_head/{id} [patch]
func (bh *businessHeadHandler) UpdateBusinessHead(c *gin.Context) {
	var body common.UpdateBusinessHeadRequest
	var params common.GetByIDRequest

	if err := c.ShouldBindUri(&params); err != nil {
		bh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		bh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	businessHead, err := bh.BusinessHeadService.UpdateBusinessHead(params.ID, body)
	if err != nil {
		bh.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(businessHead, message.GetResponseMessage(bh.handlerName, types.UPDATED)))
}
