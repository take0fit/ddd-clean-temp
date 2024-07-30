package valueobject

import (
	"github.com/take0fit/validationcontext"
)

type Password string

func NewPassword(password string, vc *validationcontext.ValidationContext, required bool) Password {
	if required {
		vc.Required(password, "Password", "Password is required", true)
	}
	vc.ValidateMinLength(password, "Password", 8, "Password must be at least 8 characters long")
	vc.ValidateContainsSpecial(password, "Password", "Password must contain at least one special character (!@#)")
	return Password(password)
}
