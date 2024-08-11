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
	var agentDetail []AgentDetailEntity
	if err := c.db.Where("user_id = ?", firebaseId).Find(&agentDetail).Error; err != nil {
		return nil, err
	}

	return &agentDetail, nil
}

func (c *Core) GetAgentByID(ctx context.Context, id int) (*AgentDetailEntity, error) {
	var agent AgentDetailEntity
	if err := c.db.First(&agent, id).Error; err != nil {
		return nil, err	
	}
	return &agent, nil
}

func (c *Core) ListAgentDetails(ctx context.Context) (*[]AgentDetailEntity, error) {
	var agents []AgentDetailEntity
	if err := c.db.Find(&agents).Error; err != nil {
		return nil, err
	}
	return &agents, nil
}

func (c *Core) UpdateAgentDetail(ctx context.Context, agentDetail AgentDetailEntity) ( error) {
	var existingAgent AgentDetailEntity
	if err := c.db.First(&existingAgent, agentDetail.ID).Error; err != nil {
		return err
	}
	if err := c.db.Model(&existingAgent).Updates(agentDetail).Error; err != nil {
		return err
	}
	return nil
}