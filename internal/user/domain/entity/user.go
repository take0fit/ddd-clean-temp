package entity

import (
	"github.com/take0fit/ddd-clean-temp/internal/user/domain/valueobject"
	"time"
)

type User struct {
	Id       int
	Name     string
	Birthday valueobject.Birthday `gorm:"type:date"`
}

func NewUser() *User {
	return &User{}
}

func (b User) GetBirthday() *time.Time {
	if b.Birthday.Valid {
		return &b.Birthday.Time
	}
	return nil
}
