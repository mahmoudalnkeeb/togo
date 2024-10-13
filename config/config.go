package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBFile     string
	TodosTable string
}

func LoadConfig(pwd string) *Config {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}

	envFilePath := fmt.Sprintf("%s/.env", rootPath)
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return &Config{
		DBFile: fmt.Sprintf("%s/%s", rootPath, os.Getenv("DB_FILE")),
		TodosTable: `
        CREATE TABLE IF NOT EXISTS togo(
            ID INTEGER PRIMARY KEY AUTOINCREMENT,
            title VARCHAR NOT NULL,
            complete INTEGER NOT NULL DEFAULT 0 CHECK(complete IN (0, 1)),
            complete_by DATETIME
        );`,
	}
}
