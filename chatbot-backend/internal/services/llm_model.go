package services

import (
    "context"
	"fmt"
	"log"
    "github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// LLMStruct defines the structure for a Large Language Model client
type LLM struct {
	ctx context.Context
    client llms.Model
}

// NewLLM constructor function to create a new LLMStruct instance
func NewLLM() *LLM {
	ctx := context.Background()
	llm, err := openai.New()
	if err != nil {
	  log.Fatal(err)
	}
    return &LLM{ctx: ctx, client: llm}
}

// GenerateCompletion function to generate text completion using the LLMStruct
func (lm *LLM) GenerateCompletion( prompt string) (string, error) {
    completion, err := llms.GenerateFromSinglePrompt(lm.ctx, lm.client, prompt)
    if err != nil {
		log.Fatal(err)
        return "", err
    }
	fmt.Println(completion)
    return completion, nil
}
