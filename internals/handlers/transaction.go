package handlers

import (
	"core_business/internals/common"
	"core_business/internals/common/types"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type transactionHandler struct {
	TransactionService ports.ITransactionService
	logger             *log.Logger
	handlerName        string
}

// NewTransactionHandler function creates a new instance for transaction handler
func NewTransactionHandler(ech ports.ITransactionService, l *log.Logger, n string) ports.ITransactionHandler {
	return &transactionHandler{
		TransactionService: ech,
		logger:             l,
		handlerName:        n,
	}
}

// GetTransactionByID godoc
// @Summary      Get a transaction
// @Description  get transaction by ID
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Transaction ID"
// @Success      200  {object}  common.GetSingleTransactionResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /transaction/{id} [get]
func (th *transactionHandler) GetTransactionByID(c *gin.Context) {
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		th.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	wallet, err := th.TransactionService.GetTransactionByID(params.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			th.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		th.logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(wallet, message.GetResponseMessage(th.handlerName, types.OKAY)))
}

// GetTransactionByCompanyID godoc
// @Summary      Get transactions by company id
// @Description  gets all transactions by company id
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Company ID"
// @Param        limit   query  int  false  "Page size"
// @Param        page   query  int  false  "Page no"
// @Param        sort   query  string  false  "Sort by"
// @Success      200  {object}  common.GetAllResponse
// @Failure      500  {object}  common.Error
// @Router       /transaction/company/{id} [get]
func (th *transactionHandler) GetTransactionByCompanyID(c *gin.Context) {

	var (
		params common.GetByIDRequest
		query  utils.Pagination
	)

	if err := c.ShouldBindUri(&params); err != nil {
		th.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		th.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	transactions, err := th.TransactionService.GetTransactionByCompanyID(params.ID, &query)

	if err != nil {
		th.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(transactions, message.GetResponseMessage(th.handlerName, types.OKAY)))

}

// GetTransactionByCardID godoc
// @Summary      Get transactions by card id
// @Description  gets all transactions by card id
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Card ID"
// @Param        limit   query  int  false  "Page size"
// @Param        page   query  int  false  "Page no"
// @Param        sort   query  string  false  "Sort by"
// @Success      200  {object}  common.GetAllResponse
// @Failure      500  {object}  common.Error
// @Router       /transaction/card/{id} [get]
func (th *transactionHandler) GetTransactionByCardID(c *gin.Context) {
	var (
		params common.GetByIDRequest
		query  utils.Pagination
	)

	if err := c.ShouldBindUri(&params); err != nil {
		th.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		th.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	transactions, err := th.TransactionService.GetTransactionByCompanyID(params.ID, &query)

	if err != nil {
		th.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(transactions, message.GetResponseMessage(th.handlerName, types.OKAY)))
}

// GetAllTransaction godoc
// @Summary      Get transactions
// @Description  gets all transactions
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        limit   query  int  false  "Page size"
// @Param        page   query  int  false  "Page no"
// @Param        sort   query  string  false  "Sort by"
// @Success      200  {object}  common.GetAllResponse
// @Failure      500  {object}  common.Error
// @Router       /transaction [get]
func (th *transactionHandler) GetAllTransaction(c *gin.Context) {
	var query utils.Pagination

	if err := c.ShouldBindQuery(&query); err != nil {
		th.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	transactions, err := th.TransactionService.GetAllTransaction(&query)

	if err != nil {
		th.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(transactions, message.GetResponseMessage(th.handlerName, types.OKAY)))
}

// CreateTransaction godoc
// @Summary      Get transactions
// @Description  gets all transactions
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        limit   query  int  false  "Page size"
// @Param        page   query  int  false  "Page no"
// @Param        sort   query  string  false  "Sort by"
// @Success      200  {object}  common.GetAllResponse
// @Failure      500  {object}  common.Error
// @Router       /transaction/webhook [get]
func (th *transactionHandler) CreateTransaction(c *gin.Context) {
	//var body common.CreateCardRequest
	//if err := c.ShouldBindJSON(&body); err != nil {
	//	th.logger.Error(err)
	//	c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
	//	return
	//}

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error
	}

	fmt.Println(jsonData)

	c.JSON(http.StatusCreated, result.ReturnSuccessResult(domain.Transaction{}, message.GetResponseMessage(th.handlerName, types.CREATED)))

}

// UpdateTransaction godoc
// @Summary      Update a transaction by ID
// @Description  update transaction by id
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Transaction ID"
// @Param company body common.UpdateTransactionRequest true "Update transaction"
// @Success      200  {object}  common.GetSingleTransactionResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /transaction/{id} [patch]
func (th *transactionHandler) UpdateTransaction(c *gin.Context) {
	var (
		body   common.UpdateTransactionRequest
		params common.GetByIDRequest
	)

	if err := c.ShouldBindUri(&params); err != nil {
		th.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		th.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	transaction, err := th.TransactionService.UpdateTransaction(params.ID, body)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			th.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		th.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(transaction, message.GetResponseMessage(th.handlerName, types.UPDATED)))
}

// DeleteTransaction godoc
// @Summary      Delete a transaction by ID
// @Description  Delete transaction by id
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Transaction ID"
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /transaction/{id} [delete]
func (th *transactionHandler) DeleteTransaction(c *gin.Context) {
	var params common.GetByIDRequest

	if err := c.ShouldBindUri(&params); err != nil {
		th.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	err := th.TransactionService.DeleteTransaction(params.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			th.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		th.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessMessage(message.GetResponseMessage(th.handlerName, types.DELETED)))
}

// LockTransaction godoc
// @Summary      Lock a transaction by ID
// @Description  Lock transaction by id
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Transaction ID"
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /transaction/{id}/lock [patch]
func (th *transactionHandler) LockTransaction(c *gin.Context) {
	var params common.GetByIDRequest

	if err := c.ShouldBindUri(&params); err != nil {
		th.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	transaction, err := th.TransactionService.LockTransaction(params.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			th.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		th.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(transaction, message.GetResponseMessage(th.handlerName, types.OKAY)))
}