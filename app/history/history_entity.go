package history

import (
	"encoding/json"
	"time"
)

type HistoryEntity struct {
	ID                int       `gorm:"autoIncrement;column:id"`
	FirebaseID        string    `gorm:"column:firebase_id"`
	AgentID           int       `gorm:"column:agent_id"`
	FrameworkID       int       `gorm:"column:framework_id"`
	Prompt            string    `gorm:"column:prompt"`
	StyleMessageID    int       `gorm:"column:style_message_id"`
	Language          string    `gorm:"column:language"`
	Result            string    `gorm:"column:result"`
	Model             string    `gorm:"column:model"`
	Completion_tokens int       `gorm:"column:completion_tokens"`
	Prompt_tokens     int       `gorm:"column:prompt_tokens"`
	TimeStamp         time.Time `gorm:"column:time_stamp"`
}

type HistoryWithAgentDetail struct {
    ID                int           `gorm:"column:id"`
    FirebaseID        string        `gorm:"column:firebase_id"`
    AgentID           int           `gorm:"column:agent_id"`
    FrameworkID       int           `gorm:"column:framework_id"`
    Prompt            string        `gorm:"column:prompt"`
    StyleMessageID    int           `gorm:"column:style_message_id"`
    Language          string        `gorm:"column:language"`
    Result            string        `gorm:"column:result"`
    Model             string        `gorm:"column:model"`
    Completion_tokens int           `gorm:"column:completion_tokens"`
    Prompt_tokens     int           `gorm:"column:prompt_tokens"`
    TimeStamp         time.Time     `gorm:"column:time_stamp"`
    Name              string        `gorm:"column:name"`
    Description       string        `gorm:"column:description"`
    ImageURL          string        `gorm:"column:image_url"`
    AgentPrompt       json.RawMessage `gorm:"column:prompt"`
    AgentFrameworkID  int           `gorm:"column:framework_id"`
    RoleFrameID       int           `gorm:"column:role_framework_id"`
    TotalUsed         int           `gorm:"column:total_used"`
}

func (HistoryEntity) TableName() string {
	return "histories"
}
