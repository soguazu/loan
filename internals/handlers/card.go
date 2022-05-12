package handlers

import (
	"core_business/internals/common"
	"core_business/internals/common/types"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type cardHandler struct {
	CardService ports.ICardService
	logger      *log.Logger
	handlerName string
}

// NewCardHandler function creates a new instance for card handler
func NewCardHandler(cs ports.ICardService, l *log.Logger, n string) ports.ICardHandler {
	return &cardHandler{
		CardService: cs,
		logger:      l,
		handlerName: n,
	}
}

// GetCardByID godoc
// @Summary      Get a card
// @Description  get card by ID
// @Tags         card
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Card ID"
// @Success      200  {object}  common.GetSingleCardResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /card/{id} [get]
func (ch *cardHandler) GetCardByID(c *gin.Context) {
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	card, err := ch.CardService.GetCardByID(params.ID)
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

	c.JSON(http.StatusOK, result.ReturnSuccessResult(card, message.GetResponseMessage(ch.handlerName, types.OKAY)))
}

// GetCardByCompanyID godoc
// @Summary      Get all cards
// @Description  gets all cards
// @Tags         card
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Company ID"
// @Param        limit   query  int  false  "Page size"
// @Param        page   query  int  false  "Page no"
// @Param        sort   query  string  false  "Sort by"
// @Param        filter   query  string  false  "Sort by"
// @Success      200  {object}  common.GetAllResponse
// @Failure      500  {object}  common.Error
// @Router       /card/company/{id} [get]
func (ch *cardHandler) GetCardByCompanyID(c *gin.Context) {
	var (
		params common.GetByIDRequest
		query  utils.Pagination
	)

	if err := c.ShouldBindUri(&params); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	cards, err := ch.CardService.GetCardByCompanyID(params.ID, &query)

	if err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(cards, message.GetResponseMessage(ch.handlerName, types.OKAY)))

}

// GetAllCard godoc
// @Summary      Get all cards
// @Description  gets all cards
// @Tags         card
// @Accept       json
// @Produce      json
// @Param        limit   query  int  false  "Page size"
// @Param        page   query  int  false  "Page no"
// @Param        sort   query  string  false  "Sort by"
// @Success      200  {object}  common.GetAllResponse
// @Failure      500  {object}  common.Error
// @Router       /card [get]
func (ch *cardHandler) GetAllCard(c *gin.Context) {
	var query utils.Pagination

	if err := c.ShouldBindQuery(&query); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	cards, err := ch.CardService.GetAllCard(&query)

	if err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(cards, message.GetResponseMessage(ch.handlerName, types.OKAY)))
}

// CreateCard godoc
// @Summary      Create Card
// @Description  creates a card
// @Tags         card
// @Accept       json
// @Produce      json
// @Param card body common.CreateCardRequest true "Add card"
// @Success      200  {object}  common.GetSingleCardResponse
// @Failure      400  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /card [post]
func (ch *cardHandler) CreateCard(c *gin.Context) {
	var body common.CreateCardRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	card, err := ch.CardService.CreateCard(body)

	if err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.ReturnSuccessResult(card, message.GetResponseMessage(ch.handlerName, types.CREATED)))
}

// UpdateCard godoc
// @Summary      Update a card by ID
// @Description  update card by id
// @Tags         card
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Card ID"
// @Param company body common.UpdateSudoCardRequest true "Add address"
// @Success      200  {object}  common.GetSingleCardResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /card/{id} [patch]
func (ch *cardHandler) UpdateCard(c *gin.Context) {
	var body common.UpdateSudoCardRequest
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

	card, err := ch.CardService.UpdateCard(params.ID, body)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ch.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		ch.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(card, message.GetResponseMessage(ch.handlerName, types.UPDATED)))
}

// CancelCard godoc
// @Summary      Cancel a card by ID
// @Description  Cancel card by id
// @Tags         card
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Card ID"
// @Param card body common.ChangeCardStatusRequest true "Cancel card"
// @Success      200  {object}  common.GetSingleCardResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /card/{id}/cancel [patch]
func (ch *cardHandler) CancelCard(c *gin.Context) {
	var body common.ChangeCardStatusRequest
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

	err := ch.CardService.CancelCard(params.ID, body)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ch.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		ch.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessMessage(message.GetResponseMessage(ch.handlerName, types.DELETED)))
}

// LockCard godoc
// @Summary      Lock a card by ID
// @Description  Lock card by id
// @Tags         card
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Card ID"
// @Param card body common.ActionOnCardRequest true "Lock card"
// @Success      200  {object}  common.GetSingleCardResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /card/{id}/lock [patch]
func (ch *cardHandler) LockCard(c *gin.Context) {
	var body common.ActionOnCardRequest
	var params common.GetByIDRequest

	if err := c.ShouldBindUri(&params); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ch.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	card, err := ch.CardService.LockCard(params.ID, body)
	if err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(card, message.GetResponseMessage(ch.handlerName, types.UPDATED)))
}

// ChangeCardPin godoc
// @Summary      Change a card pin by ID
// @Description  Change card pin by id
// @Tags         card
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Card ID"
// @Param card body common.ChangeCardPinRequest true "Lock card"
// @Success      200  {object}  common.GetSingleCardResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /card/{id}/change-pin [patch]
func (ch *cardHandler) ChangeCardPin(c *gin.Context) {
	var body common.ChangeCardPinRequest
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

	err := ch.CardService.ChangeCardPin(params.ID, body)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ch.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		ch.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessMessage(message.GetResponseMessage(ch.handlerName, types.UPDATED)))
}
