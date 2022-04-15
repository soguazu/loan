package common

import (
	uuid "github.com/satori/go.uuid"
)

// UpdateCompanyProfileRequest DTO to update company files
type UpdateCompanyProfileRequest struct {
	RCNumber           *string `json:"rc_number,omitempty"`
	BusinessTin        *string `json:"business_tin,omitempty"`
	BusinessType       *string `json:"business_type,omitempty"`
	IncorporationYear  *string `json:"incorporation_year,omitempty"`
	IncorporationState *string `json:"incorporation_state,omitempty"`
	YearsInOperation   *int    `json:"years_in_operation,omitempty"`
	BusinessActivity   *string `json:"business_activity,omitempty"`
	CACCertificateURL  *string `json:"cac_certificate_url,omitempty"`
	MermatURL          *string `json:"mermat_url,omitempty"`
	StatusReportURL    *string `json:"status_report_url,omitempty"`
}

// GetCompanyProfileResponse DTO
type GetCompanyProfileResponse struct {
	ID                 uuid.UUID `json:"id,omitempty"`
	Company            uuid.UUID `json:"company,omitempty"`
	RCNumber           string    `json:"rc_number,omitempty"`
	BusinessTin        string    `json:"business_tin,omitempty"`
	BusinessType       string    `json:"business_type,omitempty"`
	IncorporationYear  string    `json:"incorporation_year,omitempty"`
	IncorporationState string    `json:"incorporation_state,omitempty"`
	YearsInOperation   int       `json:"years_in_operation,omitempty"`
	BusinessActivity   string    `json:"business_activity,omitempty"`
	CACCertificateURL  string    `json:"cac_certificate_url,omitempty"`
	MermatURL          string    `json:"mermat_url,omitempty"`
	StatusReportURL    string    `json:"status_report_url,omitempty"`
}

// CompanyProfileDataResponse returns company file response
type CompanyProfileDataResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Data    GetCompanyProfileResponse `json:"data"`
}

// FilterCompanyProfileDataResponse returns company files responses
type FilterCompanyProfileDataResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Data    []GetCompanyProfileResponse `json:"data"`
}

// CreateCompanyProfileRequest DTO to create company file
type CreateCompanyProfileRequest struct {
	Company            uuid.UUID `json:"company" binding:"required"`
	RCNumber           string    `json:"rc_number" binding:"required"`
	BusinessTin        string    `json:"business_tin" binding:"required"`
	BusinessType       string    `json:"business_type" binding:"required"`
	IncorporationYear  string    `json:"incorporation_year" binding:"required"`
	IncorporationState string    `json:"incorporation_state" binding:"required"`
	YearsInOperation   int       `json:"years_in_operation" binding:"required"`
	BusinessActivity   string    `json:"business_activity" binding:"required"`
	CACCertificateURL  string    `json:"cac_certificate_url" binding:"required"`
	MermatURL          string    `json:"mermat_url" binding:"required"`
	StatusReportURL    string    `json:"status_report_url" binding:"required"`
}
