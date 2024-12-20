package agentdetail

import "encoding/json"

type AgentDetailEntity struct {
    ID          int             `gorm:"autoIncrement;column:id"`
    Name        string          `gorm:"column:name;uniqueIndex"`
    Description string          `gorm:"column:description"`
    ImageURL    string          `gorm:"column:image_url"`
    Prompt      json.RawMessage `gorm:"column:prompt"`
    FirebaseID  string          `gorm:"column:firebase_id"`
    FrameworkID int             `gorm:"column:framework_id"`
    RoleFrameID int             `gorm:"column:role_framework_id"`
    TotalUsed   int             `gorm:"column:total_used"`
    Status      string          `gorm:"column:status"`
    Language    string          `gorm:"-"`
    Role        RoleEntity      `gorm:"foreignKey:RoleFrameID;references:ID"`
}

func (AgentDetailEntity) TableName() string {
    return "agent_details"
}

type RoleEntity struct {
    ID       int    `gorm:"autoIncrement;column:id"`
    Language string `gorm:"column:language"`
}

func (RoleEntity) TableName() string {
    return "roles"
}
