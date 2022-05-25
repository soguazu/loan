package common

import (
	uuid "github.com/satori/go.uuid"
)

// UpdateAddressRequest DTO to update address
type UpdateAddressRequest struct {
	Address            *string `json:"address,omitempty"`
	ApartmentUnitFloor *int32  `json:"apartment_unit_floor,omitempty"`
	City               *string `json:"city,omitempty"`
	State              *string `json:"state,omitempty"`
	Country            *string `json:"country,omitempty"`
	UtilityBill        *string `json:"utility_bill,omitempty"`
	PostalCode         *string `json:"postal_code,omitempty"`
}

// GetAddressResponse DTO
type GetAddressResponse struct {
	ID                 uuid.UUID `json:"id,omitempty"`
	Company            uuid.UUID `json:"company,omitempty"`
	Address            string    `json:"address"`
	ApartmentUnitFloor int32     `json:"apartment_unit_floor"`
	City               string    `json:"city"`
	State              string    `json:"state"`
	Country            string    `json:"country"`
	UtilityBill        string    `json:"utility_bill"`
}

// CreateAddressResponse DTO
type CreateAddressResponse struct {
	ID                 uuid.UUID `json:"id,omitempty"`
	Company            uuid.UUID `json:"company,omitempty"`
	Address            string    `json:"address"`
	ApartmentUnitFloor int32     `json:"apartment_unit_floor"`
	City               string    `json:"city"`
	State              string    `json:"state"`
	Country            string    `json:"country"`
	UtilityBill        string    `json:"utility_bill"`
}

// CreateDataResponse returns address response
type CreateDataResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Data    CreateAddressResponse `json:"data"`
}

// FilterAddressDataResponse returns address responses
type FilterAddressDataResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Data    []CreateAddressResponse `json:"data"`
}

// CreateAddressRequest DTO to create address
type CreateAddressRequest struct {
	Company            uuid.UUID `json:"company" binding:"required"`
	Address            string    `json:"address" binding:"required"`
	ApartmentUnitFloor int32     `json:"apartment_unit_floor"`
	City               string    `json:"city" binding:"required"`
	State              string    `json:"state" binding:"required"`
	Country            string    `json:"country" binding:"required"`
	UtilityBill        string    `json:"utility_bill" binding:"required"`
	PostalCode         string    `json:"postal_code" binding:"required"`
}
