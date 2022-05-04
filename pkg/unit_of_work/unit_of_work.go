package tx

import (
	"core_business/internals/core/ports"
	"gorm.io/gorm"
)

type gormUnitOfWork struct {
	db *gorm.DB
}

// NewGormUnitOfWork will create a new gorm unit of work
func NewGormUnitOfWork(db *gorm.DB) ports.IUnitOfWork {
	return &gormUnitOfWork{db: db}
}

func (u *gormUnitOfWork) Begin() (*gorm.DB, error) {
	tx := u.db.Begin()
	u.db = tx
	return tx, tx.Error
}

func (u *gormUnitOfWork) Commit() error {
	return u.db.Commit().Error
}

func (u *gormUnitOfWork) Rollback() error {
	return u.db.Rollback().Error
}
