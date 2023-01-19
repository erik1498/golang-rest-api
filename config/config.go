package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	error := godotenv.Load(".env")
	if error != nil {
		fmt.Println("Error load .env file")
	}
	return os.Getenv(key)
}
