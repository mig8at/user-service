package ports

import (
	"context"
	"user_service/internal/application/dto"
)

type UserService interface {
	CreateUser(ctx context.Context, user *dto.CreateUser) (*dto.User, error)
}
