package ports

import "user_service/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id string) (*domain.User, error)
	GetAll() ([]domain.User, error)
}
