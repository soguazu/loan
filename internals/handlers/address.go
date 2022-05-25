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

type addressHandler struct {
	AddressService ports.IAddressService
	logger         *log.Logger
	handlerName    string
}

// NewAddressHandler function creates a new instance for address handler
func NewAddressHandler(cs ports.IAddressService, l *log.Logger, n string) ports.IAddressHandler {
	return &addressHandler{
		AddressService: cs,
		logger:         l,
		handlerName:    n,
	}
}

// GetAddressByID godoc
// @Summary      Get an address
// @Description  get address by ID
// @Tags         address
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Address ID"
// @Success      200  {object}  common.GetAddressResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /address/{id} [get]
func (ah *addressHandler) GetAddressByID(c *gin.Context) {
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		ah.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	company, err := ah.AddressService.GetAddressByID(params.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ah.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		ah.logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(company, message.GetResponseMessage(ah.handlerName, types.OKAY)))
}

// GetAddressByCompanyID godoc
// @Summary      Get an address
// @Description  get address by ID
// @Tags         address
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Company ID"
// @Success      200  {object}  common.FilterAddressDataResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /address/company/{id} [get]
func (ah *addressHandler) GetAddressByCompanyID(c *gin.Context) {
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		ah.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	companyID, err := uuid.FromString(params.ID)

	if err != nil {
		ah.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	filter := domain.Address{
		Company: companyID,
	}

	businessHead, err := ah.AddressService.GetAddressBy(filter)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ah.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		ah.logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(businessHead, message.GetResponseMessage(ah.handlerName, types.OKAY)))
}

// GetAllAddress godoc
// @Summary      Get all address
// @Description  gets all address
// @Tags         address
// @Accept       json
// @Produce      json
// @Param        limit   query  int  false  "Page size"
// @Param        page   query  int  false  "Page no"
// @Param        sort   query  string  false  "Sort by"
// @Success      200  {object}  common.GetAllResponse
// @Failure      500  {object}  common.Error
// @Router       /address [get]
func (ah *addressHandler) GetAllAddress(c *gin.Context) {
	var query utils.Pagination
	if err := c.ShouldBindQuery(&query); err != nil {
		ah.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	addresses, err := ah.AddressService.GetAllAddress(&query)
	if err != nil {
		ah.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(addresses, message.GetResponseMessage(ah.handlerName, types.OKAY)))
}

// CreateAddress godoc
// @Summary      Create Address
// @Description  creates an address
// @Tags         address
// @Accept       json
// @Produce      json
// @Param company body common.CreateAddressRequest true "Add company"
// @Success      200  {object}  common.CreateDataResponse
// @Failure      400  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /address [post]
func (ah *addressHandler) CreateAddress(c *gin.Context) {
	var body common.CreateAddressRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		ah.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	address := &domain.Address{
		Company:            body.Company,
		Address:            body.Address,
		City:               body.City,
		State:              body.State,
		Country:            body.Country,
		UtilityBill:        body.UtilityBill,
		ApartmentUnitFloor: body.ApartmentUnitFloor,
		PostalCode:         body.PostalCode,
	}
	err := ah.AddressService.CreateAddress(address)

	if err != nil {
		ah.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.ReturnSuccessResult(address, message.GetResponseMessage(ah.handlerName, types.CREATED)))
}

// DeleteAddress godoc
// @Summary      Delete an address by ID
// @Description  deletes address by id
// @Tags         address
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Address ID"
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /address/{id} [delete]
func (ah *addressHandler) DeleteAddress(c *gin.Context) {
	var query common.GetByIDRequest
	if err := c.ShouldBindUri(&query); err != nil {
		ah.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	err := ah.AddressService.DeleteAddress(query.ID)
	if err != nil {
		ah.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusNoContent, result.ReturnSuccessMessage(types.DELETED))
}

// UpdateAddress godoc
// @Summary      Update an address by ID
// @Description  update address by id
// @Tags         address
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Address ID"
// @Param company body common.UpdateAddressRequest true "Add address"
// @Success      200  {object}  common.GetCompanyResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /address/{id} [patch]
func (ah *addressHandler) UpdateAddress(c *gin.Context) {
	var body common.UpdateAddressRequest
	var params common.GetByIDRequest

	if err := c.ShouldBindUri(&params); err != nil {
		ah.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		ah.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	address, err := ah.AddressService.UpdateAddress(params.ID, body)
	if err != nil {
		ah.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(address, message.GetResponseMessage(ah.handlerName, types.UPDATED)))
}
