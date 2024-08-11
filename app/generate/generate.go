package generate

type GenerateRequest struct {
	FirebaseID     string `json:"firebase_id"`
	AgentID        int    `json:"agent_id"`
	Prompt         string `json:"prompt"`
	StyleMessageID int    `json:"style_message_id"`
}


type PromptAPE struct {
	Propose     string `json:"propose"`
	Expectation string `json:"expectation"`
}
type PromptRICEE struct {
	Instruction string `json:"instruction"`
	Context     string `json:"context"`
	Example     string `json:"example"`
	Execute     string `json:"execute"`
}

type PromptTAG struct {
	Task string `json:"task"`
	Goal string `json:"goal"`
}

type PromptERA struct {
	Expectation string `json:"expectation"`
	Action      string `json:"action"`
}

type PromptRPPPP struct {
	Problem  string `json:"problem"`
	Promise  string `json:"promise"`
	Prove    string `json:"prove"`
	Proposal string `json:"proposal"`
}
