package coins

import (
	"context"
	"gorm.io/gorm"
)

// Adaptor is responsible for the data storage operations and validation logic related to coins.
type Adaptor struct {
	db *gorm.DB
}

// NewAdaptor creates a new instance of Adaptor with a database connection.
func NewAdaptor(db *gorm.DB) *Adaptor {
	return &Adaptor{db: db}
}

// ValidateNewCoins performs validation on a new coins record.
func (a *Adaptor) ValidateNewCoins(ctx context.Context, coins CoinsEntity) error {
	// Add your validation logic here
	return nil
}