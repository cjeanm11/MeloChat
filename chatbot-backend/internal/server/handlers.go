package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server-template/internal/services"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

const MAX_TOKEN_PROMPT = 2000
var llm *services.LLM

func init() {
	llm = services.NewLLM()
}

type Response struct {
	Result string `json:"result"`
}

type Message struct {
	Description string
}

func (s *Server) chatHandler(c echo.Context) error {
    var chat Message
    if err := json.NewDecoder(c.Request().Body).Decode(&chat); err != nil {
        log.Printf("Error decoding request body: %v", err)
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
    }

    message, err := json.Marshal(chat.Description)
    if err != nil {
        log.Printf("Error marshaling message: %v", err)
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to process message")
    }

    ctx := c.Request().Context()

    cacheKey := fmt.Sprintf("chat:%s", string(message))
    cachedResponse, err := s.RedisClient.Get(ctx, cacheKey).Result()

    if err == nil { 
        log.Println("Cache hit for message:", string(message))
        var res Response
        if err := json.Unmarshal([]byte(cachedResponse), &res); err != nil {
            log.Printf("Error unmarshalling cached response: %v", err)
        } else {
            return c.JSON(http.StatusOK, res)
        }
    } else { 
        log.Printf("Error fetching from Redis: %v", err)
    }

    tokenCount := llm.EstimateTokenCount(string(message))
    if tokenCount > MAX_TOKEN_PROMPT {
        return echo.NewHTTPError(http.StatusBadRequest, "Message exceeds token limit")
    }

    log.Println("Generating audio for message:", string(message))
    res, err := llm.GenerateAudio(string(message))
    if err != nil {
        log.Printf("Error generating audio: %v", err)
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate audio")
    }

    responseBytes, err := json.Marshal(res)
    if err != nil {
        log.Printf("Error marshaling response: %v", err)
    } else {
        s.RedisClient.Set(ctx, cacheKey, string(responseBytes), 24*time.Hour)
    }

    return c.JSON(http.StatusOK, res)
}



func (s *Server) downloadAudioHandler(c echo.Context) error {
	fileID, err := strconv.Atoi(c.Param("file_id"))
	println(fileID)
	if err != nil {
		return err
	}
	directory := ".temp"
	filename := strconv.Itoa(fileID) + ".wav"
	return c.File(directory + "/" + filename)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, true)
}
