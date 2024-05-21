package agentdetail

import "encoding/json"

type AgentDetailEntity struct {
	ID          string          `gorm:"column:agent_detail_id;uniqueIndex"`
	Name        string          `gorm:"column:name;uniqueIndex"`
	Description string          `gorm:"column:description"`
	ImageURL    string          `gorm:"column:image_url"`
	Prompt      json.RawMessage `gorm:"column:prompt"` // for a raw json
	UserID      string          `gorm:"column:user_id"`
	FrameworkID string          `gorm:"column:framework_id"`
	RoleFrameID string          `gorm:"column:role_framework_id"`
}

func (AgentDetailEntity) TableName() string {
	return "agent_details"
}
