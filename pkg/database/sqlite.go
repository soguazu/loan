package database

import (
	"fmt"
	"github.com/soguazu/core_business/internals/core/domain"
	"github.com/soguazu/core_business/internals/core/ports"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteDatastore struct {
}

// NewSqliteDatabase creates a new instance for managing database
func NewSqliteDatabase() ports.IDatabase {
	return &sqliteDatastore{}
}

func (d *sqliteDatastore) ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(
			err.Error(),
		)
		panic("failed to connect database")
	}

	fmt.Println("Established database connection")
	return db
}

func (d *sqliteDatastore) MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Company{})
}
