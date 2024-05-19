// core.go

package agentdetail

import (
	"context"

	"gorm.io/gorm"
)

type Core struct {
	db *gorm.DB
}

func NewCore(db *gorm.DB) *Core {
	return &Core{db: db}
}

func (c *Core) CreateAgentDetail(ctx context.Context, agentDetail AgentDetailEntity) (*string, error) {

	if err := c.db.Create(&agentDetail); err.Error != nil {
		return nil, err.Error
	}
	return &agentDetail.AgentDetailID, nil
}

//can add Delete Update Get
