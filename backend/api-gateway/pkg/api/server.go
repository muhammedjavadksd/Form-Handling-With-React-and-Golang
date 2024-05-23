package api

import (
	handler "api-gateway/pkg/api/handlers"
	"api-gateway/pkg/config"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	engine *fiber.App
	port   string
}

func NewHTTPServer(c *config.Config, form handler.FormHandler) *Server {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: false,
		MaxAge:           300,
	}))

	app.Post("/form", form.SaveForm)
	return &Server{
		engine: app,
		port:   c.Port,
	}

}

func (s *Server) Start() {
	go func() {
		if err := s.engine.Listen(":" + s.port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	log.Printf("Server started on port %s", s.port)
}

func (s *Server) Shutdown() {
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.engine.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("Server shut down gracefully")
}
