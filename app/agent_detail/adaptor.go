// adaptor.go

package agentdetail

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

func (a *Adaptor) ValidateNewAgentDetail(ctx context.Context, agentDetail AgentDetail) error {
	// Validate the AgentDetail name is not empty.
	if agentDetail.Name == "" {
		return errors.New("AgentDetail name cannot be empty")
	}

	// Add additional validations as needed.

	// If all validations pass, no error is returned.
	return nil
}