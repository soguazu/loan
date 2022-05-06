package domain

// Fee model
type Fee struct {
	Base
	Channel      string      `gorm:"type:varchar(100);not null"`
	Identifier   string      `gorm:"type:varchar(100);not null;unique;"`
	Fee          float64     `gorm:"not null;" sql:"type:decimal(10,8);"`
	IsPercent    bool        `gorm:"not null"`
	IsDollar     bool        `gorm:"not null"`
	Transactions Transaction `json:"transaction" gorm:"ForeignKey:Fee;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
