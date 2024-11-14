package realtimegen

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

func (c *Core) GetFullPromptByAgentID(ctx context.Context, agentId int) (*RealtimegenFull, error) {
	var fullPrompt RealtimegenFull
	if err := c.db.Table("agent_details").
		Select(`agent_details.id as agent_id, agent_details.name as agent_name, agent_details.description as agent_description, agent_details.image_url as agent_image_url, 
			agent_details.prompt as agent_prompt, agent_details.firebase_id as agent_firebase_id, agent_details.framework_id as agent_framework_id, 
			agent_details.role_framework_id as agent_role_framework_id, agent_details.total_used as agent_total_used,
			roles.id as role_id, roles.name as role_name, roles.language as role_language,
			frameworks.id as framework_id, frameworks.name as framework_name, frameworks.detail as framework_detail, 
			frameworks.component as framework_component, frameworks.language as framework_language, frameworks.prompt as framework_prompt`).
		Joins("left join roles on roles.id = agent_details.role_framework_id").
		Joins("left join frameworks on frameworks.id = agent_details.framework_id").
		Where("agent_details.id = ?", agentId).
		Scan(&fullPrompt).Error; err != nil {
		return nil, err
	}

	return &fullPrompt, nil
}