package realtimegen

import (
	"context"
	"encoding/json"
	"fmt"
)

type storage interface {
	GetFullPromptByAgentID(ctx context.Context, agentId int) (*RealtimegenFull, error)
}

type Usecase struct {
	storage storage
}

func NewUsecase(s storage) *Usecase {
	return &Usecase{
		storage: s,
	}
}

func (u *Usecase) GetFullPromptByAgentID(ctx context.Context, id int) (*RealTimeGenPrompt, error) {
	fullPrompt, err := u.storage.GetFullPromptByAgentID(ctx, id)
	if err != nil {
		return nil, err
	}

	roleLanguage := fullPrompt.RoleLanguage
	if roleLanguage == "th" {
		roleLanguage = "Thai"
	} else if roleLanguage == "en" {
		roleLanguage = "english"
	}

	var result RealTimeGenPrompt

	if fullPrompt.FrameworkID == 1 || fullPrompt.FrameworkID == 6 { //Framework APE
		// Extract JSON fields from AgentPrompt
		var agentPrompt map[string]string
		if err := json.Unmarshal(fullPrompt.AgentPrompt, &agentPrompt); err != nil {
			return nil, err
		}

		// Format the prompt string
		result.FullPrompt = fmt.Sprintf(
			fullPrompt.FrameworkPrompt,
			fullPrompt.RoleName,
			agentPrompt["propose"],
			agentPrompt["expectation"],
			"user_input",   // Placeholder for user frontend input
			"style_prompt", // Placeholder for user frontend input
			roleLanguage,
		)
	} else if fullPrompt.FrameworkID == 2 || fullPrompt.FrameworkID == 7 {
		// Extract JSON fields from AgentPrompt
		var agentPrompt map[string]string
		if err := json.Unmarshal(fullPrompt.AgentPrompt, &agentPrompt); err != nil {
			return nil, err
		}

		// Format the prompt string for FrameworkID 2 and 7
		result.FullPrompt = fmt.Sprintf(
			"Your view as [%s] and your task is [%s]. I will expect you to [%s] about [%s]. Example is [%s]. Execute on [%s]. The article should feel like [%s] in [%s] language.",
			fullPrompt.RoleName,
			agentPrompt["context"],
			agentPrompt["instruction"],
			"user_input", // Placeholder for user frontend input
			agentPrompt["example"],
			agentPrompt["execute"],
			"style_prompt", // Placeholder for user frontend input
			roleLanguage,
		)
	} else if fullPrompt.FrameworkID == 3 || fullPrompt.FrameworkID == 8 {
		// Perform specific actions for FrameworkID 3
		result.FullPrompt = fmt.Sprintf("Framework 3: %s - %s", fullPrompt.AgentName, fullPrompt.FrameworkName)
	} else if fullPrompt.FrameworkID == 4 || fullPrompt.FrameworkID == 9 {
		// Perform specific actions for FrameworkID 4
		result.FullPrompt = fmt.Sprintf("Framework 4: %s - %s", fullPrompt.AgentName, fullPrompt.FrameworkName)
	} else if fullPrompt.FrameworkID == 5 || fullPrompt.FrameworkID == 10 {
		// Perform specific actions for FrameworkID 5
		result.FullPrompt = fmt.Sprintf("Framework 5: %s - %s", fullPrompt.AgentName, fullPrompt.FrameworkName)
	} else {
		// Perform default actions
		result.FullPrompt = fmt.Sprintf("Default: %s - %s", fullPrompt.AgentName, fullPrompt.FrameworkName)
	}

	return &result, nil
}
