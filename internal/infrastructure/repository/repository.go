package repository

import (
	"context"
	"fmt"
	"log"
	"user_service/internal/application/dto"
	"user_service/internal/domain/models"
	"user_service/internal/infrastructure/config"
	"user_service/internal/ports"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
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

func (r *repository) GetDB() *gorm.DB {
	return r.db
}

func (r *repository) Create(ctx context.Context, createUser *dto.CreateUser) (*models.User, error) {

	userModel := &models.User{}
	if err := copier.Copy(userModel, createUser); err != nil {
		return nil, err
	}
	userModel.ID = uuid.NewString()

	if err := r.db.WithContext(ctx).Create(userModel).Error; err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("operación cancelada por exceder el límite de tiempo")
		}
		return nil, err
	}

	return userModel, nil
}

func (r *repository) GetById(ctx context.Context, id string) (*models.User, error) {

	user := &models.User{}

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(user).Error; err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("operación cancelada por exceder el límite de tiempo")
		}
		return nil, err
	}

	return user, nil
}

func (r *repository) Paginate(ctx context.Context, page, limit int) ([]models.User, error) {

	// Calcular el offset para la paginación
	offset := (page - 1) * limit

	// Arreglo para almacenar los usuarios
	var users []models.User

	// Realizar la consulta con paginación
	if err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("operación cancelada por exceder el límite de tiempo")
		}
		return nil, err
	}

	return users, nil
}

func (r *repository) Update(ctx context.Context, id string, updateUser *dto.UpdateUser) (*models.User, error) {

	user := &models.User{}

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(user).Error; err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("operación cancelada por exceder el límite de tiempo")
		}
		return nil, err
	}

	if updateUser.Name != "" {
		user.Name = updateUser.Name
	}
	if updateUser.Nickname != "" {
		user.Nickname = updateUser.Nickname
	}
	if updateUser.Bio != "" {
		user.Bio = updateUser.Bio
	}
	if updateUser.Avatar != "" {
		user.Avatar = updateUser.Avatar
	}

	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("operación cancelada por exceder el límite de tiempo")
		}
		return nil, err
	}

	return user, nil
}

func (r *repository) Follow(ctx context.Context, id, followerID string) error {

	follower := &models.Follower{
		ID:         uuid.NewString(),
		UserID:     id,
		FollowerID: followerID,
	}

	if err := r.db.WithContext(ctx).Create(follower).Error; err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("operación cancelada por exceder el límite de tiempo")
		}
		return err
	}

	return nil
}

func (r *repository) Unfollow(ctx context.Context, id, followerID string) error {

	if err := r.db.WithContext(ctx).Where("user_id = ? AND follower_id = ?", id, followerID).Delete(&models.Follower{}).Error; err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("operación cancelada por exceder el límite de tiempo")
		}
		return err
	}

	return nil
}

func (r *repository) Followers(ctx context.Context, id string, page, limit int) ([]models.User, error) {

	offset := (page - 1) * limit

	var users []models.User

	if err := r.db.WithContext(ctx).
		Table("followers").
		Select("users.*").
		Joins("JOIN users ON followers.follower_id = users.id").
		Where("followers.user_id = ?", id).
		Offset(offset).
		Limit(limit).
		Find(&users).Error; err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("operación cancelada por exceder el límite de tiempo")
		}
		return nil, err
	}

	return users, nil
}

func (r *repository) Following(ctx context.Context, id string, page, limit int) ([]models.User, error) {

	offset := (page - 1) * limit

	var users []models.User

	if err := r.db.WithContext(ctx).
		Table("followers").
		Select("users.*").
		Joins("JOIN users ON followers.user_id = users.id").
		Where("followers.follower_id = ?", id).
		Offset(offset).
		Limit(limit).
		Find(&users).Error; err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("operación cancelada por exceder el límite de tiempo")
		}
		return nil, err
	}

	return users, nil
}
