package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/redis/go-redis/v9"
	"github.com/alicebob/miniredis/v2"
)

type Server struct {
	port       int
	httpServer *http.Server
	RedisClient *redis.Client 
}

type Option func(*Server)

func NewServer(options ...Option) *Server {
	s := &Server{
		port: loadPortFromEnv(),
	}

	for _, option := range options {
		option(s)
	}

	s.initHTTPServer()
	return s
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func (s *Server) initHTTPServer() {
	s.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func (s *Server) Start() {

	log.Printf("Server starting on port %d\n", s.port)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}

}

func loadPortFromEnv() int {
	portStr, exists := os.LookupEnv("PORT")
	if !exists {
		return 8080
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Printf("Warning: Invalid PORT environment variable '%s', falling back to default port 8080.\n", portStr)
		return 8080
	}

	return port
}

func WithRedis(redisDB int) Option {
    return func(s *Server) {
        mr, err := miniredis.Run()
        if err != nil {
            log.Fatalf("Failed to start in-memory Redis server: %v", err)
        }
        s.RedisClient = redis.NewClient(&redis.Options{
            Addr: mr.Addr(),
            DB:   redisDB,  
        })
    }
}