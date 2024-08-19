package generate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"gorm.io/gorm"
)

type Core struct {
	db *gorm.DB
}

func NewCore(db *gorm.DB) *Core {
	return &Core{db: db}
}

func (c *Core) Generate(ctx context.Context, prompt string, model string) (string,int,int, error) {
	float16Key := os.Getenv("FLOAT16_KEY")
	if float16Key == "" {
		return "",0,0, fmt.Errorf("FLOAT16_KEY environment variable not set")
	}

	url := "https://api.float16.cloud/v1/chat/completions"
	requestBody := map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful assistant."},
			{"role": "user", "content": prompt},
		},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return "",0,0, fmt.Errorf("error marshalling request body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "",0,0, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", float16Key))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "",0,0, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "",0,0, fmt.Errorf("error reading response body: %w", err)
	}

	var responseBody map[string]interface{}
	if err := json.Unmarshal(bodyResp, &responseBody); err != nil {
		return "",0,0, fmt.Errorf("error unmarshalling response body: %w", err)
	}
	var result string
	var completionTokens, promptTokens int

	if choices, exists := responseBody["choices"].([]interface{}); exists && len(choices) > 0 {
		if firstChoice, ok := choices[0].(map[string]interface{}); ok {
			if message, exists := firstChoice["message"].(map[string]interface{}); exists {
				if content, exists := message["content"].(string); exists {
					result = content
				}
			}
		}
	}

	if usage, exists := responseBody["usage"].(map[string]interface{}); exists {
		if completion, exists := usage["completion_tokens"].(float64); exists {
			completionTokens = int(completion)
		}
		if prompt, exists := usage["prompt_tokens"].(float64); exists {
			promptTokens = int(prompt)
		}
	}
	if result != "" && completionTokens != 0 && promptTokens != 0 {
		return result, completionTokens, promptTokens, nil
	}

	return "",0,0, fmt.Errorf("unexpected response format: %s", string(bodyResp))
}
