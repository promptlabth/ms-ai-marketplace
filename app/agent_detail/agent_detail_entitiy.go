package agentdetail

import "encoding/json"

type AgentDetailEntity struct {
    ID          int64           `gorm:"autoIncrement;column:id"`
    Name        string          `gorm:"column:name;uniqueIndex"`
    Description string          `gorm:"column:description"`
    ImageURL    string          `gorm:"column:image_url"`
    Prompt      json.RawMessage `gorm:"column:prompt"` 
    UserID      string          `gorm:"column:user_id"`
    FrameworkID int64           `gorm:"column:framework_id"`
    RoleFrameID int64           `gorm:"column:role_framework_id"`
}

func (AgentDetailEntity) TableName() string {
	return "agent_details"
}
