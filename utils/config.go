package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file ", err)
	}

	return os.Getenv(key)
}
