package history

import (
	"context"
	"time"
)

type History struct {
	ID                int
	FirebaseID        string
	AgentID           int
	FrameworkID       int
	Prompt            string
	StyleMessageID    int
	Language          string
	Result            string
	Model             string
	Completion_tokens int
	Prompt_tokens     int
	TimeStamp         time.Time
}

type HistoryInterface interface {
	CreateHistory(ctx context.Context, historyDetail History) (*int, error)
	GetHistoryByID(ctx context.Context, id int) (*History, error)
	ListHistories(ctx context.Context, userID int) (*[]History, error)
}

type NewHistoryRequest struct {
	FirebaseID        string `json:"firebase_id"`
	AgentID           int    `json:"agent_id"`
	FrameworkID       int    `json:"framework_id"`
	Prompt            string `json:"prompt"`
	StyleMessageID    int    `json:"style_message_id"`
	Language          string `json:"language"`
	Result            string `json:"result"`
	Model             string `json:"model"`
	Completion_tokens int    `json:"completion_tokens"`
	Prompt_tokens     int    `json:"prompt_tokens"`
}
