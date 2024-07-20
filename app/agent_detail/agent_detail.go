// agent_detail.go

package agentdetail

import (
	"context"
	"encoding/json"
)

type AgentDetail struct {
	ID          int64
	Name        string
	Description string
	ImageURL    string
	Prompt      json.RawMessage
	FirebaseID      string
	FrameworkID int64
	RoleFrameID int64
	TotalUsed int64
}

type AgentDetailInterface interface {
	CreateAgentDetail(ctx context.Context, agent_Detail AgentDetail) (*int64, error) // Creates a new user and returns the user ID
	GetAgentDetailByID(ctx context.Context, id string) (*AgentDetail, error)         // Fetches a user by their ID
	GetAgentByID(ctx context.Context, id int) (*AgentDetail, error)
	// UpdateAgentDetail(ctx context.Context, agent_Detail AgentDetail) error          // Updates an existing user
	// DeleteAgentDetail(ctx context.Context, id string) error           // Deletes a user by their ID
	ListAgentDetails(ctx context.Context) (*[]AgentDetail, error)
}

type NewAgentDetailRequest struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	ImageURL    string          `json:"image_url"`
	Prompt      json.RawMessage `json:"prompt"`
	FirebaseID      string          `json:"firebase_id"`
	FrameworkID int64           `json:"framework_id"`
	RoleFrameID int64           `json:"role_framework_id"`
	TotalUsed int64           `json:"total_used"`
	// Include other fields as necessary
}
