package review

import "time"

type ReviewEntity struct {
	ID       int       `gorm:"autoIncrement;column:id"`
	AdminID  string    `gorm:"column:admin_id"`
	UserID   string    `gorm:"column:user_id"`
	AgentID  int       `gorm:"column:agent_id"`
	Reason   string    `gorm:"column:reason"`
	DateTime time.Time `gorm:"column:date_time"`
}

type ReviewRequest struct {
	AdminID  string    `json:"admin_id"`
	AgentID  int       `json:"agent_id"`
	Reason   string    `json:"reason"`
	DateTime time.Time `json:"date_time"`
}

func (ReviewEntity) TableName() string {
	return "reviews"
}
