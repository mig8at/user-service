package application

import (
	"context"
	"user_service/internal/application/dto"
	"user_service/internal/ports"
)

type userService struct {
	repo ports.UserRepository
}

func NewService(repo ports.UserRepository) ports.UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *dto.CreateUser) (*dto.User, error) {

	newUser, err := s.repo.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return &dto.User{
		ID:       newUser.ID,
		Name:     newUser.Name,
		Email:    newUser.Email,
		Nickname: newUser.Nickname,
		Bio:      newUser.Bio,
		Avatar:   newUser.Avatar,
	}, nil
}
