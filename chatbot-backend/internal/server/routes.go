package server

import (
	"encoding/json"
	"net/http"
	"server-template/internal/model"
	"server-template/internal/services"

	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "server-template/internal/database"
)

var llm *services.LLM

// Response struct
type Response struct {
	Result string `json:"result"`
}

func init() {
	llm = services.NewLLM()
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	// CORS middleware configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8000"}, // Allow all origins (for demo purposes)
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/chat", s.Chat)
	e.GET("/health", s.HealthHandler)

	return e
}

// HelloWorldHandler handles requests to the root "/" endpoint
func (s *Server) Chat(c echo.Context) error {
	var chat models.Conversation
	if err := json.NewDecoder(c.Request().Body).Decode(&chat); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error decoding request body: %v", err)
	}
	jsonBytes, err := json.Marshal(chat)
	fmt.Println(string(jsonBytes), err)

	res, err := llm.GenerateCompletion(string(jsonBytes))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error generating response: %v", err)
	}

	return c.JSON(http.StatusOK, Response{Result: res})
}

// HealthHandler handles requests to the "/health" endpoint
func (s *Server) HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
