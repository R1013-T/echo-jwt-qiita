package utils

import (
	"github.com/joho/godotenv"
	"os"
)

func Getenv(key string) string {
	if err := godotenv.Load(); err != nil {
		return ""
	}

	return os.Getenv(key)
}
