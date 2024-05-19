// usecase.go

package agentdetail

import (
	"context"
	"log"
)

type storage interface {
	CreateAgentDetail(context.Context, AgentDetailEntity) (*string, error)
	GetAgentDetailsByUserID(context.Context, string) (*[]AgentDetailEntity, error) 
}

type domain interface {
	ValidateNewAgentDetail(ctx context.Context, agentDetail AgentDetail) error
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

	// First, use the domain logic to validate the new role.
	if err := u.domain.ValidateNewAgentDetail(ctx, agentDetail); err != nil {
		log.Printf("Error validating new role: %v", err)
		return err
	}

	agentDetailEntity := AgentDetailEntity{
		ID: agentDetail.ID,
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

func (u *Usecase) GetAgentDetails(ctx context.Context, firebaseId string) (*[]AgentDetailEntity, error) {
	agentDetail, err := u.storage.GetAgentDetailsByUserID(ctx, firebaseId)
    if err != nil {
        return nil, err
    }
    return agentDetail, nil
}
