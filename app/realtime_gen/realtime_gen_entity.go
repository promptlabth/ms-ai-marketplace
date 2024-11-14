package realtimegen

import "encoding/json"

type RealtimegenFull struct {
	// Fields from AgentDetailEntity
	AgentID          int             `gorm:"column:agent_id"`
	AgentName        string          `gorm:"column:agent_name"`
	AgentDescription string          `gorm:"column:agent_description"`
	AgentImageURL    string          `gorm:"column:agent_image_url"`
	AgentPrompt      json.RawMessage `gorm:"column:agent_prompt"`
	AgentFirebaseID  string          `gorm:"column:agent_firebase_id"`
	AgentFrameworkID int             `gorm:"column:agent_framework_id"`
	AgentRoleFrameID int             `gorm:"column:agent_role_framework_id"`
	AgentTotalUsed   int             `gorm:"column:agent_total_used"`

	// Fields from RoleEntity
	RoleID       int    `gorm:"column:role_id"`
	RoleName     string `gorm:"column:role_name"`
	RoleLanguage string `gorm:"column:role_language"`

	// Fields from FrameworkEntity
	FrameworkID        int             `gorm:"column:framework_id"`
	FrameworkName      string          `gorm:"column:framework_name"`
	FrameworkDetail    string          `gorm:"column:framework_detail"`
	FrameworkComponent json.RawMessage `gorm:"column:framework_component"`
	FrameworkLanguage  string          `gorm:"column:framework_language"`
	FrameworkPrompt    string          `gorm:"column:framework_prompt"`
}