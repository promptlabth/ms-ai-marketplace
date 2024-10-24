package styleprompt

import (
	"context"
	// "gorm.io/gorm"
)
type StylePrompt struct {
	ID       int   `gorm:"autoIncrement;column:id"`
	Name     string `gorm:"column:name"`
	Language string `gorm:"column:language"`
}

type StylePromptInterface interface {
	CreateStylePrompt(ctx context.Context, stylePrompt StylePrompt) (int, error) // Creates a new style prompt and returns the prompt ID
	GetStylePromptByID(ctx context.Context, id int) (*StylePrompt, error)        // Fetches a style prompt by its ID
	// UpdateStylePrompt(ctx context.Context, stylePrompt StylePrompt) error          // Updates an existing style prompt
	// DeleteStylePrompt(ctx context.Context, id int) error                         // Deletes a style prompt by its ID
	ListStylePrompts(ctx context.Context, language string) (*[]StylePrompt, error) // Lists all style prompts by language
}

type NewStylePromptRequest struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}
