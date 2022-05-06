package domain

// ExpenseCategory model
type ExpenseCategory struct {
	Base
	Title       string      `json:"title" gorm:"index;not null"`
	Transaction Transaction `json:"transaction,omitempty" gorm:"ForeignKey:ExpenseCategory;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
