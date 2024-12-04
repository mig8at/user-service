package main

import (
	"user_service/internal/application"
	"user_service/internal/infrastructure/config"
	"user_service/internal/infrastructure/http"
	"user_service/internal/infrastructure/repository"

	"github.com/go-playground/validator/v10"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()
	validate := validator.New()

	// Ejecutar seeder solo si el entorno es desarrollo
	// if cfg.Env == "development" {
	// 	seed := seeder.NewSeeder(cfg)
	// 	seed.Seed()
	// }

	// Inicializar repositorio
	repo := repository.NewRepository(cfg)

	// Inicializar servicio de aplicación
	service := application.NewService(repo)

	// Inicializar y ejecutar servidor HTTP
	httpServer := http.NewHTTPServer(cfg, service, validate)
	httpServer.Run(cfg.Port)
}
