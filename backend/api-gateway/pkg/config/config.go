package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port     string
	FormGrpc string
}

func NewConfig() *Config {
	port := os.Getenv("PORT")
	formGrpc := os.Getenv("FORM_GRPC_PORT")
	return &Config{
		Port:     port,
		FormGrpc: formGrpc,
	}
}
