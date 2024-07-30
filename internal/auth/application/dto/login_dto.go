package dto

import (
	"github.com/take0fit/ddd-clean-temp/internal/auth/domain/valueobject"
	"github.com/take0fit/validationcontext"
)

type InputLoginDTO struct {
	Email    valueobject.Email
	Password valueobject.Password
}

func NewInputLoginDTO(email, password string) (*InputLoginDTO, error) {
	vc := validationcontext.NewValidationContext()

	emailVO := valueobject.NewEmail(email, vc, true)
	passwordVO := valueobject.NewPassword(password, vc, true)

	if vc.HasErrors() {
		return nil, vc.AggregateError()
	}

	return &InputLoginDTO{
		Email:    emailVO,
		Password: passwordVO,
	}, nil
}
