package main

import (
	_ "embed"

	"github.com/joho/godotenv"
	"github.com/juniperwilson/archeyeve/external/api"
	"github.com/juniperwilson/archeyeve/internal/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	database.Start()
	defer database.Stop()

	SeedAll()

	router := api.ApiRouting()
	router.Run(":8080")
}
