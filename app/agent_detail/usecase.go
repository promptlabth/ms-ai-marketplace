// usecase.go

package agentdetail

import (
	"context"
	"log"
)

type storage interface {
	CreateAgentDetail(context.Context, AgentDetailEntity) (*string, error)
}

type domain interface {
	// ValidateNewUser(ctx context.Context, agent_detail AgentDetail) error
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

func (u *Usecase) NewAgentDetail(ctx context.Context, agentDetail AgentDetail) error {

	agentDetailEntity := AgentDetailEntity{
		AgentDetailID: agentDetail.AgentDetailID,
		Name:          agentDetail.Name,
		Description:   agentDetail.Description,
		ImageURL:      agentDetail.ImageURL,
		Prompt:        agentDetail.Prompt,
		UserID:        agentDetail.UserID,
		FrameworkID:   agentDetail.FrameworkID,
		RoleFrameID:   agentDetail.RoleFrameID,
	}
	log.Printf("AgentDetailEntity : %+v\n", agentDetailEntity)

	_, err := u.storage.CreateAgentDetail(ctx, agentDetailEntity)
	return err
}
