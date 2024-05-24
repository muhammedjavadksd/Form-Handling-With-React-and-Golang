package di

import (
	"form-manager/pkg/api"
	"form-manager/pkg/config"
)

func InitDI() *api.Server {
	cfg := config.NewConfig()
	server := api.InitGRPCServer(cfg)
	return server
}
