package styleprompt

import (
	"context"
	"log"
)


// storage outlines the methods required by the use case to interact with the data layer.
type storage interface {
	CreateStylePrompt(ctx context.Context, stylePrompt StylePromptEntity) (uint, error)
	GetStylePromptByID(ctx context.Context, id uint,language string) (*StylePromptEntity, error)
	// UpdateStylePrompt(ctx context.Context, stylePrompt StylePromptEntity) error
	// DeleteStylePrompt(ctx context.Context, id uint) error
	ListStylePrompts(ctx context.Context, language string) (*[]StylePromptEntity, error)
}

// domain outlines the methods required by the use case for domain logic and validations.
type domain interface {
	ValidateNewStylePrompt(ctx context.Context, stylePrompt StylePrompt) error
	// Add other domain methods as necessary
}

// Usecase struct that combines storage and domain to execute style prompt-related business logic.
type Usecase struct {
	storage storage
	domain  domain
}

// NewUsecase creates a new Usecase instance with the provided storage and domain logic implementations.
func NewUsecase(s storage, d domain) *Usecase {
	return &Usecase{
		storage: s,
		domain:  d,
	}
}

// CreateStylePrompt orchestrates the process of validating and creating a new style prompt.
func (u *Usecase) CreateStylePrompt(ctx context.Context, stylePrompt StylePrompt) (uint, error) {

	// First, use the domain logic to validate the new style prompt.
	if err := u.domain.ValidateNewStylePrompt(ctx, stylePrompt); err != nil {
		log.Printf("Error validating new style prompt: %v", err)
		return 0, err
	}

	stylePromptEntity := StylePromptEntity{
		Name:     stylePrompt.Name,
		Language: stylePrompt.Language,
	}
	log.Printf("StylePromptEntity: %+v\n", stylePromptEntity)

	// If validation passes, proceed to create the style prompt in the storage layer.
	id, err := u.storage.CreateStylePrompt(ctx, stylePromptEntity)
	return id, err
}
// ListStylePrompts gets a list of style prompts by language
func (u *Usecase) ListStylePrompts(ctx context.Context, language string) (*[]StylePromptEntity, error) {
	stylePrompts, err := u.storage.ListStylePrompts(ctx, language)
	if err != nil {
		log.Printf("Error listing style prompts: %v", err)
		return nil, err
	}
	return stylePrompts, nil
}


func (u *Usecase) GetStylePromptByID(ctx context.Context, id uint,language string) (*StylePromptEntity, error) {
	stylePrompt, err := u.storage.GetStylePromptByID(ctx, id,language)
	if err != nil {
		log.Printf("Error getting style prompt by ID: %v", err)
		return nil, err
	}
	return stylePrompt, nil
}

// func (u *Usecase) UpdateStylePrompt(ctx context.Context, stylePrompt StylePromptEntity) error {
// 	err := u.storage.UpdateStylePrompt(ctx, stylePrompt)
// 	if err != nil {
// 		log.Printf("Error updating style prompt: %v", err)
// 		return err
// 	}
// 	return nil
// }

// func (u *Usecase) DeleteStylePrompt(ctx context.Context, id uint) error {
// 	err := u.storage.DeleteStylePrompt(ctx, id)
// 	if err != nil {
// 		log.Printf("Error deleting style prompt: %v", err)
// 		return err
// 	}
// 	return nil
// }

