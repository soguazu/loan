package domain

import (
	"github.com/satori/go.uuid"
	"time"
)

// Company model
type Company struct {
	ID              uuid.UUID         `gorm:"type:uuid;primaryKey;autoIncrement:false"`
	Owner           uuid.UUID         `json:"owner" gorm:"not null;index"`
	Name            string            `json:"name" gorm:"UNIQUE_INDEX:business;index;not null"`
	Website         string            `json:"website" gorm:"index"`
	Type            string            `json:"type" gorm:"index"`
	FundingSource   string            `json:"funding_source"`
	NoOfEmployee    int32             `json:"no_of_employee" gorm:"not null;default:0"`
	Address         []Address         `json:"address,omitempty" gorm:"foreignKey:Company;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BusinessHead    BusinessHead      `json:"business_head,omitempty" gorm:"foreignKey:Company;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BusinessPartner []BusinessPartner `json:"business_partner,omitempty" gorm:"foreignKey:Company;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CompanyProfile  CompanyProfile    `json:"company_profile,omitempty" gorm:"foreignKey:Company;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Wallet          Wallet            `json:"wallet,omitempty" gorm:"foreignKey:Company;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}
