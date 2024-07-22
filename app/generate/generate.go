package generate

type GenerateRequest struct {
	FirebaseID     string `json:"firebase_id"`
	AgentID        int    `json:"agent_id"`
	FrameworkID    int    `json:"framework_id"`
	Prompt         string `json:"prompt"`
	StyleMessageID int    `json:"style_message_id"`
}
