package ports

import (
	"context"
	"user_service/internal/application/dto"
	"user_service/internal/domain/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *dto.CreateUser) (*models.User, error)
}
