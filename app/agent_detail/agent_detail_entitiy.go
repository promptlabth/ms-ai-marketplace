package agentdetail

import "encoding/json"

type AgentDetailEntity struct {
	AgentDetailID string `gorm:"column:agent_detail_id"`
	Name        string          `gorm:"column:name"`
	Description string          `gorm:"column:descriptoin"`
	ImageURL    string          `gorm:"column:image_url"`
	Prompt      json.RawMessage `gorm:"column:prompt"` // for a raw json
	UserID      int64           `gorm:"column:user_id"`
	FrameworkID int64           `gorm:"column:framework_id"`
	RoleFrameID int64           `gorm:"column:role_framework_id"`
}

func (AgentDetailEntity) TableName() string {
	return "agent_Details"
}
