package repository

import (
	"context"
	"fmt"
	"log"
	"time"
	"user_service/internal/application/dto"
	"user_service/internal/domain/models"
	"user_service/internal/infrastructure/config"
	"user_service/internal/ports"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(cfg *config.Config) ports.UserRepository {
	db, err := gorm.Open(sqlite.Open(cfg.SqlitePath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos SQLite: %v", err)
	}

	// Migrar los modelos para crear tablas automáticamente
	if err := db.AutoMigrate(&models.User{}, &models.Follower{}); err != nil {
		log.Fatalf("Error al migrar las tablas: %v", err)
	}
	db.Exec("PRAGMA foreign_keys = ON;")

	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, createUser *dto.CreateUser) (*models.User, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	newUser := &models.User{
		ID:       uuid.NewString(),
		Name:     createUser.Name,
		Email:    createUser.Email,
		Nickname: createUser.Nickname,
		Bio:      createUser.Bio,
		Avatar:   createUser.Avatar,
	}

	if err := r.db.WithContext(ctx).Create(newUser).Error; err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("operación cancelada por exceder el límite de tiempo")
		}
		return nil, err
	}

	return newUser, nil
}
