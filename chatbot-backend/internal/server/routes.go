package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server-template/internal/services"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var llm *services.LLM

type Response struct {
	Result string `json:"result"`
}

func init() {
	llm = services.NewLLM()
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8000"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	e.Static("/temp", "./.temp")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/api/chat", s.Chat)
	e.GET("/api/download_audio/:file_id", s.downloadAudio)
	e.GET("/health", s.HealthHandler)

	return e
}

type Message struct {
	Description string
}

func (s *Server) Chat(c echo.Context) error {
	var chat Message
	if err := json.NewDecoder(c.Request().Body).Decode(&chat); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Result: err.Error()})
	}

	jsonBytes, err := json.Marshal(chat.Description)
	fmt.Println(string(jsonBytes), err)
	res, err := llm.GenerateAudio(string(jsonBytes))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Result: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (s *Server) downloadAudio(c echo.Context) error {
	fileID, err := strconv.Atoi(c.Param("file_id"))
	println(fileID)
	if err != nil {
		return err
	}

	directory := ".temp"

	filename := strconv.Itoa(fileID) + ".wav"

	return c.File(directory + "/" + filename)
}

func (s *Server) HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, true)
}
