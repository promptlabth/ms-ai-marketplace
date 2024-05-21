package framework

import (
	"context"
	"encoding/json"
)

type Framework struct {
	ID   string
	Name string
	Detail string 
	Component json.RawMessage
}

type FrameworkInterface interface {
	CreateFramework(ctx context.Context, agent_Detail Framework) (*string, error) 
	GetFrameworkByID(ctx context.Context, id string) (*Framework, error)     
	ListFrameworks(ctx context.Context) (*[]Framework, error)     
	UpdateFramework(ctx context.Context, agent_Detail Framework) error      
	DeleteFramework(ctx context.Context, id string) error                         
}

type NewFrameworkRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Detail string `json:"detail"`
	Component json.RawMessage `json:"component"`
}
