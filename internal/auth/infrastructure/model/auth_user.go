package model

import "time"

type AuthUser struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"uniqueIndex"`
	Email        string `gorm:"uniqueIndex"`
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
