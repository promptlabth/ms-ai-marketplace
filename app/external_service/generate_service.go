package generateservice

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	// "golang.org/x/oauth2"
	"github.com/promptlabth/ms-ai-marketplace/config"
	"golang.org/x/oauth2/google"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type GenerateService struct {
	ProjectID    string
	LocationID   string
	ModelID      string
	Client       *storage.Client
	AnthropicKey string
}

func NewGenerateService() (*GenerateService, error) {

	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(config.Val.GCP.GoogleAppleciationCredential)))
	if err != nil {
		return nil, fmt.Errorf("failed to create storage client: %w", err)
	}

	// Initialize Vertex AI client
	projectID := config.Val.GCP.ProjectId
	if projectID == "" {
		log.Fatalf("GCP_PROJECT_ID environment variable not set")
		return nil, errors.New("GCP_PROJECT_ID environment variable not set")
	}

	locationID := "asia-southeast1"
	modelID := "gemini-1.0-pro-002"

	// Placeholder for initializing Vertex AI, as Go SDK for Vertex AI might differ
	log.Printf("Initialized Vertex AI client for project %s in location %s\n", projectID, locationID)

	// Set up the Anthropic client with the API key from the environment variable
	apiKey := config.Val.NLPApiKey.Anthropic
	if apiKey == "" {
		log.Fatalf("ANTHROPIC_API_KEY environment variable not set")
		return nil, errors.New("ANTHROPIC_API_KEY environment variable not set")
	}

	return &GenerateService{
		ProjectID:    projectID,
		LocationID:   locationID,
		ModelID:      modelID,
		Client:       client,
		AnthropicKey: apiKey,
	}, nil
}

func (s *GenerateService) getModelAndParameter(featureName string) (map[string]interface{}, error) {
	model := "gemini-1.0-pro-002"
	modelList := map[string]map[string]interface{}{
		"APE": {
			"model": model,
			"parameter": map[string]interface{}{
				"max_output_tokens": 4096,
				"temperature":       0.6,
				"top_p":             0.8,
				"top_k":             40,
			},
		},
		"RICEE": {
			"model": model,
			"parameter": map[string]interface{}{
				"max_output_tokens": 8192,
				"temperature":       0.2,
				"top_p":             0.8,
				"top_k":             40,
			},
		},
		"TAG": {
			"model": model,
			"parameter": map[string]interface{}{
				"max_output_tokens": 8192,
				"temperature":       0.2,
				"top_p":             0.8,
				"top_k":             40,
			},
		},
		"ERP": {
			"model": model,
			"parameter": map[string]interface{}{
				"max_output_tokens": 8192,
				"temperature":       0.4,
				"top_p":             0.8,
				"top_k":             40,
			},
		},
		"RPPPP": {
			"model": model,
			"parameter": map[string]interface{}{
				"max_output_tokens": 1024,
				"temperature":       0.5,
				"top_p":             0.8,
				"top_k":             40,
			},
		},
	}

	if params, ok := modelList[featureName]; ok {
		return params, nil
	}
	return nil, fmt.Errorf("feature name not found: %s", featureName)
}

func (s *GenerateService) GenerateMessageVertexAI(inputPrompt, featureName string) (string, error) {
	ctx := context.Background()

	// Get vertex parameter
	modelParams, err := s.getModelAndParameter(featureName)
	if err != nil {
		return "", err
	}

	modelName := modelParams["model"].(string)
	generationConfig := modelParams["parameter"].(map[string]interface{})

	// Prepare request payload
	requestPayload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"role": "user",
				"parts": []interface{}{
					inputPrompt,
				},
			},
		},
		"generationConfig": generationConfig,
		"safetySettings": []map[string]interface{}{
			{"category": "HARM_CATEGORY_HATE_SPEECH", "threshold": "BLOCK_MEDIUM_AND_ABOVE"},
			{"category": "HARM_CATEGORY_DANGEROUS_CONTENT", "threshold": "BLOCK_MEDIUM_AND_ABOVE"},
			{"category": "HARM_CATEGORY_SEXUALLY_EXPLICIT", "threshold": "BLOCK_MEDIUM_AND_ABOVE"},
			{"category": "HARM_CATEGORY_HARASSMENT", "threshold": "BLOCK_MEDIUM_AND_ABOVE"},
		},
	}

	requestBody, err := json.Marshal(requestPayload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request payload: %v", err)
	}

	// Send request
	apiEndpoint := "us-central1-aiplatform.googleapis.com"
	url := fmt.Sprintf("https://%s/v1/projects/%s/locations/%s/publishers/google/models/%s:streamGenerateContent",
		apiEndpoint, s.ProjectID, s.LocationID, modelName)

	// Use the client to send the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %v", err)
	}

	tokenSource, err := google.DefaultTokenSource(ctx, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return "", fmt.Errorf("failed to create token source: %v", err)
	}

	token, err := tokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve token: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(body), nil
}

func (g *GenerateService) GenerateMessageOpenAI(inputPrompt string) (string, error) {
	apiKey := os.Getenv("OPENAI_KEY")
	if apiKey == "" {
		log.Fatal("environment variable not set")
	}

	url := "https://api.openai.com/v1/completions"

	requestBody, err := json.Marshal(map[string]interface{}{
		"model":             "gpt-3.5-turbo-instruct",
		"prompt":            inputPrompt,
		"temperature":       1,
		"max_tokens":        1024,
		"top_p":             1,
		"frequency_penalty": 0,
		"presence_penalty":  0,
	})
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", errors.New("ChatCompletion error" + err.Error())
	}
	var responseBody map[string]interface{}
	if err := json.Unmarshal(body, &responseBody); err == nil {
		if errorDetails, exists := responseBody["error"]; exists {
			if errorMap, ok := errorDetails.(map[string]interface{}); ok {
				if errorMessage, ok := errorMap["message"].(string); ok {
					return "", errors.New(errorMessage)
				}
			}
		}
	}
	return string(body), nil
}
