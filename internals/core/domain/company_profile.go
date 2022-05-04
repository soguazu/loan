package domain

import (
	"github.com/satori/go.uuid"
)

// CompanyProfile model
type CompanyProfile struct {
	Base
	Company            uuid.UUID `json:"company,omitempty" gorm:"column:company"`
	RCNumber           string    `json:"rc_number" gorm:"index;not null"`
	BusinessTin        string    `json:"business_tin" gorm:"not null"`
	BusinessType       string    `json:"business_type" gorm:"not null"`
	IncorporationYear  string    `json:"incorporation_year" gorm:"not null"`
	YearsInOperation   int       `json:"years_in_operation" gorm:"not null"`
	IncorporationState string    `json:"incorporation_state" gorm:"not null"`
	BusinessActivity   string    `json:"business_activity" gorm:"not null"`
	CACCertificateURL  string    `json:"cac_certificate_url" gorm:"not null"`
	MermatURL          string    `json:"mermat_url" gorm:"not null"`
	StatusReportURL    string    `json:"status_report_url" gorm:"not null"`
}
