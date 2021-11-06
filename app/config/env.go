package config

import (
	"github.com/joho/godotenv"
)

func LoadEnvironment() error {
	return godotenv.Load(".env")
}
