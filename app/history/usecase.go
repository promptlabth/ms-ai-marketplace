package history

import (
	"context"
	"errors"
	"fmt"

	// "fmt"
	"log"
	"math/rand"
	"strings"
	"text/template"
	// "os"
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
	existsInDatabase(ctx context.Context, tableName string, id string) bool
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

func (u *Usecase) CreateHistory(ctx context.Context, history History) (*string, error) {
	// // Validate the new history
	// if err := u.domain.ValidateNewHistory(ctx, history); err != nil {
	// 	return nil,err
	// }

	// // Check if UserID exists
	// if !u.domain.existsInDatabase(ctx, "users", history.UserID) {
	// 	return nil,err
	// }

	// Get Agent by ID
	agent, err := u.storage.GetAgentByID(ctx, history.AgentID)
	if err != nil {
		return nil, err
	}

	// // Check if FrameworkID exists
	// framework, err := u.storage.GetFrameworkByID(ctx, history.FrameworkID)
	// if err != nil {
	// 	log.Printf("Error fetching framework: %v", err)
	// 	return nil, err
	// }
	// fmt.Print(framework)

	// Check if StyleMessageID exists (optional)
	styleMessage, err := u.storage.GetStyleMessageByID(ctx, history.StyleMessageID)
	if err != nil {
		return nil, err
	}
	// Check if role exists (optional)
	role, err := u.storage.GetRoleByID(ctx, int(agent.RoleFrameID))
	if err != nil {
		return nil, err
	}
	
	// Generate framework detail from agent.Prompt
	var frameworkDetail strings.Builder
	for key, value := range agent.Prompt {
		frameworkDetail.WriteString(fmt.Sprintf("%s is %s ", key, value))
	}

	// Prepare the input prompt
	inputPromptTemplate := " Provide guidance in the role of {{.role}} which includes {{.frameworkDetail}} needing an answer in the style of {{.styleMessage}} language {{.language}}"
	inputPrompt, err := formatInputPrompt(inputPromptTemplate, role, frameworkDetail.String(), styleMessage.Name, history)
	if err != nil {
		return nil, err
	}

	result, err :=  handleModelGeneration(inputPrompt);
	if err != nil {
		return nil, err
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

	_, err = u.storage.CreateHistory(ctx, historyEntity)
	if err != nil {
		return nil, err
	}
	
	return &result,nil
}

func handleModelGeneration(imputPromtp string) (string, error) {
	generateService := generateservice.Generate{}

	modelLanguageChoices := []string{"GIMINI", "GIMINI"} 
	// modelLanguageChoices := []string{"GPT", "GIMINI"} 
	weights := []float64{0.6, 0.4}        
	
	// In development environment, prioritize VERTEX
	// if os.Getenv("ENV") == "DEV" {
	// 	modelLanguageChoices = []string{"VERTEX"}
	// 	weights = []float64{1.0}
	// }

	for len(modelLanguageChoices) > 0 {
		modelLanguage := randomChoice(modelLanguageChoices, weights)

		switch modelLanguage {
		case "GIMINI":
			result, err := generateService.GenerateMessageGemini(imputPromtp)
			if err != nil {
				log.Printf("Error generating message: %v", err)
				// Remove the failing model from the list and adjust weights
				index := findIndex(modelLanguageChoices, modelLanguage)
				modelLanguageChoices = append(modelLanguageChoices[:index], modelLanguageChoices[index+1:]...)
				weights = append(weights[:index], weights[index+1:]...)
				continue
			}
			return result, nil
		default:
			return "", errors.New("unsupported model")
		}
	}
	return "", errors.New("all models failed")
}

func formatInputPrompt(templateStr string, role *role.RoleEntity, frameworkDetail, styleMessage string, history History) (string, error) {
	tmpl, err := template.New("inputPrompt").Parse(templateStr)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	err = tmpl.Execute(&sb, map[string]interface{}{
		"role":            role,
		"frameworkDetail": frameworkDetail,
		"styleMessage":    styleMessage,
		"language":        history.Language,
	})
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}
// Helper function to randomly select a model based on weights
func randomChoice(choices []string, weights []float64) string {
	sum := 0.0
	for _, w := range weights {
		sum += w
	}

	r := rand.Float64() * sum
	for i, w := range weights {
		r -= w
		if r <= 0 {
			return choices[i]
		}
	}

	return choices[len(choices)-1]
}

// Helper function to find the index of an element in a slice
func findIndex(slice []string, element string) int {
	for i, v := range slice {
		if v == element {
			return i
		}
	}
	return -1
}


// // CreateHistory orchestrates the process of validating and creating a new history record
// func (u *Usecase) CreateHistory(ctx context.Context, history History) (*string, error) {
// 	// Optional: Validate the new history using domain logic
// 	if err := u.domain.ValidateNewHistory(ctx, history); err != nil {
// 		log.Printf("Validation error: %v", err)
// 		return nil,err
// 	}
// 	result, err := u.generateMessage(history)
// 	if err != nil {
// 		log.Printf("Error generating message: %v", err)
// 		return nil, err
// 	}
// 	historyEntity := HistoryEntity{
// 		UserID:         history.UserID,
// 		AgentID:        history.AgentID,
// 		FrameworkID:    history.ID,
// 		Prompt:         history.Prompt,
// 		StyleMessageID: history.StyleMessageID,
// 		LanguageID:     history.LanguageID,
// 		Result:         result,
// 		TimeStamp:      time.Now(),
// 	}

// 	_, err := u.storage.CreateHistory(ctx, historyEntity)
// 	if err != nil {
// 		log.Printf("Error getting history by ID: %v", err)
// 		return nil, err
// 	}

// 	return result,err
// }


// func (u *Usecase) GetHistoryByID(ctx context.Context, id int) (*History, error) {
// 	history, err := u.storage.GetHistoryByID(ctx, id)
// 	if err != nil {
// 		log.Printf("Error getting history by ID: %v", err)
// 		return nil, err
// 	}
// 	return history, nil
// }

// func (u *Usecase) ListHistories(ctx context.Context, userID int) (*[]History, error) {
// 	histories, err := u.storage.ListHistories(ctx, userID)
// 	if err != nil {
// 		log.Printf("Error listing histories: %v", err)
// 		return nil, err
// 	}
// 	return histories, nil
// }

// func (u *Usecase) UpdateHistory(ctx context.Context, history History) error {
// 	if err := u.storage.UpdateHistory(ctx, history); err != nil {
// 		log.Printf("Error updating history: %v", err)
// 		return err
// 	}
// 	return nil
// }

// func (u *Usecase) DeleteHistory(ctx context.Context, id int) error {
// 	if err := u.storage.DeleteHistory(ctx, id); err != nil {
// 		log.Printf("Error deleting history: %v", err)
// 		return err
// 	}
// 	return nil
// }