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
		roleLanguage = "ตอบภาษาไทยเท่านั้น"
	} else if roleLanguage == "en" {
		roleLanguage = "answer in english only"
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
			fullPrompt.FrameworkPrompt,
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
		// Extract JSON fields from AgentPrompt
		var agentPrompt map[string]string
		if err := json.Unmarshal(fullPrompt.AgentPrompt, &agentPrompt); err != nil {
			return nil, err
		}

        result.FullPrompt = fmt.Sprintf(
			fullPrompt.FrameworkPrompt,
			fullPrompt.RoleName,
			agentPrompt["task"],
			agentPrompt["goal"],
			"user_input", // Placeholder for user frontend input
			"style_prompt", // Placeholder for user frontend input
			roleLanguage,
		)
	} else if fullPrompt.FrameworkID == 4 || fullPrompt.FrameworkID == 9 {
		// Extract JSON fields from AgentPrompt
		var agentPrompt map[string]string
		if err := json.Unmarshal(fullPrompt.AgentPrompt, &agentPrompt); err != nil {
			return nil, err
		}

        result.FullPrompt = fmt.Sprintf(
			fullPrompt.FrameworkPrompt,
			fullPrompt.RoleName,
			agentPrompt["action"],
			agentPrompt["expectation"],
			"user_input", // Placeholder for user frontend input
			"style_prompt", // Placeholder for user frontend input
			roleLanguage,
		)
	} else if fullPrompt.FrameworkID == 5 || fullPrompt.FrameworkID == 10 {
		// Extract JSON fields from AgentPrompt
		var agentPrompt map[string]string
		if err := json.Unmarshal(fullPrompt.AgentPrompt, &agentPrompt); err != nil {
			return nil, err
		}

        result.FullPrompt = fmt.Sprintf(
			fullPrompt.FrameworkPrompt,
			fullPrompt.RoleName,
			agentPrompt["problem"],
			agentPrompt["promise"],
            agentPrompt["prove"],
			"user_input", // Placeholder for user frontend input
            agentPrompt["proposal"],
			"style_prompt", // Placeholder for user frontend input
			roleLanguage,
		)
	} else {
		// Perform default actions
		result.FullPrompt = fmt.Sprintf("Default: %s - %s", fullPrompt.AgentName, fullPrompt.FrameworkName)
	}

	return &result, nil
}
