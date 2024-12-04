package models

import "time"

type User struct {
	ID        string    `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Nickname  string    `gorm:"not null;unique"`
	Bio       string    `gorm:"not null"`
	Avatar    string    `gorm:"not null"`
	Followers int       `gorm:"default:0"`
	Following int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Follower struct {
	UserID     string `gorm:"primaryKey"`
	FollowerID string `gorm:"primaryKey"`
}
