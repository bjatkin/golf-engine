package main

import (
	"fmt"
	"net/http"
)

func startGameServer(args []string) error {
	gameDir := "./" + args[0]
	fs = http.FileServer(http.Dir(gameDir))

	if serverRunning {
		fmt.Println("   swaping to serve " + args[0] + " directory")
		return nil
	}

	err := startServer()
	if err == nil {
		fmt.Println("   now serving the " + args[0] + " directory")
	}
	serverRunning = true

	return err
}
