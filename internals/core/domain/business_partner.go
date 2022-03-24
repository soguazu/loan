package domain

import (
	"github.com/satori/go.uuid"
)

// BusinessPartner model
type BusinessPartner struct {
	Base
	Company uuid.UUID `json:"company,omitempty" gorm:"foreignKey:Company;references:ID"`
	Phone   string    `json:"phone" gorm:"not null"`
	Name    string    `json:"name" gorm:"not null"`
}
