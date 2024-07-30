package valueobject

import (
	"testing"

	"github.com/take0fit/validationcontext"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		required  bool
		expectErr bool
	}{
		{"ValidEmail", "test@example.com", true, false},
		{"InvalidEmailFormat", "invalid-email", true, true},
		{"EmptyEmail", "", true, true},
		{"OptionalEmptyEmail", "", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := validationcontext.NewValidationContext()
			NewEmail(tt.email, vc, tt.required)

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
