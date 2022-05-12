package common

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// SpendingControls set spending limit
type SpendingControls struct {
	AllowedCategories []string `json:"allowedCategories"`
	BlockedCategories []string `json:"blockedCategories"`
	Channels          `json:"channels"`
	SpendingLimits    `json:"spendingLimits"`
}

// Channels available means of transaction
type Channels struct {
	Atm    bool `json:"atm" binding:"default:true"`
	Web    bool `json:"web" binding:"default:true"`
	Pos    bool `json:"pos" binding:"default:true"`
	Mobile bool `json:"mobile" binding:"default:false"`
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
	ID               uuid.UUID `json:"id,omitempty" binding:"required"`
	Name             string    `json:"card_name" binding:"required"`
	Company          uuid.UUID `json:"company,omitempty" binding:"required"`
	Type             string    `json:"type" binding:"required;default:virtual"`
	Brand            string    `json:"brand" binding:"required;default:verve"`
	Number           string    `json:"number"`
	Currency         string    `json:"currency" binding:"default:'NG'"`
	Status           string    `json:"status" binding:"default:'active'"`
	Lock             bool      `json:"lock" binding:"default:false"`
	CardAuth         string    `json:"card_auth" binding:"required;default:1234"`
	Summary          string    `json:"summary"`
	User             `json:"user" binding:"required"`
	SpendingControls `json:"spendingControls"`
}

// CreateSudoCardRequest DTO
type CreateSudoCardRequest struct {
	Type             string `json:"type"`
	Name             string `json:"name"`
	Brand            string `json:"brand"`
	Number           string `json:"number,omitempty"`
	Currency         string `json:"currency"`
	Status           string `json:"status"`
	CustomerID       string `json:"customerId"`
	FundingSourceID  string `json:"fundingSourceId"`
	SpendingControls `json:"spendingControls"`
}

// Company DTO
type Company struct {
	Name string `json:"name"`
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
	Status         string `json:"status;"`
	BillingAddress `json:"billingAddress"`
	Name           string `json:"name"`
	Phone          string `json:"phoneNumber"`
	Email          string `json:"emailAddress"`
}

// CreateSudoCustomerResponse DTO
type CreateSudoCustomerResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
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
	Lock bool `json:"status" binding:"required"`
}

// ChangeCardStatusRequest DTO to change card status
type ChangeCardStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

// CancelCardRequest DTO to cancel card
type CancelCardRequest struct {
	Status           string `json:"status"`
	SpendingControls `json:"spendingControls"`
}

// ChangeCardPinRequest DTO to Change card pin
type ChangeCardPinRequest struct {
	OldPin string `json:"oldPin"`
	NewPin string `json:"newPin"`
}

// CreateSudoCardResponse DTO response for create card
type CreateSudoCardResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
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