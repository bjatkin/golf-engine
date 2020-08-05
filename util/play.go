package main

import (
	"context"
	"net/http"
)

func startGameServer(args []string) error {
	gameDir := "./" + args[0]
	fs = http.FileServer(http.Dir(gameDir))

	// Shut down any currently running server
	err := server.Shutdown(context.Background())
	if err != nil {
		return err
	}
	return startServer()
}
