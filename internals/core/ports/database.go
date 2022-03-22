package ports

import "gorm.io/gorm"

// IDatabase defines the interface for database management
type IDatabase interface {
	ConnectDB(url string) *gorm.DB
	MigrateAll(db *gorm.DB) error
}
