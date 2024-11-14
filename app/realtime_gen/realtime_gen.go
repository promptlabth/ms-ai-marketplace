package realtimegen

import "context"

type RealTimeGenPrompt struct {
	FullPrompt string `json:"FullPrompt"`
}

type RealTimeGenPromptInterface interface {
	GetRealTimePromptByID(ctx context.Context, id int) (*RealTimeGenPrompt, error)
}
