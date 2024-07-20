package role

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

// Adaptor is responsible for implementing domain-specific logic for roles.
type Adaptor struct {
	db *gorm.DB
}

// NewAdaptor creates a new instance of Adaptor with a database connection.
func NewAdaptor(db *gorm.DB) *Adaptor {
	return &Adaptor{db: db}
}

// ValidateNewRole contains domain logic for creating a new role.
func (a *Adaptor) ValidateNewRole(ctx context.Context, role Role) error {
	// Validate the role name is not empty.
	if role.Name == "" {
		return errors.New("role name cannot be empty")
	}

	// Add additional validations as needed.

	// If all validations pass, no error is returned.
	return nil
}

// Additional domain-specific methods can be added here...
