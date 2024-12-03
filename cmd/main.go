package main

import (
	"user_service/internal/application"
	"user_service/internal/infrastructure/config"
	"user_service/internal/infrastructure/http"
	"user_service/internal/infrastructure/repository"
	"user_service/internal/infrastructure/seeder"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Ejecutar seeder solo si el entorno es desarrollo
	if cfg.Env == "development" {
		seed := seeder.NewSeeder(cfg)
		seed.Seed()
	}

	// Inicializar repositorio
	userRepo := repository.NewBadgerRepository(cfg)

	// Inicializar servicio de aplicación
	userService := application.NewUserService(userRepo)

	// Inicializar y ejecutar servidor HTTP
	httpServer := http.NewHTTPServer(cfg, userService)
	httpServer.Run(cfg.Port)
}
