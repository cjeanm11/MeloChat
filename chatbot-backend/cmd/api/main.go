package main

import (
	"server-template/internal/server"
)

func main() {
	srv := server.NewServer(
		server.WithRedis(0),
	)
	srv.Start()
}
