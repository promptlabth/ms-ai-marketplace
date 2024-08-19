package generate

import (
	"context"

	agentdetail "github.com/promptlabth/ms-ai-marketplace/app/agent_detail"
	"github.com/promptlabth/ms-ai-marketplace/app/framework"
	"github.com/promptlabth/ms-ai-marketplace/app/history"
	"github.com/promptlabth/ms-ai-marketplace/app/role"
	styleprompt "github.com/promptlabth/ms-ai-marketplace/app/style_prompt"
)

type generateAdaptor interface {
	GenerateAdaptor(ctx context.Context, prompt, model string) (string, error)
}

type agentStorage interface {
	GetAgentByID(ctx context.Context, id int) (*agentdetail.AgentDetailEntity, error)
	UpdateAgentDetail(ctx context.Context, agentDetail agentdetail.AgentDetailEntity) error
}
type stylepromptStorage interface {
	GetStylePromptByID(ctx context.Context, id int) (*styleprompt.StylePromptEntity, error)
}
type frameworkStorage interface {
	GetFrameworkByID(ctx context.Context, id int) (*framework.FrameworkEntity, error)
}
type roleStorage interface {
	GetRoleByID(ctx context.Context, id int) (*role.RoleEntity, error)
}
type historyStorage interface {
	CreateHistory(ctx context.Context, history history.HistoryEntity) (*int, error)
}

//secondary port
type storage interface {
	Generate(ctx context.Context, prompt string,model string) (string,int,int, error)
}
