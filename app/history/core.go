package history

import (
	"context"

	agentdetail "github.com/promptlabth/ms-ai-marketplace/app/agent_detail"
	"github.com/promptlabth/ms-ai-marketplace/app/framework"
	"github.com/promptlabth/ms-ai-marketplace/app/role"
	styleprompt "github.com/promptlabth/ms-ai-marketplace/app/style_prompt"
	"gorm.io/gorm"
)

// Core is responsible for the data storage operations related to histories.
type Core struct {
	db *gorm.DB
}

// NewCore creates a new instance of Core with a database connection.
func NewCore(db *gorm.DB) *Core {
	return &Core{db: db}
}

// CreateHistory inserts a new history record into the database.
func (c *Core) CreateHistory(ctx context.Context, history HistoryEntity) (*int, error) {
	if err := c.db.Create(&history).Error; err != nil {
		return nil, err
	}
	return &history.ID, nil
}

func (c *Core) GetAgentByID(ctx context.Context, id int) (*agentdetail.AgentDetailEntity, error) {
	var agent agentdetail.AgentDetailEntity
	if err := c.db.First(&agent, id).Error; err != nil {
		return nil, err
	}
	return &agent, nil
}
func (c *Core) GetFrameworkByID(ctx context.Context, id int) (*framework.FrameworkEntity, error) {
	var framework framework.FrameworkEntity
	if err := c.db.First(&framework, id).Error; err != nil {
		return nil, err
	}
	return &framework, nil
}
func (c *Core) GetStyleMessageByID(ctx context.Context, id int) (*styleprompt.StylePromptEntity, error) {
	var styleMessage styleprompt.StylePromptEntity
	if err := c.db.First(&styleMessage, id).Error; err != nil {
		return nil, err
	}
	return &styleMessage, nil
}
func (c *Core) GetRoleByID(ctx context.Context, id int) (*role.RoleEntity, error) {
	var role role.RoleEntity
	if err := c.db.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
