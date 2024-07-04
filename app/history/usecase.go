package history

import (
	"context"
	"errors"

	// "fmt"
	"log"
	"math/rand"

	// "os"
	"time"

	agentdetail "github.com/promptlabth/ms-orch-user-service/app/agent_detail"
)


type storage interface {
	CreateHistory(ctx context.Context, history HistoryEntity) (*int, error)
	GetAgentByID(ctx context.Context, id int) (*agentdetail.AgentDetailEntity, error)
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
	// Validate the new history
	if err := u.domain.ValidateNewHistory(ctx, history); err != nil {
		log.Printf("Validation error: %v", err)
		return nil,err
	}

	// Check if UserID exists
	if !u.domain.existsInDatabase(ctx, "users", history.UserID) {
		return nil,errors.New("User ID does not exist")
	}

	// Get Agent by ID
	agent, err := u.storage.GetAgentByID(ctx, history.AgentID)
	if err != nil {
		log.Printf("Error fetching agent: %v", err)
		return nil, errors.New("Failed to fetch agent. Please try again later.")
	}

	// Check if FrameworkID exists
	if !u.domain.existsInDatabase(ctx, "frameworks", string(history.FrameworkID)) {
		return nil,errors.New("Framework ID does not exist")
	}

	// Check if StyleMessageID exists (optional)
	if history.StyleMessageID != 0 && !u.domain.existsInDatabase(ctx, "style_prompts", string(history.StyleMessageID)) {
		return nil,errors.New("Style Message ID does not exist")
	}

	// Initialize the generate service
	// generateService := GenerateService()

	// Randomly select a model with specified weights
	modelLanguageChoices := []string{"GPT", "VERTEX"}
	// weights := []float64{0.5, 0.3, 0.2}

	// In development environment, prioritize VERTEX
	// if os.Getenv("ENV") == "DEV" {
	// 	modelLanguageChoices = []string{"VERTEX"}
	// 	weights = []float64{1.0}
	// }

	var result string
	// var err error

	// Attempt to generate a text message using one of the models
	for len(modelLanguageChoices) > 0 {
		// modelLanguage := randomChoice(modelLanguageChoices, weights)
		// switch modelLanguage {
		// case "GPT":
		// 	result, err = generateWithModel(u, generateService, history, "GPT")
		// case "VERTEX":
		// 	result, err = generateWithModel(u, generateService, history, "VERTEX")
		// }

		// if err == nil {
		// 	break
		// }

		// Handle model failure by adjusting weights
		// index := findIndex(modelLanguageChoices, modelLanguage)
		// if index >= 0 {
		// 	data := weights[index]
		// 	weights = append(weights[:index], weights[index+1:]...)
		// 	if len(weights) > 0 {
		// 		weights[0] += data
		// 	}
		// 	modelLanguageChoices = append(modelLanguageChoices[:index], modelLanguageChoices[index+1:]...)
		// }
	}

	if err != nil || result == "" || agent == nil {
		log.Printf("Error generating message: %v", err)
		return nil,errors.New("Failed to generate message. Please try again later.")
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
	return nil,err
}

// Helper function to generate message using specified model
// func generateWithModel(u *Usecase, generateService GenerateService, history History, modelName string) (string, error) {
	// model, err := u.domain.modelUsecase.get_by_name(modelName)
	// if err != nil {
	// 	return "", err
	// }

	// dbPrompt, err := u.domain.inputPromptUsecase.get_by_feature_id_and_model_id(history.FeatureID, model.ID, history.LanguageID)
	// if err != nil {
	// 	return "", err
	// }

	// if dbPrompt == nil {
	// 	return "", errors.New("prompt not found")
	// }

	// inputPrompt := fmt.Sprintf(dbPrompt.PromptInput, history.Prompt, history.StyleMessageID)
	// switch modelName {
	// case "GPT":
	// 	return generateService.generateMessageOpenAI(inputPrompt)
	// case "GEMINI":
	// 	return generateService.generateMessageGeminiAI(inputPrompt, history.StyleMessageID)
	// default:
	// 	return "", errors.New("unknown model")
	// }
// }

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