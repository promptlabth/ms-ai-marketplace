package history

import (
	"context"
	"time"
)

type storage interface {
	CreateHistory(ctx context.Context, history HistoryEntity) (*int,error)
}

type domain interface {
	ValidateNewHistory(ctx context.Context, history History) error
}

type Usecase struct {
	storage      storage
	domain       domain
}

func NewUsecase(s storage, d domain, ) *Usecase {
	return &Usecase{
		storage:      s,
		domain:       d,
	}
}
func (u *Usecase) CreateHistory(ctx context.Context, history History)  error {

	err := u.domain.ValidateNewHistory(ctx, history)
	if err != nil {
		return err
	}

	historyEntity := HistoryEntity{
		FirebaseID:     history.FirebaseID,
		AgentID:        history.AgentID,
		FrameworkID:    history.FrameworkID,
		Prompt:         history.Prompt,
		StyleMessageID: history.StyleMessageID,
		Language:       history.Language,
		Result:         history.Result,
		TimeStamp:      time.Now(),
	}

	_, err = u.storage.CreateHistory(ctx, historyEntity)
	return err
}
