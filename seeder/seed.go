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
			Channel:    "Transaction POS - VAN card",
			Identifier: "ngn-van-card-pos",
			Fee:        3,
			IsDollar:   false,
			IsPercent:  true,
		},
		{
			Channel:    "Transaction WEB - VAN card",
			Identifier: "ngn-van-card-web",
			Fee:        3,
			IsDollar:   false,
			IsPercent:  true,
		},
		{
			Channel:    "Transaction ATM - VAN card",
			Identifier: "ngn-van-card-atm",
			Fee:        5,
			IsDollar:   false,
			IsPercent:  true,
		},
		{
			Channel:    "MGN Virtual Card Creation",
			Identifier: "ngn-vc-create",
			Fee:        3,
			IsDollar:   true,
			IsPercent:  false,
		},
		{
			Channel:    "Card Shipping",
			Identifier: "card-shipping",
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

	if err := tx.Create(&domain.Fee{}).Error; err != nil {
		return err
	}

	if err := tx.Create(&pricing).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}
