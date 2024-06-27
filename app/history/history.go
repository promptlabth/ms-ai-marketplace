package history

import (
	"context"
	"time"
)

type History struct {
	ID             int     
	UserID         int      
	AgentID        int     
	FrameworkID    int       
	Prompt         string   
	StyleMessageID int      
	LanguageID     int      
	Result         string 
	TimeStamp      time.Time 
}

type HistoryInterface interface {
	CreateHistory(ctx context.Context, historyDetail History) (*int, error)
	GetHistoryByID(ctx context.Context, id int) (*History, error)
	ListHistories(ctx context.Context, userID int) (*[]History, error)
	// UpdateHistory(ctx context.Context, historyDetail History) error
	// DeleteHistory(ctx context.Context, id int) error
}

type NewHistoryRequest struct {
	UserID         int    `json:"user_id"` 
	AgentID        int    `json:"agent_id"`
	FrameworkID    int    `json:"framework_id"`
	Prompt         string `json:"prompt"`
	StyleMessageID int    `json:"style_message_id"`
	LanguageID     int    `json:"language_id"`
	Result         string `json:"result"`
}
