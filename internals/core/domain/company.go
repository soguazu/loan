package domain

import (
	"github.com/satori/go.uuid"
)

// Company model
type Company struct {
	Base
	Owner         uuid.UUID `json:"owner," gorm:"not null;index"`
	Name          string    `json:"name" gorm:"UNIQUE_INDEX:business;index;not null"`
	Website       string    `json:"website" gorm:"index"`
	Type          string    `json:"type" gorm:"index"`
	FundingSource string    `json:"funding_source"`
	NoOfEmployee  int32     `json:"no_of_employee" gorm:"not null;default:0"`
}
