package tx

import (
	log "github.com/sirupsen/logrus"
	"core_business/internals/core/ports"
	"gorm.io/gorm"
)

type gormUnitOfWork struct {
	db     *gorm.DB
	logger *log.Logger
}

// NewGormUnitOfWork will create a new gorm unit of work
func NewGormUnitOfWork(db *gorm.DB, l *log.Logger) ports.IUnitOfWork {
	return &gormUnitOfWork{db: db, logger: l}
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
