// adaptor.go

package framework

import (
	"context"
	"errors"
	// "regexp"
	"gorm.io/gorm"
)


type Adaptor struct {
	db *gorm.DB
}

func NewAdaptor(db *gorm.DB) *Adaptor {
	return &Adaptor{db: db}
}

func (a *Adaptor) ValidateNewFramework(ctx context.Context, framework Framework) error {
	// Validate the framework name is not empty.
	if framework.Name == "" {
		return errors.New("Framework name cannot be empty")
	}

	return nil
}