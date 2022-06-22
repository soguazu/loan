package domain

// PAN model
type PAN struct {
	Base
	Number string `json:"number" gorm:"index;not null"`
	Status bool   `json:"status" gorm:"not null;default:true"`
}
