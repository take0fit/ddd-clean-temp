package valueobject

import (
	"github.com/take0fit/validationcontext"
)

type UserID int
type UserName string

func NewUserID(id int, vc *validationcontext.ValidationContext) UserID {
	if id <= 0 {
		vc.AddError("Id", "Id must be greater than 0")
	}
	return UserID(id)
}

func NewUserName(name string, vc *validationcontext.ValidationContext) UserName {
	if name == "" {
		vc.AddError("Name", "Name is required")
	}
	return UserName(name)
}
