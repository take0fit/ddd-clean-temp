package valueobject

import "github.com/take0fit/validationcontext"

type Email string

func NewEmail(email string, vc *validationcontext.ValidationContext) Email {
	if email == "" {
		vc.AddError("Email", "Email is required")
	}
	return Email(email)
}
