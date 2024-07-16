package generateservice

import (
	"context"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)


type Generate struct {
	
}

func NewGenerate() *Generate {
	return &Generate{}
}

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