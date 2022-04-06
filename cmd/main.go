package main

import (
	"core_business/cmd/server"
	_ "core_business/docs"
	"core_business/pkg/database"
	"log"
)

// @title Evea Core Business Swagger API
// @version 1.0
// @description Evea Core Business Swagger API.
// @termsOfService http://swagger.io/terms/

// @contact.name Evea Team API Support
// @contact.email info@evea.com

// @license.name MIT
// @license.url https://github.com/sguazu

// @BasePath /v1
func main() {
	var DBConnection = database.NewDatabase()
	err := server.Run(DBConnection)
	if err != nil {
		log.Fatal(err)
		return
	}
	server.Injection()
}
