package history

import (
	"time"
)

// History represents the history of interactions with the framework
type HistoryEntity struct {
	ID             int             `gorm:"autoIncrement;column:id"`
	UserID         string             `gorm:"column:user_id"`                
	AgentID        int             `gorm:"column:agent_id"`
	FrameworkID    int             `gorm:"column:framework_id"`
	Prompt         string          `gorm:"column:prompt"`
	StyleMessageID int             `gorm:"column:style_message_id"`
	Language     string             `gorm:"column:language_id"`
	Result         string          `gorm:"column:result"`
	TimeStamp      time.Time       `gorm:"column:time_stamp"`
}


func (HistoryEntity) TableName() string {
	return "histories"
}
