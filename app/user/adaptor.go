// adaptor.go

package user

import (
	"context"
	"errors"
	"regexp"

	"gorm.io/gorm"
)

// Adaptor is responsible for implementing domain-specific logic for users.
type Adaptor struct {
	db *gorm.DB
}

// NewAdaptor creates a new instance of Adaptor with a database connection.
func NewAdaptor(db *gorm.DB) *Adaptor {
	return &Adaptor{db: db}
}

// NewUser contains domain logic for creating a new user.
// This example function might validate user data before creation.
func (a *Adaptor) ValidateNewUser(ctx context.Context, user User) error {
	// Validate the email address format.
	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	// Validate the username is not empty.
	if user.FirebaseID == "" {
		return errors.New("username cannot be empty")
	}

	// Add additional validations as needed.

	// If all validations pass, no error is returned.
	return nil
}

// Helper function to validate email format.
func isValidEmail(email string) bool {
	// Simple regex for demonstration purposes; consider a more robust solution for production.
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

// Additional domain-specific methods can be added here...
