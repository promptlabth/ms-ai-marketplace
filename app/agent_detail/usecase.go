// usecase.go

package agentdetail

import (
    "context"
    "log"
    "github.com/promptlabth/ms-ai-marketplace/app/payment/coins"
)

type storage interface {
    CreateAgentDetail(context.Context, AgentDetailEntity) (*int, error)
    GetAgentDetailsByUserID(context.Context, string) (*[]AgentDetailEntity, error)
    ListAgentDetails(context.Context) (*[]AgentDetailEntity, error)
    ListAgentDetailsThatApprove(context.Context) (*[]AgentDetailEntity, error)
    GetAgentByID(context.Context, int) (*AgentDetailEntity, error)
    UpdateAgentDetail(context.Context, AgentDetailEntity) error
    IncrementTotalUsed(context.Context, int) error
    UpdateAgentStatus(context.Context, int, string) error
    DeleteAgentDetail(context.Context, int) error
}

type domain interface {
    ValidateNewAgentDetail(ctx context.Context, agentDetail AgentDetail) error
}

type Usecase struct {
    storage      storage
    domain       domain
    coinsUsecase coins.Usecase
}

func NewUsecase(s storage, d domain, cu coins.Usecase) *Usecase {
    return &Usecase{
        storage:      s,
        domain:       d,
        coinsUsecase: cu,
    }
}

func (u *Usecase) NewAgentDetail(ctx context.Context, agentDetail AgentDetail) error {
    agentDetailEntity := AgentDetailEntity{
        Name:        agentDetail.Name,
        Description: agentDetail.Description,
        ImageURL:    agentDetail.ImageURL,
        Prompt:      agentDetail.Prompt,
        FirebaseID:  agentDetail.FirebaseID,
        FrameworkID: agentDetail.FrameworkID,
        RoleFrameID: agentDetail.RoleFrameID,
        TotalUsed:   agentDetail.TotalUsed,
        Status:      agentDetail.Status,
        Language:    agentDetail.Language,
    }
    log.Printf("AgentDetailEntity : %+v\n", agentDetailEntity)

    // Create the agent detail
    agentID, err := u.storage.CreateAgentDetail(ctx, agentDetailEntity)
    if err != nil {
        return err
    }

    // Create coins for the agent detail
    coinsEntity := coins.CoinsEntity{
        FirebaseID: agentDetail.FirebaseID,
        AgentID:    *agentID, // Set the AgentID to the newly created agent's ID
        Coins:      0, // Set default coins to zero
    }
    _, err = u.coinsUsecase.CreateCoins(ctx, coinsEntity)
    if err != nil {
        return err
    }

    return nil
}

func (u *Usecase) GetAgentDetails(ctx context.Context, firebaseId string) (*[]AgentDetailEntity, error) {
    agentDetail, err := u.storage.GetAgentDetailsByUserID(ctx, firebaseId)
    if err != nil {
        return nil, err
    }
    return agentDetail, nil
}

func (u *Usecase) GetAgentByID(ctx context.Context, id int) (*AgentDetailEntity, error) {
    agent, err := u.storage.GetAgentByID(ctx, id)
    if err != nil {
        log.Printf("Error getting agent by ID: %v", err)
        return nil, err
    }
    return agent, nil
}

func (u *Usecase) ListAgentDetails(ctx context.Context) (*[]AgentDetailEntity, error) {
    agents, err := u.storage.ListAgentDetails(ctx)
    if err != nil {
        return nil, err
    }
    return agents, nil
}

func (u *Usecase) ListAgentDetailsThatApprove(ctx context.Context) (*[]AgentDetailEntity, error) {
    agents, err := u.storage.ListAgentDetailsThatApprove(ctx)
    if err != nil {
        return nil, err
    }
    return agents, nil
}

func (u *Usecase) UpdateAgentDetail(ctx context.Context, agentDetail AgentDetail) error {
    agentDetailEntity := AgentDetailEntity{
        ID:          agentDetail.ID,
        Name:        agentDetail.Name,
        Description: agentDetail.Description,
        ImageURL:    agentDetail.ImageURL,
        Prompt:      agentDetail.Prompt,
        FirebaseID:  agentDetail.FirebaseID,
        FrameworkID: agentDetail.FrameworkID,
        RoleFrameID: agentDetail.RoleFrameID,
        TotalUsed:   agentDetail.TotalUsed,
        Status:      "pending",
    }

    err := u.storage.UpdateAgentDetail(ctx, agentDetailEntity)
    return err
}

func (u *Usecase) IncrementTotalUsed(ctx context.Context, agentID int) error {
    return u.storage.IncrementTotalUsed(ctx, agentID)
}

func (u *Usecase) UpdateAgentStatus(ctx context.Context, agentID int, status string) error {
    return u.storage.UpdateAgentStatus(ctx, agentID, status)
}

func (u *Usecase) DeleteAgentDetail(ctx context.Context, agentID int) error {
    return u.storage.DeleteAgentDetail(ctx, agentID)
}
