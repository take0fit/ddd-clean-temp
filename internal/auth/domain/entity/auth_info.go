package entity

import (
	"github.com/take0fit/ddd-clean-temp/internal/auth/domain/valueobject"
	"golang.org/x/crypto/bcrypt"
)

type AuthInfo struct {
	UserID       int
	Email        valueobject.Email
	PasswordHash string
}

func NewAuthInfo(userID int, email valueobject.Email, password valueobject.Password) (*AuthInfo, error) {
	passwordHash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	return &AuthInfo{
		UserID:       userID,
		Email:        email,
		PasswordHash: passwordHash,
	}, nil
}

// NewAuthInfoFromDB データベースから読み込むためのファクトリメソッド
func NewAuthInfoFromDB(userID int, email valueobject.Email, passwordHash string) *AuthInfo {
	return &AuthInfo{
		UserID:       userID,
		Email:        email,
		PasswordHash: passwordHash,
	}
}

func hashPassword(password valueobject.Password) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (a *AuthInfo) VerifyPassword(password valueobject.Password) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.PasswordHash), []byte(password))
	return err == nil
}
