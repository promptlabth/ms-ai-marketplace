package generateservice

import (
	"context"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// import (
// 	"context"
// 	// "encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"

// 	openai "github.com/sashabaranov/go-openai"
// 	"cloud.google.com/go/compute/metadata"
// 	"google.golang.org/api/option"
// 	"google.golang.org/api/storage/v1"
// )

type Generate struct {
	
}

func NewGenerate() *Generate {
	return &Generate{}
}

// type VertexModel struct {
// 	ModelName string
// }

// type ModelParameter struct {
// 	MaxOutputTokens int     `json:"max_output_tokens"`
// 	Temperature     float64 `json:"temperature"`
// 	TopP            float64 `json:"top_p"`
// 	TopK            int     `json:"top_k"`
// }

// func main() {
// 	app := App{}
// 	app.init()
// 	inputPrompt := "Your prompt here"
// 	featureName := "เขียนบทความ"
// 	response := app.generateMessageOpenAI(inputPrompt)
// 	fmt.Println("OpenAI Response: ", response)

// 	vertexResponse := app.generateMessageVertexAI(inputPrompt, featureName)
// 	fmt.Println("Vertex AI Response: ", vertexResponse)
// }

// func (a *App) init() {
// 	apiKey := os.Getenv("OPENAI_KEY")
// 	gcpProjectID := os.Getenv("GCP_PROJECT_ID")

// 	// Initialize OpenAI client
// 	a.anthropicClient = openai.Client{
// 		AuthToken: apiKey,
// 	}

// 	// Initialize Google Cloud credentials
// 	cred, err := ioutil.ReadFile("gcp_sa_key.json")
// 	if err != nil {
// 		log.Fatalf("Failed to read GCP service account key file: %v", err)
// 	}

// 	ctx := context.Background()
// 	storageService, err := storage.NewService(ctx, option.WithCredentialsJSON(cred))
// 	if err != nil {
// 		log.Fatalf("Failed to create storage service: %v", err)
// 	}

// 	// Use storageService as needed...
// 	_ = storageService

// 	fmt.Println("Initialization complete")
// }

// func (a *App) generateMessageOpenAI(inputPrompt string) string {
// 	resp, err := a.anthropicClient.Completions.Create(context.TODO(), openai.CompletionRequest{
// 		Model: openai.GPT3_5_Turbo,
// 		Messages: []openai.ChatCompletionMessage{
// 			{Role: "user", Content: inputPrompt},
// 		},
// 	})

// 	if err != nil {
// 		log.Fatalf("OpenAI API call failed: %v", err)
// 	}

// 	return resp.Choices[0].Message.Content
// }

// func (a *App) getModelAndParameter(featureName string) (string, ModelParameter) {
// 	model := "gemini-1.0-pro-002"
// 	modelList := map[string]ModelParameter{
// 		"เขียนแคปชั่นขายของ":  {4096, 0.6, 0.8, 40},
// 		"ช่วยคิดคอนเทนต์":   {8192, 0.2, 0.8, 40},
// 		"เขียนบทความ":       {8192, 0.2, 0.8, 40},
// 		"เขียนสคริปวิดีโอสั้น": {8192, 0.4, 0.8, 40},
// 		"เขียนประโยคเปิดคลิป":  {1024, 0.5, 0.8, 40},
// 	}

// 	return model, modelList[featureName]
// }

// func (a *App) getVertexModel(modelName string) VertexModel {
// 	return VertexModel{ModelName: modelName}
// }

// func (a *App) generateMessageVertexAI(inputPrompt, featureName string) string {
// 	model, params := a.getModelAndParameter(featureName)

// 	generationConfig := ModelParameter{
// 		MaxOutputTokens: 8192,
// 		Temperature:     1,
// 		TopP:            0.95,
// 	}

// 	vertexModel := a.getVertexModel(model)

// 	// Assuming vertexModel.GenerateContent is the method to generate content
// 	// Replace with actual method to generate content using the vertexModel

// 	// For demonstration, we're returning a placeholder string
// 	return fmt.Sprintf("Generated content for prompt: %s with model: %s and parameters: %+v", inputPrompt, vertexModel.ModelName, generationConfig)
// }

func (g *Generate) GenerateMessageGemini(inputPrompt string) (string, error) { // Exported method
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(inputPrompt))
	if err != nil {
		return "", err
	}
	return resp.PromptFeedback.BlockReason.String(), nil
}