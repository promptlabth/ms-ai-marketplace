package review

import (
	"context"

	"github.com/promptlabth/ms-ai-marketplace/app/agent_detail"
)

type storage interface {
	CreateReview(context.Context, ReviewEntity) (*int, error)
}

type agentDetailStorage interface {
	GetAgentByID(ctx context.Context, id int) (*agentdetail.AgentDetailEntity, error)
}

type Usecase struct {
	storage           storage
	agentDetailStorage agentDetailStorage
}

func NewUsecase(s storage, ads agentDetailStorage) *Usecase {
	return &Usecase{
		storage:           s,
		agentDetailStorage: ads,
	}
}

func (u *Usecase) NewReview(ctx context.Context, review ReviewEntity) error {
    agentDetail, err := u.agentDetailStorage.GetAgentByID(ctx, review.AgentID)
    if err != nil {
        return err
    }

    review.UserID = agentDetail.FirebaseID

    _, err = u.storage.CreateReview(ctx, review)
    return err
}
