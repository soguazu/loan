package handlers

import (
	"core_business/internals/common"
	"core_business/internals/common/types"
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type expenseCategoryHandler struct {
	ExpenseCategoryService ports.IExpenseCategoryService
	logger                 *log.Logger
	handlerName            string
}

// NewExpenseCategoryHandler function creates a new instance for expense category handler
func NewExpenseCategoryHandler(ech ports.IExpenseCategoryService, l *log.Logger, n string) ports.IExpenseCategoryHandler {
	return &expenseCategoryHandler{
		ExpenseCategoryService: ech,
		logger:                 l,
		handlerName:            n,
	}
}

// GetExpenseCategoryByID godoc
// @Summary      Get an expense category
// @Description  get expense category by ID
// @Tags         expense_category
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Expense Category ID"
// @Success      200  {object}  common.GetExpenseCategoryResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /expense_category/{id} [get]
func (ech *expenseCategoryHandler) GetExpenseCategoryByID(c *gin.Context) {
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		ech.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	company, err := ech.ExpenseCategoryService.GetExpenseCategoryByID(params.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ech.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		ech.logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(company, message.GetResponseMessage(ech.handlerName, types.OKAY)))
}

// GetAllExpenseCategory godoc
// @Summary      Get all expense category
// @Description  gets all expense category
// @Tags         expense_category
// @Accept       json
// @Produce      json
// @Param        limit   query  int  false  "Page size"
// @Param        page   query  int  false  "Page no"
// @Param        sort   query  string  false  "Sort by"
// @Success      200  {object}  common.GetAllResponse
// @Failure      500  {object}  common.Error
// @Router       /expense_category [get]
func (ech *expenseCategoryHandler) GetAllExpenseCategory(c *gin.Context) {
	var query utils.Pagination
	if err := c.ShouldBindQuery(&query); err != nil {
		ech.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	expenseCategory, err := ech.ExpenseCategoryService.GetAllExpenseCategory(&query)
	if err != nil {
		ech.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(expenseCategory, message.GetResponseMessage(ech.handlerName, types.OKAY)))
}

// CreateExpenseCategory godoc
// @Summary      Create expense category
// @Description  creates an expense category
// @Tags         expense_category
// @Accept       json
// @Produce      json
// @Param company body common.CreateExpenseCategoryRequest true "Add category"
// @Success      200  {object}  common.CreateDataResponse
// @Failure      400  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /expense_category [post]
func (ech *expenseCategoryHandler) CreateExpenseCategory(c *gin.Context) {
	var body common.CreateExpenseCategoryRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		ech.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	expenseCategory := &domain.ExpenseCategory{
		Title: body.Title,
	}

	err := ech.ExpenseCategoryService.CreateExpenseCategory(expenseCategory)

	if err != nil {
		ech.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.ReturnSuccessResult(expenseCategory, message.GetResponseMessage(ech.handlerName, types.CREATED)))
}

// DeleteExpenseCategory godoc
// @Summary      Delete an expense category by ID
// @Description  deletes expense category by id
// @Tags         expense_category
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Expense category  ID"
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /expense_category/{id} [delete]
func (ech *expenseCategoryHandler) DeleteExpenseCategory(c *gin.Context) {
	var query common.GetByIDRequest
	if err := c.ShouldBindUri(&query); err != nil {
		ech.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	err := ech.ExpenseCategoryService.DeleteExpenseCategory(query.ID)
	if err != nil {
		ech.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusNoContent, result.ReturnSuccessMessage(types.DELETED))
}

// UpdateExpenseCategory godoc
// @Summary      Update an expense category by ID
// @Description  update expense category by id
// @Tags         expense_category
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Expense category ID"
// @Param company body common.UpdateExpenseCategoryRequest true "Add address"
// @Success      200  {object}  common.GetExpenseCategoryDataResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /expense_category/{id} [patch]
func (ech *expenseCategoryHandler) UpdateExpenseCategory(c *gin.Context) {
	var body common.UpdateExpenseCategoryRequest
	var params common.GetByIDRequest

	if err := c.ShouldBindUri(&params); err != nil {
		ech.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		ech.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	expenseCategory, err := ech.ExpenseCategoryService.UpdateExpenseCategory(params.ID, body)
	if err != nil {
		ech.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(expenseCategory, message.GetResponseMessage(ech.handlerName, types.UPDATED)))
}
