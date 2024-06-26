package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8000", "http://localhost"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	e.Static("/temp", "./.temp")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/api/chat", s.chatHandler)
	e.GET("/api/download_audio/:file_id", s.downloadAudioHandler)
	e.GET("/health", s.healthHandler)

	return e
}

