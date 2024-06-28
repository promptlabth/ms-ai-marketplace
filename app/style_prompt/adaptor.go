package styleprompt

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

func (a *Adaptor) ValidateNewStylePrompt(ctx context.Context, stylePrompt StylePrompt) error {
	if stylePrompt.Name == "" {
		return errors.New("StylePrompt name cannot be empty")
	}

	return nil
}

