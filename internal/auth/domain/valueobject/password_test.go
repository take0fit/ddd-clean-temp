package valueobject

import (
	"testing"

	"github.com/take0fit/validationcontext"
)

func TestNewPassword(t *testing.T) {
	tests := []struct {
		name      string
		password  string
		required  bool
		expectErr bool
	}{
		{"ValidPassword", "Password@123", true, false},
		{"ShortPassword", "Pass@1", true, true},
		{"NoSpecialCharacter", "Password123", true, true},
		{"EmptyPassword", "", true, true},
		{"OptionalEmptyPassword", "", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := validationcontext.NewValidationContext()
			NewPassword(tt.password, vc, tt.required)

			hasErrors := vc.HasErrors()
			if hasErrors != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, hasErrors)
			}

			if hasErrors && len(vc.Errors()) != 1 {
				t.Errorf("Expected 1 error, got: %d", len(vc.Errors()))
			}
		})
	}
}
