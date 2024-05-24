package di

import (
	"api-gateway/pkg/api"
	handler "api-gateway/pkg/api/handlers"
	"api-gateway/pkg/config"
	formclient "api-gateway/pkg/form-client"
)

func InitDI() *api.Server {
	cfg := config.NewConfig()
	conn := formclient.InitClient(cfg)
	form := handler.NewFormHandler(conn)
	server := api.NewHTTPServer(cfg, *form)
	return server

}
