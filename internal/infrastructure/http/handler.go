package http

import (
	"fmt"
	"net/http"
	"user_service/internal/application/dto"
	"user_service/internal/infrastructure/config"
	"user_service/internal/ports"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type HTTPServer struct {
	engine      *gin.Engine
	validate    *validator.Validate
	userService ports.UserService
}

func NewHTTPServer(cfg *config.Config, userService ports.UserService, validate *validator.Validate) *HTTPServer {
	engine := gin.Default()
	server := &HTTPServer{
		engine:      engine,
		validate:    validate,
		userService: userService,
	}
	server.registerRoutes()
	return server
}

func (s *HTTPServer) Run(port string) {
	if err := s.engine.Run(port); err != nil {
		panic(err)
	}
}

func (s *HTTPServer) registerRoutes() {
	s.engine.POST("/users", s.createUser)
}

func (s *HTTPServer) createUser(c *gin.Context) {
	var user dto.CreateUser

	// Intentar vincular el cuerpo JSON al DTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar los datos con el validador
	if err := s.validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error de validación: %s", err.Error())})
		return
	}

	// Crear el usuario usando el servicio
	createdUser, err := s.userService.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Responder con el usuario creado
	c.JSON(http.StatusCreated, createdUser)
}
