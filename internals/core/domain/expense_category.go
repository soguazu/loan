package domain

// ExpenseCategory model
type ExpenseCategory struct {
	Base
	Title string `json:"title" gorm:"index;not null"`
}
