package generate

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	agentdetail "github.com/promptlabth/ms-ai-marketplace/app/agent_detail"
	"github.com/promptlabth/ms-ai-marketplace/app/history"
)

type GenerateService struct {
	generateAdaptor    generateAdaptor
	agentStorage       agentStorage
	stylepromptStorage stylepromptStorage
	frameworkStorage   frameworkStorage
	roleStorage        roleStorage
	historyStorage     historyStorage
	storage            storage
}

func NewService(
	generateAdaptor generateAdaptor,
	agentStorage agentStorage,
	stylepromptStorage stylepromptStorage,
	frameworkStorage frameworkStorage,
	roleStorage roleStorage,
	historyStorage historyStorage,
	storage storage,

) *GenerateService {
	return &GenerateService{
		generateAdaptor:    generateAdaptor,
		agentStorage:       agentStorage,
		historyStorage:     historyStorage,
		stylepromptStorage: stylepromptStorage,
		roleStorage:        roleStorage,
		frameworkStorage:   frameworkStorage,
		storage:            storage,
	}
}

func (s *GenerateService) Generate(ctx context.Context, generateRequest GenerateRequest, language string) (string, error) {

	agent, err := s.agentStorage.GetAgentByID(ctx, generateRequest.AgentID)
	if err != nil {
		return "", err
	}
	// var agent_result agentdetail.AgentDetailEntity

	// if tx := entity.DB().Where("id = ?", agentDetail.ID).First(&agent_result); tx.RowsAffected == 0 {

	// 	return fmt.Errorf("error menu not found")
	// }

	stylePrompt, err := s.stylepromptStorage.GetStylePromptByID(ctx, generateRequest.StyleMessageID)
	if err != nil {
		return "", err
	}

	framework, err := s.frameworkStorage.GetFrameworkByID(ctx, agent.FrameworkID)
	if err != nil {
		return "", err
	}
	role, err := s.roleStorage.GetRoleByID(ctx, agent.RoleFrameID)
	if err != nil {
		return "", err
	}

	promptData, err := getPromptdata(agent.Prompt, framework.Name)
	if err != nil {
		return "", err
	}

	promptMessage, err := getPromptMessage(promptData, role.Name, generateRequest.Prompt, stylePrompt.Name, language,framework.Prompt)
	if err != nil {
		return "", err
	}
	fmt.Print("promptMessage: %s", promptMessage)

	model := "SeaLLM-7B-v2.5"
	// promptMessage = "Your view as [Doctor] and your task is [talk with ผู้ป่วย]. I will expect you to [ผู้ป่วย halp full] that article should feel like [funny] in th language."
	message,completion_tokens,prompt_tokens, err := s.storage.Generate(ctx, promptMessage, model)
	if err != nil {
		return "", err
	}

	history := history.HistoryEntity{
		FirebaseID:     generateRequest.FirebaseID,
		AgentID:        agent.ID,
		FrameworkID:    framework.ID,
		Prompt:         generateRequest.Prompt,
		StyleMessageID: stylePrompt.ID,
		Result:         message,
		Model:             model,
		Completion_tokens: completion_tokens,
		Prompt_tokens:     prompt_tokens,
		Language:       language,
		TimeStamp:      time.Now(),
	}

	id, err := s.historyStorage.CreateHistory(ctx, history)
	if err != nil {
		return "", err
	}

	fmt.Sprintf("History created with ID: %d", *id)
// ("name","description","image_url","prompt","firebase_id","framework_id","role_framework_id","total_used")
	agentDetailEntity := agentdetail.AgentDetailEntity{
		ID:          agent.ID,
		Name:        agent.Name,
		Description: agent.Description,
		ImageURL:    agent.ImageURL,
		Prompt:      agent.Prompt,
		FirebaseID:  agent.FirebaseID,
		FrameworkID: agent.FrameworkID,
		RoleFrameID: agent.RoleFrameID,
		TotalUsed:   agent.TotalUsed + 1,
	}
	if err := s.agentStorage.UpdateAgentDetail(ctx, agentDetailEntity); err != nil {
		return "", err
	}
	return message, nil
}

func getPromptdata(promptJSON json.RawMessage, nameFramework string) (interface{}, error) {
	var promptData interface{}

	switch nameFramework {
	case "RICEE":
		var data PromptRICEE
		if err := json.Unmarshal(promptJSON, &data); err != nil {
			return nil, err
		}
		promptData = data
	case "APE":
		var data PromptAPE
		if err := json.Unmarshal(promptJSON, &data); err != nil {
			return nil, err
		}
		promptData = data
	case "TAG":
		var data PromptTAG
		if err := json.Unmarshal(promptJSON, &data); err != nil {
			return nil, err
		}
		promptData = data
	case "ERA":
		var data PromptERA
		if err := json.Unmarshal(promptJSON, &data); err != nil {
			return nil, err
		}
		promptData = data
	case "RPPPP":
		var data PromptRPPPP
		if err := json.Unmarshal(promptJSON, &data); err != nil {
			return nil, err
		}
		promptData = data
	default:
		return nil, fmt.Errorf("unknown framework type")
	}

	return promptData, nil
}

func getPromptMessage(promptData interface{}, role, propmt_input, styleName, language string,promptFormat string) (string, error) {
	var message string
	
	switch data := promptData.(type) {
	case PromptRICEE:
		message = fmt.Sprintf(
			promptFormat,
			role, data.Context, data.Instruction, propmt_input, data.Example, data.Execute, styleName, language,
		)
	case PromptAPE:
		message = fmt.Sprintf(
			promptFormat,
			role, data.Propose, data.Expectation, propmt_input, styleName, language,
		)
	case PromptTAG:
		message = fmt.Sprintf(
			promptFormat,
			role, data.Task, data.Goal, propmt_input, styleName, language,
		)
	case PromptERA:
		message = fmt.Sprintf(
			promptFormat,
			role, data.Action, data.Expectation, propmt_input, styleName, language,
		)
	case PromptRPPPP:
		message = fmt.Sprintf(
			promptFormat,
			role, data.Problem, data.Promise, data.Prove, propmt_input, data.Proposal, styleName, language,
		)
	default:
		return "", fmt.Errorf("unsupported prompt data type")
	}
	return message, nil
}
