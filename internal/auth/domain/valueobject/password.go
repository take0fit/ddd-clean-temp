package valueobject

import "github.com/take0fit/validationcontext"

type Password string

func NewPassword(password string, vc *validationcontext.ValidationContext) Password {
	if password == "" {
		vc.AddError("Password", "Password is required")
	}
	return Password(password)
}
