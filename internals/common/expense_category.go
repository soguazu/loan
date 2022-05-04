package common

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// UpdateExpenseCategoryRequest DTO to update wallet
type UpdateExpenseCategoryRequest struct {
	Title *string `json:"title,omitempty"`
}

// CreateExpenseCategoryRequest DTO to update wallet
type CreateExpenseCategoryRequest struct {
	Title string `json:"title,omitempty"`
}

// GetExpenseCategoryResponse DTO
type GetExpenseCategoryResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// CreateExpenseCategoryResponse DTO
type CreateExpenseCategoryResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// CreateExpenseCategoryDataResponse returns wallet response
type CreateExpenseCategoryDataResponse struct {
	Success bool                          `json:"success"`
	Message string                        `json:"message"`
	Data    CreateExpenseCategoryResponse `json:"data"`
}

// GetExpenseCategoryDataResponse returns wallet response
type GetExpenseCategoryDataResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Data    GetExpenseCategoryResponse `json:"data"`
}
