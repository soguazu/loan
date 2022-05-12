package main

import (
	"core_business/internals/core/domain"
	"core_business/pkg/config"
	"core_business/pkg/database"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	database := database.NewDatabase()
	conn := database.ConnectDB(config.Instance.DatabaseURL)

	err = database.MigrateAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		sqlDB, _ := conn.DB()
		err := sqlDB.Close()
		if err != nil {
			return
		}
	}()

	seed(conn)

}

// seed the database with some data
func seed(db *gorm.DB) error {

	pricing := []domain.Fee{
		{
			Channel:    "Transaction POS - card",
			Identifier: "ngn-card-pos",
			Fee:        3,
			IsDollar:   false,
			IsPercent:  true,
		},
		{
			Channel:    "Transaction WEB - card",
			Identifier: "ngn-card-web",
			Fee:        3,
			IsDollar:   false,
			IsPercent:  true,
		},
		{
			Channel:    "Transaction ATM - card",
			Identifier: "ngn-card-atm",
			Fee:        5,
			IsDollar:   false,
			IsPercent:  true,
		},
		{
			Channel:    "NGN Virtual Card Creation - card",
			Identifier: "ngn-vc-create",
			Fee:        1700,
			IsDollar:   false,
			IsPercent:  false,
		},
		{
			Channel:    "Card Shipping - card",
			Identifier: "ngn-card-shipping",
			Fee:        1000,
			IsDollar:   false,
			IsPercent:  false,
		},
	}

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil || tx.Error != nil {
			tx.Rollback()
		}
	}()

	var fees []domain.Fee

	if err := tx.Find(&fees).Error; err != nil {
		return err
	}

	if len(fees) > 0 {
		return nil
	}

	if err := tx.Create(&pricing).Error; err != nil {
		return err
	}

	return tx.Commit().Error

}
