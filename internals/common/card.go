package common

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// SpendingControls set spending limit
type SpendingControls struct {
	AllowedCategories []string         `json:"allowedCategories"`
	BlockedCategories []string         `json:"blockedCategories"`
	Channels          Channels         `json:"channels"`
	SpendingLimits    []SpendingLimits `json:"spendingLimits"`
}

// Channels available means of transaction
type Channels struct {
	Atm    bool `json:"atm"`
	Web    bool `json:"web"`
	Pos    bool `json:"pos"`
	Mobile bool `json:"mobile"`
}

//SpendingLimits set transaction limit on intervals
type SpendingLimits struct {
	Amount   int    `json:"amount"`
	Interval string `json:"interval"`
}

// User DTO
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

// CreateCardRequest DTO
type CreateCardRequest struct {
	Name             string           `json:"card_name" binding:"required"`
	Company          uuid.UUID        `json:"company,omitempty" binding:"required"`
	Type             string           `json:"type" binding:"required" form:"default:virtual"`
	Brand            string           `json:"brand" binding:"required" form:"default:verve"`
	Status           string           `json:"status" form:"default:'active'"`
	Summary          string           `json:"summary"`
	User             User             `json:"user" binding:"required"`
	SpendingControls SpendingControls `json:"spendingControls" binding:"required"`
}

// CreateSudoCardRequest DTO
type CreateSudoCardRequest struct {
	Type             string           `json:"type"`
	Brand            string           `json:"brand"`
	Number           string           `json:"number,omitempty"`
	Currency         string           `json:"currency"`
	Status           string           `json:"status"`
	CustomerID       string           `json:"customerId"`
	FundingSourceID  string           `json:"fundingSourceId"`
	SpendingControls SpendingControls `json:"spendingControls"`
}

// Company DTO
type Company struct {
	Name string `json:"name"`
}

type Individual struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// BillingAddress DTO
type BillingAddress struct {
	Line1      string `json:"line1"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

// CreateCustomerRequest DTO
type CreateCustomerRequest struct {
	Type           string `json:"type"`
	Company        `json:"company"`
	Individual     `json:"individual"`
	Status         string `json:"status"`
	BillingAddress `json:"billingAddress"`
	Name           string `json:"name"`
	Phone          string `json:"phoneNumber"`
	Email          string `json:"emailAddress"`
}

// CreateSudoCustomerResponse DTO
type CreateSudoCustomerResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    interface{} `json:"message"`
	Data       struct {
		Business    string `json:"business"`
		Type        string `json:"type"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phoneNumber"`
		Status      string `json:"status"`
		Company     struct {
			Name string `json:"name"`
			ID   string `json:"_id"`
		} `json:"company"`
		BillingAddress struct {
			Line1      string `json:"line1"`
			City       string `json:"city"`
			State      string `json:"state"`
			Country    string `json:"country"`
			PostalCode string `json:"postalCode"`
			ID         string `json:"_id"`
		} `json:"billingAddress"`
		IsDeleted bool      `json:"isDeleted"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		ID        string    `json:"_id"`
		V         int       `json:"__v"`
	} `json:"data"`
}

// UpdateSudoCardRequest UPDATE card struct
type UpdateSudoCardRequest struct {
	Status           *string `json:"status,omitempty"`
	SpendingControls struct {
		SpendingLimits struct {
			Amount   *int    `json:"amount,omitempty"`
			Interval *string `json:"interval,omitempty"`
		} `json:"spendingLimits,omitempty"`
		Channels struct {
			Pos *bool `json:"pos,omitempty"`
			Web *bool `json:"web,omitempty"`
			Atm *bool `json:"atm,omitempty"`
		} `json:"channels"`
		AllowedCategories []string `json:"allowedCategories"`
		BlockedCategories []string `json:"blockedCategories"`
	}
}

// ActionOnCardRequest DTO lock card
type ActionOnCardRequest struct {
	Lock bool `json:"lock"`
}

// ChangeCardStatusRequest DTO to change card status
type ChangeCardStatusRequest struct {
	Status string `json:"status" binding:"required"`
	Reason string `json:"reason"`
}

// CancelCardRequest DTO to cancel card
type CancelCardRequest struct {
	Status           string `json:"status"`
	Reason           string `json:"cancellationReason"`
	SpendingControls `json:"spendingControls"`
}

// ChangeCardPinRequest DTO to Change card pin
type ChangeCardPinRequest struct {
	OldPin string `json:"oldPin"`
	NewPin string `json:"newPin"`
}

// CreateSudoCardResponse DTO response for create card
type CreateSudoCardResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    interface{} `json:"message"`
	Error      string      `json:"error"`
	Data       struct {
		Business         string `json:"business"`
		Customer         string `json:"customer"`
		Account          string `json:"account"`
		FundingSource    string `json:"fundingSource"`
		Type             string `json:"type"`
		Brand            string `json:"brand"`
		Currency         string `json:"currency"`
		MaskedPan        string `json:"maskedPan"`
		ExpiryMonth      string `json:"expiryMonth"`
		ExpiryYear       string `json:"expiryYear"`
		Status           string `json:"status"`
		SpendingControls struct {
			Channels struct {
				Atm    bool   `json:"atm"`
				Pos    bool   `json:"pos"`
				Web    bool   `json:"web"`
				Mobile bool   `json:"mobile"`
				ID     string `json:"_id"`
			} `json:"channels"`
			AllowedCategories []string `json:"allowedCategories"`
			BlockedCategories []string `json:"blockedCategories"`
			SpendingLimits    []struct {
				Amount     int           `json:"amount"`
				Interval   string        `json:"interval"`
				Categories []interface{} `json:"categories"`
				ID         string        `json:"_id"`
			} `json:"spendingLimits"`
			ID string `json:"_id"`
		} `json:"spendingControls"`
		Is2FAEnrolled bool      `json:"is2FAEnrolled"`
		IsDeleted     bool      `json:"isDeleted"`
		CreatedAt     time.Time `json:"createdAt"`
		UpdatedAt     time.Time `json:"updatedAt"`
		ID            string    `json:"_id"`
		V             int       `json:"__v"`
	} `json:"data"`
}

// GetCardResponse DTO response for get single card
type GetCardResponse struct {
	ID                uuid.UUID `json:"ID,omitempty"`
	Company           uuid.UUID `json:"company,omitempty"`
	Wallet            uuid.UUID `json:"wallet,omitempty"`
	Name              string    `json:"name"`
	PartnerCustomerID string    `json:"customerId"`
	Type              string    `json:"type"`
	Brand             string    `json:"brand"`
	Number            string    `json:"number"`
	Currency          string    `json:"currency"`
	Status            string    `json:"status"`
	Lock              bool      `json:"lock"`
	PartnerCardID     string    `json:"partner_card_id"`
	Partner           string    `json:"partner"`
	CardAuth          string    `json:"card_auth"`
	Summary           string    `json:"summary"`
	Customer          uuid.UUID `json:"customer"`
	SpendingControls  struct {
		Channels struct {
			Atm    bool `json:"atm"`
			Pos    bool `json:"pos"`
			Web    bool `json:"web"`
			Mobile bool `json:"mobile"`
		} `json:"channels"`
		AllowedCategories []string `json:"allowedCategories"`
		BlockedCategories []string `json:"blockedCategories"`
		SpendingLimits    struct {
			Amount   int    `json:"amount"`
			Interval string `json:"interval"`
		} `json:"spendingLimits"`
		ID string `json:"_id"`
	} `json:"spendingControls" gorm:"type:text"`
	Business    string `json:"business"`
	Account     string `json:"account"`
	MaskedPan   string `json:"maskedPan"`
	ExpiryMonth string `json:"expiryMonth"`
	ExpiryYear  string `json:"expiryYear"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// GetSingleCardResponse DTO get a card
type GetSingleCardResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    GetCardResponse `json:"data"`
}

// GetSinglePANResponse DTO get a pan
type GetSinglePANResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		ID        uuid.UUID `json:"id"`
		Number    string    `json:"number"`
		Status    bool      `json:"status"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt time.Time
	} `json:"data"`
}

type ProcessCardUpdate struct {
	StatusCode int         `json:"statusCode"`
	Message    interface{} `json:"message"`
	Error      string      `json:"error"`
}

type AddPANRequest struct {
	Numbers []string `json:"numbers"`
}

type SudoPANNumber struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       struct {
		Brand       string `json:"brand"`
		Number      string `json:"number"`
		ExpiryMonth string `json:"expiryMonth"`
		ExpiryYear  string `json:"expiryYear"`
		Cvv2        string `json:"cvv2"`
		DefaultPin  string `json:"defaultPin"`
	} `json:"data"`
}
