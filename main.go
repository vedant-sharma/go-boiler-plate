package main

import (
	"log"

	"go-api-template/cmd"
	"go-api-template/config"
	"go-api-template/db"

	"github.com/joho/godotenv"
)

func main() {

	if godotenv.Load(".env") != nil {
		log.Fatal("Error loading .env file")
	}

	confVars, configErr := config.New()

	if configErr != nil {
		log.Fatal(configErr)
	}

	dbErr := db.Init()

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	defer db.Close()

	app := cmd.InitApp()

	app.Listen(confVars.Port)
}
