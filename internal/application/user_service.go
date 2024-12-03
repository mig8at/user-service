package application

import (
	"user_service/internal/domain"
	"user_service/internal/ports"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	repo      ports.UserRepository
	validator *validator.Validate
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{
		repo:      repo,
		validator: validator.New(),
	}
}

func (s *userService) RegisterUser(user *domain.User) error {
	// Validar los datos del usuario
	if err := s.validator.Struct(user); err != nil {
		return err
	}

	// Guardar el usuario en el repositorio
	return s.repo.Create(user)
}

func (s *userService) GetUser(id string) (*domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) GetUsers() ([]domain.User, error) {
	return s.repo.GetAll()
}
