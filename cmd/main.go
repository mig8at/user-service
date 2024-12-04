package main

import (
	"user_service/internal/application"
	"user_service/internal/infrastructure/config"
	"user_service/internal/infrastructure/http"
	"user_service/internal/infrastructure/repository"
	"user_service/internal/infrastructure/seeder"

	"github.com/go-playground/validator/v10"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()
	validate := validator.New()

	// Inicializar repositorio
	repo := repository.NewRepository(cfg)

	// Inicializar seeder utilizando la conexión existente del repositorio
	seed := seeder.NewSeeder(repo.GetDB())

	// Ejecutar el seeder solo en entornos de desarrollo o prueba
	if cfg.Env == "development" || cfg.Env == "test" {
		seed.Seed()
	}

	// Inicializar servicio de aplicación
	service := application.NewService(repo)

	// Inicializar y ejecutar servidor HTTP
	httpServer := http.NewHTTPServer(cfg, service, validate)
	httpServer.Run(cfg.Port)
}
