package history

import (
	"context"
	"log"
	// "errors"
	// // "fmt"

	// "log"
	// "math/rand"
	// "strings"
	// "text/template"
	// // "os"
	"time"

	"github.com/promptlabth/ms-orch-user-service/app/__mock__/role"
	agentdetail "github.com/promptlabth/ms-orch-user-service/app/agent_detail"
	generateservice "github.com/promptlabth/ms-orch-user-service/app/external_service"
	"github.com/promptlabth/ms-orch-user-service/app/framework"
	styleprompt "github.com/promptlabth/ms-orch-user-service/app/style_prompt"
)

type storage interface {
	CreateHistory(ctx context.Context, history HistoryEntity) (*int, error)
	GetAgentByID(ctx context.Context, id int) (*agentdetail.AgentDetailEntity, error)
	GetFrameworkByID(ctx context.Context, id int) (*framework.FrameworkEntity, error)
	GetStyleMessageByID(ctx context.Context, id int) (*styleprompt.StylePromptEntity, error)
	GetRoleByID(ctx context.Context, id int) (*role.RoleEntity, error)
	// GetHistoryByID(ctx context.Context, id int) (*History, error)
	// ListHistories(ctx context.Context, userID int) (*[]History, error)
	// UpdateHistory(ctx context.Context, history History) error
	// DeleteHistory(ctx context.Context, id int) error
}

type domain interface {
	ValidateNewHistory(ctx context.Context, history History) error
	// existsInDatabase(ctx context.Context, tableName string, id string) bool
}

type Usecase struct {
	storage storage
	domain  domain
}

func NewUsecase(s storage, d domain) *Usecase {
	return &Usecase{
		storage: s,
		domain:  d,
	}
}
func (u *Usecase) CreateHistory(ctx context.Context, history History) (*string, string) {
	var result = "result"

	err := u.domain.ValidateNewHistory(ctx, history)
	if err != nil {
		return nil, "validation error: " + err.Error()
	}

	
	

	historyEntity := HistoryEntity{
		UserID:         history.UserID,
		AgentID:        history.AgentID,
		FrameworkID:    history.FrameworkID,
		Prompt:         history.Prompt,
		StyleMessageID: history.StyleMessageID,
		Language:       history.Language,
		Result:         result,
		TimeStamp:      time.Now(),
	}

	

	generateService := generateservice.NewGenerate()
	resultOpenAI, err := generateService.GenerateMessageOpenAI("Hi, What is OpenAI")

	if err != nil {
		log.Printf("Error generating message with OpenAI: %v", err)
		return nil, "Error generating message with OpenAI: " + err.Error()
	}

	result = result+resultOpenAI
	_, err = u.storage.CreateHistory(ctx, historyEntity)
	if err != nil {
		return nil, "storage error: " + err.Error()
	}

	return &result, ""
}

// func (u *Usecase) CreateHistory(ctx context.Context, history History) (*string, string) {
// // Validate the new history
// if err := u.domain.ValidateNewHistory(ctx, history); err != nil {
// 	return nil,err
// }

// // Check if UserID exists
// if !u.domain.existsInDatabase(ctx, "users", history.UserID) {
// 	return nil,err
// }

// Get Agent by ID
// agent, err := u.storage.GetAgentByID(ctx, history.AgentID)
// if err != nil {
// 	return nil, "err GetAgentByID"
// }

// // Check if FrameworkID exists
// framework, err := u.storage.GetFrameworkByID(ctx, history.FrameworkID)
// if err != nil {
// 	log.Printf("Error fetching framework: %v", err)
// 	return nil, err
// }
// fmt.Print(framework)

// Check if StyleMessageID exists (optional)
// styleMessage, err := u.storage.GetStyleMessageByID(ctx, history.StyleMessageID)
// if err != nil {
// 	return nil, "err GetStyleMessageByID"
// }
// fmt.Print(styleMessage)

// // // Check if role exists (optional)
// role, err := u.storage.GetRoleByID(ctx, 1)
// if err != nil {
// 	return nil, "err GetRoleByID"
// }
// fmt.Print(role)

// Generate framework detail from agent.Prompt
// var frameworkDetail strings.Builder
// for key, value := range agent.Prompt {
// 	frameworkDetail.WriteString(fmt.Sprintf("%s is %s ", key, value))
// }

// Prepare the input prompt
// inputPrompt := " Provide guidance in the role of {{docter}} which includes {{Safety}} needing an answer in the style of happy language th"
// inputPromptTemplate := " Provide guidance in the role of {{.role}} which includes {{.frameworkDetail}} needing an answer in the style of {{.styleMessage}} language {{.language}}"
// inputPrompt, err := formatInputPrompt(inputPromptTemplate, role, frameworkDetail.String(), styleMessage.Name, history)
// if err != nil {
// 	return nil, "err formatInputPrompt"
// }

// result, err :=  handleModelGeneration(inputPrompt);
// if err != nil {
// 	return nil, "err handleModelGeneration"
// }
// 	var result = "result"

// 	historyEntity := HistoryEntity{
// 		UserID:         history.UserID,
// 		AgentID:        history.AgentID,
// 		FrameworkID:    history.FrameworkID,
// 		Prompt:         history.Prompt,
// 		StyleMessageID: history.StyleMessageID,
// 		Language:       history.Language,
// 		Result:         result,
// 		TimeStamp:      time.Now(),
// 	}

// 	_, err = u.storage.CreateHistory(ctx, historyEntity)
// 	if err != nil {
// 		return nil, "u CreateHistory"
// 	}

// 	return &result,""
// }

// func handleModelGeneration(imputPromtp string) (string, error) {
// 	generateService := generateservice.Generate{}

// 	modelLanguageChoices := []string{"GIMINI", "GIMINI"}
// 	// modelLanguageChoices := []string{"GPT", "GIMINI"}
// 	weights := []float64{0.6, 0.4}

// 	// In development environment, prioritize VERTEX
// 	// if os.Getenv("ENV") == "DEV" {
// 	// 	modelLanguageChoices = []string{"VERTEX"}
// 	// 	weights = []float64{1.0}
// 	// }

// 	for len(modelLanguageChoices) > 0 {
// 		modelLanguage := randomChoice(modelLanguageChoices, weights)

// 		switch modelLanguage {
// 		case "GIMINI":
// 			result, err := generateService.GenerateMessageGemini(imputPromtp)
// 			if err != nil {
// 				log.Printf("Error generating message: %v", err)
// 				// Remove the failing model from the list and adjust weights
// 				index := findIndex(modelLanguageChoices, modelLanguage)
// 				modelLanguageChoices = append(modelLanguageChoices[:index], modelLanguageChoices[index+1:]...)
// 				weights = append(weights[:index], weights[index+1:]...)
// 				continue
// 			}
// 			return result, nil
// 		default:
// 			return "", errors.New("unsupported model")
// 		}
// 	}
// 	return "", errors.New("all models failed")
// }

// func formatInputPrompt(templateStr string, role *role.RoleEntity, frameworkDetail, styleMessage string, history History) (string, error) {
// 	tmpl, err := template.New("inputPrompt").Parse(templateStr)
// 	if err != nil {
// 		return "", err
// 	}

// 	var sb strings.Builder
// 	err = tmpl.Execute(&sb, map[string]interface{}{
// 		"role":            role,
// 		"frameworkDetail": frameworkDetail,
// 		"styleMessage":    styleMessage,
// 		"language":        history.Language,
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	return sb.String(), nil
// }
// // Helper function to randomly select a model based on weights
// func randomChoice(choices []string, weights []float64) string {
// 	sum := 0.0
// 	for _, w := range weights {
// 		sum += w
// 	}

// 	r := rand.Float64() * sum
// 	for i, w := range weights {
// 		r -= w
// 		if r <= 0 {
// 			return choices[i]
// 		}
// 	}

// 	return choices[len(choices)-1]
// }

// // Helper function to find the index of an element in a slice
// func findIndex(slice []string, element string) int {
// 	for i, v := range slice {
// 		if v == element {
// 			return i
// 		}
// 	}
// 	return -1
// }
