package history

import (
	"context"
	"regexp"
	// "strings"
	"time"

	"github.com/promptlabth/ms-ai-marketplace/app/payment/coins"
)

type storage interface {
	CreateHistory(ctx context.Context, history HistoryEntity) (*int, error)
	GetHistoryByFirebaseID(ctx context.Context, firebaseID string) ([]HistoryWithAgentDetail, error)
}

type domain interface {
	ValidateNewHistory(ctx context.Context, history History) error
}

type Usecase struct {
	storage     storage
	domain      domain
	coinsUsecase coins.Usecase
}

func NewUsecase(s storage, d domain, cu coins.Usecase) *Usecase {
	return &Usecase{
		storage:     s,
		domain:      d,
		coinsUsecase: cu,
	}
}

func countTokens(text string) int {
    // Define a regular expression to match words and punctuation marks
    re := regexp.MustCompile(`\w+|[^\w\s]`)
    // Find all matches and return the count
    return len(re.FindAllString(text, -1))
}
func (u *Usecase) CreateHistory(ctx context.Context, history History) error {
	err := u.domain.ValidateNewHistory(ctx, history)
	if err != nil {
		return err
	}

	// Calculate tokens
	history.Completion_tokens = countTokens(history.Result)
	history.Prompt_tokens = countTokens(history.Prompt)

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
	if err != nil {
		return err
	}

	// Increment coins for the agent
	err = u.coinsUsecase.IncreaseCoinsByFirebaseIDAndAgentID(ctx, history.FirebaseID, history.AgentID)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) GetHistoryByFirebaseID(ctx context.Context, firebaseID string) ([]HistoryWithAgentDetail, error) {
	histories, err := u.storage.GetHistoryByFirebaseID(ctx, firebaseID)
	if err != nil {
		return nil, err
	}

	var result []HistoryWithAgentDetail
	for _, h := range histories {
		result = append(result, HistoryWithAgentDetail{
			ID:                h.ID,
			AgentID:           h.AgentID,
			FrameworkID:       h.FrameworkID,
			Prompt:            h.Prompt,
			StyleMessageID:    h.StyleMessageID,
			Language:          h.Language,
			Result:            h.Result,
			Model:             h.Model,
			TimeStamp:         h.TimeStamp,
			Name:              h.Name,
			Description:       h.Description,
			ImageURL:          h.ImageURL,
			AgentFrameworkID:  h.AgentFrameworkID,
			RoleFrameID:       h.RoleFrameID,
			TotalUsed:         h.TotalUsed,
		})
	}

	return result, nil
}
