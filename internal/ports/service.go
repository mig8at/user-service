package ports

import "user_service/internal/domain"

type UserService interface {
	RegisterUser(user *domain.User) error
	GetUser(id string) (*domain.User, error)
	GetUsers() ([]domain.User, error)
}
