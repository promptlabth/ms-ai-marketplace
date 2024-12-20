// core.go

package agentdetail

import (
	"context"

	"gorm.io/gorm"
)

type Core struct {
	db *gorm.DB
}

func NewCore(db *gorm.DB) *Core {
	return &Core{db: db}
}

func (c *Core) CreateAgentDetail(ctx context.Context, agentDetail AgentDetailEntity) (*int, error) {
	if err := c.db.Create(&agentDetail); err.Error != nil {
		return nil, err.Error
	}
	return &agentDetail.ID, nil
}

func (c *Core) GetAgentDetailsByUserID(ctx context.Context, firebaseId string) (*[]AgentDetailEntity, error) {
	var agentDetails []AgentDetailEntity
	if err := c.db.Preload("Role").Where("firebase_id = ?", firebaseId).Find(&agentDetails).Error; err != nil {
		return nil, err
	}
	for i := range agentDetails {
		agentDetails[i].Language = agentDetails[i].Role.Language
	}
	return &agentDetails, nil
}

func (c *Core) GetAgentByID(ctx context.Context, id int) (*AgentDetailEntity, error) {
	var agent AgentDetailEntity
	if err := c.db.Preload("Role").Where("id = ?", id).First(&agent).Error; err != nil {
		return nil, err
	}
	agent.Language = agent.Role.Language
	return &agent, nil
}

func (c *Core) ListAgentDetails(ctx context.Context) (*[]AgentDetailEntity, error) {
	var agents []AgentDetailEntity
	if err := c.db.Preload("Role").Find(&agents).Error; err != nil {
		return nil, err
	}
	for i := range agents {
		agents[i].Language = agents[i].Role.Language
	}
	return &agents, nil
}

func (c *Core) ListAgentDetailsThatApprove(ctx context.Context) (*[]AgentDetailEntity, error) {
    var agents []AgentDetailEntity
    if err := c.db.Preload("Role").Where("status = ?", "approve").Find(&agents).Error; err != nil {
        return nil, err
    }
    for i := range agents {
        agents[i].Language = agents[i].Role.Language
    }
    return &agents, nil
}

func (c *Core) UpdateAgentDetail(ctx context.Context, agentDetail AgentDetailEntity) error {
	var existingAgent AgentDetailEntity
	if err := c.db.First(&existingAgent, agentDetail.ID).Error; err != nil {
		return err
	}
	if err := c.db.Model(&existingAgent).Updates(agentDetail).Error; err != nil {
		return err
	}
	return nil
}

func (c *Core) IncrementTotalUsed(ctx context.Context, agentID int) error {
    return c.db.Model(&AgentDetailEntity{}).Where("id = ?", agentID).Update("total_used", gorm.Expr("total_used + ?", 1)).Error
}

func (c *Core) UpdateAgentStatus(ctx context.Context, agentID int, status string) error {
    return c.db.Model(&AgentDetailEntity{}).Where("id = ?", agentID).Update("status", status).Error
}