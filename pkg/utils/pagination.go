package utils

import (
	"gorm.io/gorm"
	"math"
)

// Pagination manages pagination
type Pagination struct {
	Limit      int         `form:"limit,default=5" binding:"min=5"`
	Page       int         `form:"page,default=1" binding:"min=1"`
	Sort       string      `json:"sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

// GetOffset calculate the offset
func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

// GetLimit assigns default value to limit if it is 0
func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

// GetPage Assign default value to page if it is 0
func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

// GetSort assign the default to sort if it wasn't set
func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "created_at asc"
	}
	return p.Sort
}

// Paginate this handles the magic of pagination
func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
