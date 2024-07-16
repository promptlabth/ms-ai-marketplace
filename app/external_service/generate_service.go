package generateservice

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	openai "github.com/sashabaranov/go-openai"
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

func (g *Generate) GenerateMessageOpenAI(inputPrompt string) (string, error) { 
	var openai_api_key = os.Getenv("OPENAI_KEY")
	client := openai.NewClient(openai_api_key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: inputPrompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", errors.New("ChatCompletion error" + err.Error())
	}

	return resp.Choices[0].Message.Content,nil
}