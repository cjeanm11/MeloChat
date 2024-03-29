package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "server-template/internal/database"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	// CORS middleware configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8000"}, // Allow all origins (for demo purposes)
		AllowMethods:     []string{"POST", "OPTIONS"},
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
	resp := map[string]string{"message": "Hello World"}
	return c.JSON(http.StatusOK, resp)
}

// HealthHandler handles requests to the "/health" endpoint
func (s *Server) HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
