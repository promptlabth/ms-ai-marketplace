package generateservice

// import (
// 	"context"
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	openai "github.com/sashabaranov/go-openai"
// 	"golang.org/x/oauth2/google"
// 	"google.golang.org/api/option"
// 	vertexai "google.golang.org/api/vertexai/v1"
// )

// type GenerateService struct {
// 	openaiClient     *openai.Client
// 	vertexAIService  *vertexai.Service
// 	anthropicAPIKey  string
// }

// func NewGenerateService() *GenerateService {
// 	// Initialize OpenAI client
// 	openaiClient := openai.NewClient(os.Getenv("OPENAI_KEY"))

// 	// Initialize Vertex AI client
// 	ctx := context.Background()
// 	cred, err := google.CredentialsFromJSON(ctx, []byte(os.Getenv("GCP_SA_KEY_JSON")), vertexai.VertexaiScope)
// 	if err != nil {
// 		panic(err)
// 	}
// 	vertexAIService, err := vertexai.NewService(ctx, option.WithCredentials(cred))
// 	if err != nil {
// 		panic(err)
// 	}

// 	return &GenerateService{
// 		openaiClient:    openaiClient,
// 		vertexAIService: vertexAIService,
// 		anthropicAPIKey: os.Getenv("ANTHROPIC_API_KEY"),
// 	}
// }

// func (gs *GenerateService) GenerateMessageOpenAI(inputPrompt string) (string, error) {
// 	req := openai.ChatCompletionRequest{
// 		Model: "gpt-3.5-turbo",
// 		Messages: []openai.ChatCompletionMessage{
// 			{Role: "user", Content: inputPrompt},
// 		},
// 	}
// 	resp, err := gs.openaiClient.CreateChatCompletion(context.Background(), req)
// 	if err != nil {
// 		return "", err
// 	}
// 	return resp.Choices[0].Message.Content, nil
// }

// func (gs *GenerateService) GetModelAndParameter(featureName string) (string, map[string]interface{}) {
// 	model := "gemini-1.0-pro-002"
// 	modelList := map[string]map[string]interface{}{
// 		"เขียนแคปชั่นขายของ": {
// 			"model": model,
// 			"parametor": map[string]interface{}{
// 				"max_output_tokens": 4096,
// 				"temperature":       0.6,
// 				"top_p":             0.8,
// 				"top_k":             40,
// 			},
// 		},
// 		"ช่วยคิดคอนเทนต์": {
// 			"model": model,
// 			"parametor": map[string]interface{}{
// 				"max_output_tokens": 8192,
// 				"temperature":       0.2,
// 				"top_p":             0.8,
// 				"top_k":             40,
// 			},
// 		},
// 		"เขียนบทความ": {
// 			"model": model,
// 			"parametor": map[string]interface{}{
// 				"max_output_tokens": 8192,
// 				"temperature":       0.2,
// 				"top_p":             0.8,
// 				"top_k":             40,
// 			},
// 		},
// 		"เขียนสคริปวิดีโอสั้น": {
// 			"model": model,
// 			"parametor": map[string]interface{}{
// 				"max_output_tokens": 8192,
// 				"temperature":       0.4,
// 				"top_p":             0.8,
// 				"top_k":             40,
// 			},
// 		},
// 		"เขียนประโยคเปิดคลิป": {
// 			"model": model,
// 			"parametor": map[string]interface{}{
// 				"max_output_tokens": 1024,
// 				"temperature":       0.5,
// 				"top_p":             0.8,
// 				"top_k":             40,
// 			},
// 		},
// 	}
// 	return modelList[featureName]["model"].(string), modelList[featureName]["parametor"].(map[string]interface{})
// }

// func (gs *GenerateService) GenerateMessageVertexAI(inputPrompt, featureName string) (string, error) {
// 	model, params := gs.GetModelAndParameter(featureName)
// 	generationConfig := map[string]interface{}{
// 		"max_output_tokens": 8192,
// 		"temperature":       1,
// 		"top_p":             0.95,
// 	}

// 	requestBody, err := json.Marshal(map[string]interface{}{
// 		"prompt":         inputPrompt,
// 		"generationConfig": generationConfig,
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	req, err := http.NewRequest("POST", "https://vertex.googleapis.com/v1/projects/"+os.Getenv("GCP_PROJECT_ID")+"/locations/asia-southeast1/models/"+model+":predict", ioutil.NopCloser(bytes.NewReader(requestBody)))
// 	if err != nil {
// 		return "", err
// 	}
// 	req.Header.Set("Authorization", "Bearer "+cred.Token)
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	var result map[string]interface{}
// 	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 		return "", err
// 	}

// 	return result["predictions"].([]interface{})[0].(map[string]interface{})["content"].(string), nil
// }

