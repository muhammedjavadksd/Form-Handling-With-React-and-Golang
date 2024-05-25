package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	GrpcPort string
}

func NewConfig() *Config {

	formGrpc := os.Getenv("GRPC_PORT")
	return &Config{

		GrpcPort: formGrpc,
	}
}
