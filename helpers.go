package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func loadEnvironmentVars() map[string]string {
	var envs map[string]string

	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return envs
}

func buildConnectionString(envars *map[string]string) string {
	return fmt.Sprintf("mongodb://%s:%s@localhost:27017", (*envars)["MONGO_USER"], (*envars)["MONGO_PASSWORD"])
}
