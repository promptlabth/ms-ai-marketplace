package generate

import (
	"context"

	agentdetail "github.com/promptlabth/ms-ai-marketplace/app/agent_detail"
	"github.com/promptlabth/ms-ai-marketplace/app/history"
)

type generateAdaptor interface {
	GenerateAdaptor(ctx context.Context, prompt, model string) (string, error)
}

type agentStorage interface {
	GetAgentByID(context.Context, int) (*agentdetail.AgentDetailEntity, error)
}

type historyStorage interface {
	CreateHistory(ctx context.Context, history history.HistoryEntity) (*int, error)
}
