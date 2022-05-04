package handlers

import (
	"core_business/internals/common"
	"core_business/internals/common/types"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type walletHandler struct {
	WalletService ports.IWalletService
	logger        *log.Logger
	handlerName   string
}

// NewWalletHandler function creates a new instance for wallet handler
func NewWalletHandler(ws ports.IWalletService, l *log.Logger, n string) ports.IWalletHandler {
	return &walletHandler{
		WalletService: ws,
		logger:        l,
		handlerName:   n,
	}
}

// GetWalletByID godoc
// @Summary      Get a wallet
// @Description  get wallet by ID
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Wallet ID"
// @Success      200  {object}  common.GetWalletResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /wallet/{id} [get]
func (wh *walletHandler) GetWalletByID(c *gin.Context) {
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	wallet, err := wh.WalletService.GetWalletByID(params.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			wh.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		wh.logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(wallet, message.GetResponseMessage(wh.handlerName, types.OKAY)))
}

// CreateWallet godoc
// @Summary      Create Wallet
// @Description  creates a wallet
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Param wallet body common.CreateWalletRequest true "Add wallet"
// @Success      200  {object}  common.CreateWalletDataResponse
// @Failure      400  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /wallet/webhook [post]
func (wh *walletHandler) CreateWallet(c *gin.Context) {
	type ModelX[T any] struct {
		Data T
	}

	var body common.CreateWalletRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	wallet := &domain.Wallet{
		Company:         body.Company,
		AccountID:       body.AccountID,
		CustomerID:      body.CustomerID,
		PreviousBalance: 0,
		CurrentSpending: 0,
		AvailableCredit: 0,
		CreditLimit:     0,
		CashBackPayment: 0,
		TotalBalance:    0,
		Status:          true,
	}
	err := wh.WalletService.CreateWallet(wallet)

	if err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.ReturnSuccessResult(wallet, message.GetResponseMessage(wh.handlerName, types.CREATED)))
}

// DeleteWallet godoc
// @Summary      Delete a wallet by ID
// @Description  deletes wallet by id
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Wallet ID"
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /wallet/{id} [delete]
func (wh *walletHandler) DeleteWallet(c *gin.Context) {
	var query common.GetByIDRequest
	if err := c.ShouldBindUri(&query); err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	err := wh.WalletService.DeleteWallet(query.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			wh.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusNoContent, result.ReturnSuccessMessage(types.DELETED))
}

// UpdateWallet godoc
// @Summary      Update a wallet by ID
// @Description  update wallet by id
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Wallet ID"
// @Param wallet body common.UpdateWalletRequest true "Update Wallet"
// @Success      200  {object}  common.GetWalletDataResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /wallet/{id} [patch]
func (wh *walletHandler) UpdateWallet(c *gin.Context) {
	var body common.UpdateWalletRequest
	var params common.GetByIDRequest

	if err := c.ShouldBindUri(&params); err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	wallet, err := wh.WalletService.UpdateWallet(params.ID, body)
	if err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(wallet, message.GetResponseMessage(wh.handlerName, types.UPDATED)))
}
