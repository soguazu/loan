package common

import (
	uuid "github.com/satori/go.uuid"
)

// UpdateBusinessHeadRequest DTO to update business head
type UpdateBusinessHeadRequest struct {
	JobTitle               *string `json:"job_title,omitempty"`
	Phone                  *string `json:"phone,omitempty"`
	IdentificationType     *string `json:"identification_type,omitempty"`
	IdentificationNumber   *string `json:"identification_number,omitempty"`
	IdentificationImageURL *string `json:"identification_image_url,omitempty"`
	CompanyIDUrl           *string `json:"company_id_url,omitempty"`
}

// GetBusinessHeadResponse DTO
type GetBusinessHeadResponse struct {
	ID                     uuid.UUID `json:"id,omitempty"`
	Company                uuid.UUID `json:"company,omitempty"`
	JobTitle               string    `json:"job_title"`
	Phone                  string    `json:"phone"`
	IdentificationType     string    `json:"identification_type"`
	IdentificationNumber   string    `json:"identification_number"`
	IdentificationImageURL string    `json:"identification_image_url"`
	CompanyIDUrl           string    `json:"company_id_url"`
}

// BusinessHeadDataResponse returns business head response
type BusinessHeadDataResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    GetBusinessHeadResponse
}

// FilterBusinessHeadDataResponse returns business heads responses
type FilterBusinessHeadDataResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []GetBusinessHeadResponse
}

// CreateBusinessHeadRequest DTO to create business head
type CreateBusinessHeadRequest struct {
	Company                uuid.UUID `json:"company" binding:"required"`
	JobTitle               string    `json:"job_title" binding:"required"`
	Phone                  string    `json:"phone" binding:"required,e164"`
	IdentificationType     string    `json:"identification_type" binding:"required"`
	IdentificationNumber   string    `json:"identification_number" binding:"required"`
	IdentificationImageURL string    `json:"identification_image_url" binding:"required"`
	CompanyIDUrl           string    `json:"company_id_url" binding:"required"`
}
