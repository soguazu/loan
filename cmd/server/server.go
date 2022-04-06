package server

import (
	"core_business/internals/core/ports"
	"core_business/pkg/config"
	"gorm.io/gorm"
	"log"
)

// DBConnection stores the instance of the Database
var DBConnection *gorm.DB

// Run function starts the database connection
func Run(database ports.IDatabase) error {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
		return err
	}

	DBConnection = database.ConnectDB(config.Instance.DatabaseURL)
	err = database.MigrateAll(DBConnection)
	if err != nil {
		log.Fatal(err)
		return err
	}

	//defer func() {
	//	sqlDB, _ := DBConnection.DB()
	//	err := sqlDB.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}()

	return nil
}
