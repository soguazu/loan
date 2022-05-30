package domain

import (
	"github.com/satori/go.uuid"
	"time"
)

// Company model
type Company struct {
	ID              uuid.UUID         `gorm:"primary_key; unique;autoIncrement:false;type:uuid; column:id"`
	Owner           string            `json:"owner" gorm:"not null;index"`
	Name            string            `json:"name" gorm:"UNIQUE_INDEX:business;index;not null"`
	Website         string            `json:"website" gorm:"index"`
	Type            string            `json:"type" gorm:"index"`
	FundingSource   string            `json:"funding_source"`
	NoOfEmployee    string            `json:"no_of_employee" gorm:"not null;default:0"`
	Address         []Address         `json:"address,omitempty" gorm:"ForeignKey:Company;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BusinessHead    BusinessHead      `json:"business_head,omitempty" gorm:"ForeignKey:Company;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BusinessPartner []BusinessPartner `json:"business_partner,omitempty" gorm:"ForeignKey:Company;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CompanyProfile  CompanyProfile    `json:"company_profile,omitempty" gorm:"ForeignKey:Company;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Wallet          Wallet            `json:"wallet,omitempty" gorm:"ForeignKey:Company;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Transaction     []Transaction     `json:"transaction" gorm:"ForeignKey:Company;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Card            []Card            `json:"card,omitempty" gorm:"ForeignKey:Company;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}
