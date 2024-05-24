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

func (c *Core) CreateAgentDetail(ctx context.Context, agentDetail AgentDetailEntity) (*int64, error) {

	if err := c.db.Create(&agentDetail); err.Error != nil {
		return nil, err.Error
	}
	return &agentDetail.ID, nil
}

func (c *Core) GetAgentDetailsByUserID(ctx context.Context, firebaseId string) (*[]AgentDetailEntity, error){
	var agentDetail []AgentDetailEntity
	if err := c.db.Where("user_id = ?", firebaseId).Find(&agentDetail).Error; err != nil {
		return nil, err
	}

	return &agentDetail, nil
}

//can add Delete Update
// UpdateRole updates a role's information in the database.
// func (c *Core) UpdateRole(ctx context.Context, agentDetail AgentDetailEntity) error {
// 	return c.db.Model(AgentDetailEntity{}).Where("id = ?", agentDetail.ID).Updates(agentDetail).Error
// }

// // DeleteRole removes a role from the database by their ID.
// func (c *Core) DeleteRole(ctx context.Context, id int) error {
// 	return c.db.Delete(&AgentDetailEntity{}, id).Error
// }
