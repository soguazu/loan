package main

import (
	"core_business/internals/core/domain"
	"core_business/pkg/config"
	"core_business/pkg/database"
	"fmt"
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

	fmt.Println("seeding data...")
	seed(conn)

}

// seed the database with some data
func seed(db *gorm.DB) error {

	categories := []domain.ExpenseCategory{
		{
			Title: "Marketing",
		},
		{
			Title: "Meetings",
		},
		{
			Title: "Shipping/Logistics",
		},
		{
			Title: "Food",
		},
		{
			Title: "Fuel",
		},
		{
			Title: "Car Rental/Rideshare",
		},
		{
			Title: "Subscription",
		},
		{
			Title: "Office Supplies",
		},
		{
			Title: "Entertainment",
		},
		{
			Title: "Maintenance",
		},
		{
			Title: "Transportation",
		},
		{
			Title: "Telecom/Airtime",
		},
		{
			Title: "Flights",
		},
		{
			Title: "Lodging",
		},
		{
			Title: "Training",
		},
		{
			Title: "Others",
		},
	}

	pricing := []domain.Fee{
		{
			Channel:    "Transaction both WEB and POS - card",
			Identifier: "ngn-card-both",
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
			Identifier: "ngn-card-create",
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

	if len(fees) < 1 {
		if err := tx.Create(&pricing).Error; err != nil {
			return err
		}
	}

	var expenseCategory []domain.ExpenseCategory

	if err := tx.Find(&expenseCategory).Error; err != nil {
		return err
	}

	if len(expenseCategory) < 1 {
		if err := tx.Create(&categories).Error; err != nil {
			return err
		}
	}

	return tx.Commit().Error

}
