package review

import (
	"context"
	"errors"

	"gorm.io/gorm"
)


type Adaptor struct {
	db *gorm.DB
}

func NewAdaptor(db *gorm.DB) *Adaptor {
	return &Adaptor{db: db}
}

func (a *Adaptor) ValidateNewReview(ctx context.Context, review ReviewEntity) error {
	// Validate the AgentDetail name is not empty.
	if review.ID == 0 {
		return errors.New("AgentDetail ID cannot be zero")
	}

	// Add additional validations as needed.

	// If all validations pass, no error is returned.
	return nil
}