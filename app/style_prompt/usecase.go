package styleprompt

import (
	"context"
	"log"
)

type storage interface {
	CreateStylePrompt(ctx context.Context, stylePrompt StylePromptEntity) (int, error)
	GetStylePromptByID(ctx context.Context, id int) (*StylePromptEntity, error)
	// UpdateStylePrompt(ctx context.Context, stylePrompt StylePromptEntity) error
	// DeleteStylePrompt(ctx context.Context, id int) error
	ListStylePrompts(ctx context.Context, language string) (*[]StylePromptEntity, error)
}

type domain interface {
	ValidateNewStylePrompt(ctx context.Context, stylePrompt StylePrompt) error
}

type Usecase struct {
	storage storage
	domain  domain
}

func NewUsecase(s storage, d domain) *Usecase {
	return &Usecase{
		storage: s,
		domain:  d,
	}
}

func (u *Usecase) CreateStylePrompt(ctx context.Context, stylePrompt StylePrompt) (int, error) {

	if err := u.domain.ValidateNewStylePrompt(ctx, stylePrompt); err != nil {
		log.Printf("Error validating new style prompt: %v", err)
		return 0, err
	}

	stylePromptEntity := StylePromptEntity{
		Name:     stylePrompt.Name,
		Language: stylePrompt.Language,
	}
	log.Printf("StylePromptEntity: %+v\n", stylePromptEntity)

	id, err := u.storage.CreateStylePrompt(ctx, stylePromptEntity)
	return id, err
}
func (u *Usecase) ListStylePrompts(ctx context.Context, language string) (*[]StylePromptEntity, error) {
	stylePrompts, err := u.storage.ListStylePrompts(ctx, language)
	if err != nil {
		log.Printf("Error listing style prompts: %v", err)
		return nil, err
	}
	return stylePrompts, nil
}


func (u *Usecase) GetStylePromptByID(ctx context.Context, id int) (*StylePromptEntity, error) {
	stylePrompt, err := u.storage.GetStylePromptByID(ctx, id)
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

// func (u *Usecase) DeleteStylePrompt(ctx context.Context, id int) error {
// 	err := u.storage.DeleteStylePrompt(ctx, id)
// 	if err != nil {
// 		log.Printf("Error deleting style prompt: %v", err)
// 		return err
// 	}
// 	return nil
// }

