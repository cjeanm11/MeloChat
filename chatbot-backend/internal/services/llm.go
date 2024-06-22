package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

)

// LLMStruct defines the structure for a Large Language Model client
type LLM struct {
	ctx    context.Context
	// client llms.Model
}

type Response struct {
	Result string `json:"result"`
}

// NewLLM constructor function to create a new LLMStruct instance
func NewLLM() *LLM {
	ctx := context.Background()
	return &LLM{ctx: ctx}
}


func (lm *LLM) GenerateAudio(prompt string) (Response, error) {
	url := "http://127.0.0.1:5000/generate_audio"

	payload := map[string]interface{}{
		"description": prompt, 
	}

	fmt.Println(payload)

	payloadBytes, err := json.Marshal(payload)
	fmt.Println(string(payloadBytes), err)

	if err != nil {
		fmt.Println("Error marshaling JSON payload:", err)
		return Response{}, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected response status code:", resp.StatusCode)
		return Response{}, errors.New(resp.Proto)
	}

	var result Response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response body:", err)
		return Response{}, err
	}

	fmt.Println("Result:", result)
	return result, nil
}
