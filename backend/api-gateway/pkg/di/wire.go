package di

import (
	"api-gateway/pkg/api"
	handler "api-gateway/pkg/api/handlers"
	"api-gateway/pkg/config"
)

func InitDI() *api.Server {
	cfg := config.NewConfig()
	form := handler.NewFormHandler()
	server := api.NewHTTPServer(cfg, *form)
	return server

}
