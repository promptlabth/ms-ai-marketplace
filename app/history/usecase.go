package history

import (
	"context"
	"log"
	"time"
)


type storage interface {
	CreateHistory(ctx context.Context, history HistoryEntity) (*int, error)
	// GetHistoryByID(ctx context.Context, id int) (*History, error)
	// ListHistories(ctx context.Context, userID int) (*[]History, error)
	// UpdateHistory(ctx context.Context, history History) error
	// DeleteHistory(ctx context.Context, id int) error
}

type domain interface {
	ValidateNewHistory(ctx context.Context, history History) error
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

// CreateHistory orchestrates the process of validating and creating a new history record
func (u *Usecase) CreateHistory(ctx context.Context, history History) error {
	// Optional: Validate the new history using domain logic
	if err := u.domain.ValidateNewHistory(ctx, history); err != nil {
		log.Printf("Validation error: %v", err)
		return err
	}

	historyEntity := HistoryEntity{
		UserID:         history.UserID,
		AgentID:        history.AgentID,
		FrameworkID:    history.ID,
		Prompt:         history.Prompt,
		StyleMessageID: history.StyleMessageID,
		LanguageID:     history.LanguageID,
		Result:         history.Result,
		TimeStamp:      time.Now(),
	}

	_, err := u.storage.CreateHistory(ctx, historyEntity)
	return err
}
// // CreateHistory orchestrates the process of validating and creating a new history record
// func (u *Usecase) CreateHistory(ctx context.Context, history History) (*string, error) {
// 	// Optional: Validate the new history using domain logic
// 	if err := u.domain.ValidateNewHistory(ctx, history); err != nil {
// 		log.Printf("Validation error: %v", err)
// 		return nil,err
// 	}
// 	result, err := u.generateMessage(history)
// 	if err != nil {
// 		log.Printf("Error generating message: %v", err)
// 		return nil, err
// 	}
// 	historyEntity := HistoryEntity{
// 		UserID:         history.UserID,
// 		AgentID:        history.AgentID,
// 		FrameworkID:    history.ID,
// 		Prompt:         history.Prompt,
// 		StyleMessageID: history.StyleMessageID,
// 		LanguageID:     history.LanguageID,
// 		Result:         result,
// 		TimeStamp:      time.Now(),
// 	}

// 	_, err := u.storage.CreateHistory(ctx, historyEntity)
// 	if err != nil {
// 		log.Printf("Error getting history by ID: %v", err)
// 		return nil, err
// 	}

// 	return result,err
// }


// func (u *Usecase) GetHistoryByID(ctx context.Context, id int) (*History, error) {
// 	history, err := u.storage.GetHistoryByID(ctx, id)
// 	if err != nil {
// 		log.Printf("Error getting history by ID: %v", err)
// 		return nil, err
// 	}
// 	return history, nil
// }

// func (u *Usecase) ListHistories(ctx context.Context, userID int) (*[]History, error) {
// 	histories, err := u.storage.ListHistories(ctx, userID)
// 	if err != nil {
// 		log.Printf("Error listing histories: %v", err)
// 		return nil, err
// 	}
// 	return histories, nil
// }

// func (u *Usecase) UpdateHistory(ctx context.Context, history History) error {
// 	if err := u.storage.UpdateHistory(ctx, history); err != nil {
// 		log.Printf("Error updating history: %v", err)
// 		return err
// 	}
// 	return nil
// }

// func (u *Usecase) DeleteHistory(ctx context.Context, id int) error {
// 	if err := u.storage.DeleteHistory(ctx, id); err != nil {
// 		log.Printf("Error deleting history: %v", err)
// 		return err
// 	}
// 	return nil
// }
