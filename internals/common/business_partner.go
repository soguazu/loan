package common

import (
	uuid "github.com/satori/go.uuid"
)

// UpdateBusinessPartnerRequest DTO to update business partner
type UpdateBusinessPartnerRequest struct {
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty" binding:"e164"`
}

// GetBusinessPartnerResponse DTO
type GetBusinessPartnerResponse struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Company uuid.UUID `json:"company,omitempty"`
	Name    string    `json:"name"`
	Phone   string    `json:"phone"`
}

// BusinessPartnerDataResponse returns business partner response
type BusinessPartnerDataResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Data    GetBusinessPartnerResponse `json:"data"`
}

// FilterBusinessPartnerDataResponse returns business partner responses
type FilterBusinessPartnerDataResponse struct {
	Success bool                         `json:"success"`
	Message string                       `json:"message"`
	Data    []GetBusinessPartnerResponse `json:"data"`
}

// CreateBusinessPartnerRequest DTO to create business partner
type CreateBusinessPartnerRequest struct {
	Company uuid.UUID `json:"company" binding:"required"`
	Name    string    `json:"name" binding:"required"`
	Phone   string    `json:"phone" binding:"required,e164"`
}
