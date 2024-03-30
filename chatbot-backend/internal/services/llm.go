package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// LLMStruct defines the structure for a Large Language Model client
type LLM struct {
	ctx    context.Context
	client llms.Model
}

type Response struct {
	Result string `json:"result"`
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
func (lm *LLM) GenerateCompletion(prompt string) (string, error) {
	completion, err := llms.GenerateFromSinglePrompt(lm.ctx, lm.client, prompt)
	if err != nil {
		return "", err
	}
	fmt.Println(completion)
	return completion, nil
}

func (lm *LLM) GenerateAudio(prompt string) (Response, error) {
	url := "http://127.0.0.1:5000/generate_audio"

	// Define the JSON payload with the prompt string
	payload := map[string]interface{}{
		"description": prompt, 
	}

	fmt.Println(payload)

	// Marshal the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	fmt.Println(string(payloadBytes), err)

	if err != nil {
		fmt.Println("Error marshaling JSON payload:", err)
		return Response{}, err
	}

	// Send POST request with JSON payload
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return Response{}, err
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected response status code:", resp.StatusCode)
		return Response{}, errors.New(resp.Proto)
	}

	// Read response body
	var result Response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response body:", err)
		return Response{}, err
	}

	// Print results
	fmt.Println("Result:", result)
	return result, nil
}
