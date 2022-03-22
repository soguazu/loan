package ports

import "gorm.io/gorm"

//IUnitOfWork creates an instance of gorm transaction
type IUnitOfWork interface {
	Begin() (*gorm.DB, error)
	Commit() error
	Rollback() error
}
