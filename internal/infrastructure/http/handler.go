package http

import (
	"net/http"
	"user_service/internal/domain"
	"user_service/internal/infrastructure/config"
	"user_service/internal/ports"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HTTPServer struct {
	engine      *gin.Engine
	userService ports.UserService
}

func NewHTTPServer(cfg *config.Config, userService ports.UserService) *HTTPServer {
	engine := gin.Default()
	server := &HTTPServer{
		engine:      engine,
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
	s.engine.GET("/users/:id", s.getUser)
	s.engine.GET("/users", s.getUsers)
}

func (s *HTTPServer) createUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generar UUID para el usuario
	user.ID = uuid.NewString()

	if err := s.userService.RegisterUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (s *HTTPServer) getUser(c *gin.Context) {
	id := c.Param("id")
	user, err := s.userService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *HTTPServer) getUsers(c *gin.Context) {
	users, err := s.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuarios no encontrados"})
		return
	}
	c.JSON(http.StatusOK, users)
}
