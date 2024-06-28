package styleprompt

import (
	"context"
	// "gorm.io/gorm"
)

// StylePrompt represents the structure of a style prompt in the system.
type StylePrompt struct {
	ID       uint   `gorm:"autoIncrement;column:id"`
	Name     string `gorm:"column:name"`
	Language string `gorm:"column:language"`
}

// StylePromptInterface defines the set of methods that any implementation of the StylePrompt service must provide.
type StylePromptInterface interface {
	CreateStylePrompt(ctx context.Context, stylePrompt StylePrompt) (uint, error) // Creates a new style prompt and returns the prompt ID
	GetStylePromptByID(ctx context.Context, id uint) (*StylePrompt, error)        // Fetches a style prompt by its ID
	// UpdateStylePrompt(ctx context.Context, stylePrompt StylePrompt) error          // Updates an existing style prompt
	// DeleteStylePrompt(ctx context.Context, id uint) error                         // Deletes a style prompt by its ID
	ListStylePrompts(ctx context.Context, language string) (*[]StylePrompt, error) // Lists all style prompts by language
}

// NewStylePromptRequest defines the structure of a request to create a new style prompt.
type NewStylePromptRequest struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}
