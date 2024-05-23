package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port string
}

func NewConfig() *Config {
	port := os.Getenv("PORT")
	return &Config{
		Port: port,
	}
}
