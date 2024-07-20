package history

import (
	"context"
	"errors"
	"log"
	// "fmt"
	"math/rand"
	// "strings"
	// "text/template"
	// "os"
	"time"

	agentdetail "github.com/promptlabth/ms-ai-marketplace/app/agent_detail"
	generateservice "github.com/promptlabth/ms-ai-marketplace/app/external_service"
	"github.com/promptlabth/ms-ai-marketplace/app/framework"
	"github.com/promptlabth/ms-ai-marketplace/app/role"
	styleprompt "github.com/promptlabth/ms-ai-marketplace/app/style_prompt"
)

type storage interface {
	CreateHistory(ctx context.Context, history HistoryEntity) (*int, error)
	GetAgentByID(ctx context.Context, id int) (*agentdetail.AgentDetailEntity, error)
	GetFrameworkByID(ctx context.Context, id int) (*framework.FrameworkEntity, error)
	GetStyleMessageByID(ctx context.Context, id int) (*styleprompt.StylePromptEntity, error)
	GetRoleByID(ctx context.Context, id int) (*role.RoleEntity, error)
}

type domain interface {
	ValidateNewHistory(ctx context.Context, history History) error
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

	err := u.domain.ValidateNewHistory(ctx, history)
	if err != nil {
		return nil, "validation error: " + err.Error()
	}

	result, err := handleModelGeneration(history.Prompt)
	if err != nil {
		return nil, "handleModelGeneration error: " + err.Error()
	}

	historyEntity := HistoryEntity{
		FirebaseID:     history.FirebaseID,
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
		return nil, "storage error: " + err.Error()
	}

	return &result, ""
}

func handleModelGeneration(imputPromtp string) (string, error) {
	generateService, err := generateservice.NewGenerateService()
	if err != nil {
		return "", err
	}
	modelLanguageChoices := []string{"GIMINI", "GIMINI"}
	// modelLanguageChoices := []string{"GPT", "GIMINI"}
	weights := []float64{0.6, 0.4}

	// In development environment, prioritize VERTEX
	// if os.Getenv("ENV") == "DEV" {
	// 	modelLanguageChoices = []string{"VERTEX"}
	// 	weights = []float64{1.0}
	// }
	var error_log = ""
	for len(modelLanguageChoices) > 0 {
		modelLanguage := randomChoice(modelLanguageChoices, weights)
		log.Printf("modelLanguage %v", modelLanguage)

		switch modelLanguage {
		case "GIMINI":
			result, err := generateService.GenerateMessageVertexAI(imputPromtp, "APE")
			if err != nil {
				error_log = err.Error()
				log.Printf("Error generating message: %v", err)
				// Remove the failing model from the list and adjust weights
				index := findIndex(modelLanguageChoices, modelLanguage)
				modelLanguageChoices = append(modelLanguageChoices[:index], modelLanguageChoices[index+1:]...)
				weights = append(weights[:index], weights[index+1:]...)
				continue
			}
			return result, nil
		case "GPT":
			result, err := generateService.GenerateMessageOpenAI(imputPromtp)
			if err != nil {
				error_log = err.Error()
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
	return "", errors.New(error_log)
}

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

//		return sb.String(), nil
//	}
//
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
