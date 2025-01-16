package history

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"time"

	"github.com/promptlabth/ms-ai-marketplace/app/agent_detail"
	"github.com/promptlabth/ms-ai-marketplace/app/payment/coins"
	"github.com/promptlabth/ms-ai-marketplace/app/user/service"
)

type storage interface {
	CreateHistory(ctx context.Context, history HistoryEntity) (*int, error)
	GetHistoryByFirebaseID(ctx context.Context, firebaseID string) ([]HistoryWithAgentDetail, error)
	GetHistoriesByAgentIDs(ctx context.Context, agentIDs []int) ([]HistoryWithAgentDetail, error)
}

type domain interface {
	ValidateNewHistory(ctx context.Context, history History) error
}

type Usecase struct {
	storage      storage
	domain       domain
	coinsUsecase coins.Usecase
	agentDetail  agentdetail.Core
	userService  service.UserService // Add userService to Usecase
}

func NewUsecase(s storage, d domain, cu coins.Usecase, ad agentdetail.Core, us service.UserService) *Usecase {
	return &Usecase{
		storage:      s,
		domain:       d,
		coinsUsecase: cu,
		agentDetail:  ad,
		userService:  us,
	}
}

func countTokens(text string) int {
	// Define a regular expression to match words and punctuation marks
	re := regexp.MustCompile(`\w+|[^\w\s]`)
	// Find all matches and return the count
	return len(re.FindAllString(text, -1))
}

func (u *Usecase) CreateHistory(ctx context.Context, history History) error {
	err := u.domain.ValidateNewHistory(ctx, history)
	if err != nil {
		return err
	}

	// Calculate tokens
	history.Completion_tokens = countTokens(history.Result)
	history.Prompt_tokens = countTokens(history.Prompt)

	historyEntity := HistoryEntity{
		FirebaseID:        history.FirebaseID,
		AgentID:           history.AgentID,
		FrameworkID:       history.FrameworkID,
		Prompt:            history.Prompt,
		StyleMessageID:    history.StyleMessageID,
		Result:            history.Result,
		Model:             history.Model,
		Completion_tokens: history.Completion_tokens,
		Prompt_tokens:     history.Prompt_tokens,
		Language:          history.Language,
		TimeStamp:         time.Now(),
	}

	_, err = u.storage.CreateHistory(ctx, historyEntity)
	if err != nil {
		return err
	}

	// Increment coins for the agent
	err = u.coinsUsecase.IncreaseCoinsByFirebaseIDAndAgentID(ctx, history.FirebaseID, history.AgentID)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) GetHistoryByFirebaseID(ctx context.Context, firebaseID string) ([]HistoryWithAgentDetail, error) {
	histories, err := u.storage.GetHistoryByFirebaseID(ctx, firebaseID)
	if err != nil {
		return nil, err
	}

	var result []HistoryWithAgentDetail
	for _, h := range histories {
		result = append(result, HistoryWithAgentDetail{
			ID:                h.ID,
			AgentID:           h.AgentID,
			FrameworkID:       h.FrameworkID,
			Prompt:            h.Prompt,
			StyleMessageID:    h.StyleMessageID,
			Language:          h.Language,
			Result:            h.Result,
			Model:             h.Model,
			TimeStamp:         h.TimeStamp,
			Name:              h.Name,
			Description:       h.Description,
			ImageURL:          h.ImageURL,
			AgentFrameworkID:  h.AgentFrameworkID,
			RoleFrameID:       h.RoleFrameID,
			TotalUsed:         h.TotalUsed,
		})
	}

	return result, nil
}

func (u *Usecase) GetHistoriesByFirebaseID(ctx context.Context, firebaseID string) ([]HistoryWithAgentDetail, error) {
	// Get the list of Agent IDs from firebaseID
	agentDetails, err := u.agentDetail.GetAgentDetailsByUserID(ctx, firebaseID)
	if err != nil {
		return nil, err
	}

	// Extract the Agent IDs
	var agentIDs []int
	for _, agent := range *agentDetails {
		agentIDs = append(agentIDs, agent.ID)
	}

	// Query the histories table using the list of Agent IDs
	histories, err := u.storage.GetHistoriesByAgentIDs(ctx, agentIDs)

	if err != nil {
		return nil, err
	}
	fmt.Printf("Histories: %+v\n", histories)

	return histories, nil

}

func (u *Usecase) GetHistoriesByFirebaseID_Extracted(ctx context.Context, firebaseID string) ([]AgentUsageResponse, error) {
	// Get the list of Agent IDs from firebaseID
	agentDetails, err := u.agentDetail.GetAgentDetailsByUserID(ctx, firebaseID)
	if err != nil {
		return nil, err
	}

	// Extract the Agent IDs
	var agentIDs []int
	for _, agent := range *agentDetails {
		agentIDs = append(agentIDs, agent.ID)
	}

	// Query the histories table using the list of Agent IDs
	histories, err := u.storage.GetHistoriesByAgentIDs(ctx, agentIDs)
	if err != nil {
		return nil, err
	}

	// Aggregate the data
	agentUsageMap := make(map[int]map[string]int)
	agentDetailsMap := make(map[int]HistoryWithAgentDetail)
	for _, history := range histories {
		if _, exists := agentUsageMap[history.AgentID]; !exists {
			agentUsageMap[history.AgentID] = make(map[string]int)
			agentDetailsMap[history.AgentID] = history
		}
		agentUsageMap[history.AgentID][history.FirebaseID]++
	}

	// Format the response
	var response []AgentUsageResponse
	for agentID, userUsageMap := range agentUsageMap {
		var userUsageList []UserUsage
		totalUsage := 0
		for firebaseID, usageCount := range userUsageMap {
			user, err := u.userService.GetUser(firebaseID)
			if err != nil {
				return nil, err
			}
			userUsageList = append(userUsageList, UserUsage{
				// FirebaseID: firebaseID,
				UserName:       user.Name,
				UsageCount: usageCount,
			})
			totalUsage += usageCount // Increment total usage
		}
		// Sort the user usage list by usage count in descending order
		sort.Slice(userUsageList, func(i, j int) bool {
			return userUsageList[i].UsageCount > userUsageList[j].UsageCount
		})
		agentDetail := agentDetailsMap[agentID]
		response = append(response, AgentUsageResponse{
			AgentName:  agentDetail.Name,
			AgentID:    agentDetail.AgentID,
			ImageURL:   agentDetail.ImageURL,
			UserUsage:  userUsageList,
			TotalUsage: totalUsage, // Set total usage
		})
	}

	// Sort the response by total usage in descending order
	sort.Slice(response, func(i, j int) bool {
		return response[i].TotalUsage > response[j].TotalUsage
	})

	return response, nil
}