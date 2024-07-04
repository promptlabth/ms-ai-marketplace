package history

import (
	"context"
	"time"
)

type History struct {
	ID             int     
	UserID         string      
	AgentID        int     
	FrameworkID    int       
	Prompt         string   
	StyleMessageID int      
	Language       string      
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
	UserID         string    `json:"user_id"` 
	AgentID        int    `json:"agent_id"`
	FrameworkID    int    `json:"framework_id"`
	Prompt         string `json:"prompt"`
	StyleMessageID int    `json:"style_message_id"`
}
