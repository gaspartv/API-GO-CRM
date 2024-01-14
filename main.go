package main

import (
	"os"

	"github.com/gaspartv/API-GO-CRM/configs"
	"github.com/gaspartv/API-GO-CRM/routers"
	"github.com/joho/godotenv"
)

var (
	logger configs.Logger
)

func main() {
	if err := configs.Load(); err != nil {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file")
		os.Exit(1)
	}

	logger = *configs.GetLogger("main")

	routers.Initialize()
}
