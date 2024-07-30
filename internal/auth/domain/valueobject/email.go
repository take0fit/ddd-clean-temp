package valueobject

import (
	"github.com/take0fit/validationcontext"
)

type Email string

func NewEmail(email string, vc *validationcontext.ValidationContext, required bool) Email {
	if required {
		vc.Required(email, "Email", "Email is required", true)
	}
	vc.ValidateEmail(email, "Email", "Invalid email format")
	return Email(email)
}
