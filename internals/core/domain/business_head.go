package domain

import (
	"github.com/satori/go.uuid"
)

// BusinessHead model
type BusinessHead struct {
	Base
	Company                uuid.UUID `json:"company,omitempty" gorm:"foreignKey:Company;references:ID"`
	JobTitle               string    `json:"job_title" gorm:"index;not null"`
	Phone                  string    `json:"phone" gorm:"not null"`
	IdentificationType     string    `json:"identification_type" gorm:"not null"`
	IdentificationNumber   string    `json:"identification_number" gorm:"not null"`
	IdentificationImageURL string    `json:"identification_image_url" gorm:"not null"`
	CompanyIDUrl           string    `json:"company_id_url" gorm:"not null"`
}
