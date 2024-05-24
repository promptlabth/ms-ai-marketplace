// agent_detail.go

package agentdetail

import (
	"context"
	"encoding/json"
)

type AgentDetail struct {
	ID string
	Name          string
	Description   string
	ImageURL      string
	Prompt        json.RawMessage
	UserID        string
	FrameworkID   int64
	RoleFrameID   int64
}

// UserInterface defines the set of methods that any implementation of the User service must provide.
type AgentDetailInterface interface {
	CreateAgentDetail(ctx context.Context, agent_Detail AgentDetail) (*string, error) // Creates a new user and returns the user ID
	GetAgentDetailByID(ctx context.Context, id string) (*AgentDetail, error) // Fetches a user by their ID
	// UpdateAgentDetail(ctx context.Context, agent_Detail AgentDetail) error          // Updates an existing user
	// DeleteAgentDetail(ctx context.Context, id string) error           // Deletes a user by their ID
}

type NewAgentDetailRequest struct {
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	ImageURL      string          `json:"image_url"`
	Prompt        json.RawMessage `json:"prompt"`
	UserID        string           `json:"user_id"`
	FrameworkID   int64           `json:"framework_id"`
	RoleFrameID   int64           `json:"role_frame_id"`
	// Include other fields as necessary
}