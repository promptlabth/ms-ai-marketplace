package history

import (
	"context"
	"time"
)

type storage interface {
	CreateHistory(ctx context.Context, history HistoryEntity) (*int, error)
	GetHistoryByFirebaseID(ctx context.Context, firebaseID string) ([]HistoryEntity, error)
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
func (u *Usecase) CreateHistory(ctx context.Context, history History) error {

	err := u.domain.ValidateNewHistory(ctx, history)
	if err != nil {
		return err
	}

	historyEntity := HistoryEntity{
		FirebaseID:        history.FirebaseID,
		AgentID:           history.AgentID,
		FrameworkID:       history.FrameworkID,
		Prompt:            history.Prompt,
		StyleMessageID:    history.StyleMessageID,
		Result:            history.Result,
		Model:             history.Model,
		Completion_tokens: history.Completion_tokens,
		Prompt_tokens:     history.Prompt_tokens,
		Language:          history.Language,
		TimeStamp:         time.Now(),
	}

	_, err = u.storage.CreateHistory(ctx, historyEntity)
	return err
}

func (u *Usecase) GetHistoryByFirebaseID(ctx context.Context, firebaseID string) ([]History, error) {
	histories, err := u.storage.GetHistoryByFirebaseID(ctx, firebaseID)
	if err != nil {
		return nil, err
	}

	var result []History
	for _, h := range histories {
		result = append(result, History{
			ID:                h.ID,
			FirebaseID:        h.FirebaseID,
			AgentID:           h.AgentID,
			FrameworkID:       h.FrameworkID,
			StyleMessageID:    h.StyleMessageID,
			Language:          h.Language,
			TimeStamp:         h.TimeStamp,
		})
	}

	return result, nil
}
