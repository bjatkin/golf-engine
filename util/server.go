package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"sync"
)

func startServer(args []string) error {
	serverSG.Add(1)

	go func() {
		defer serverSG.Done()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf(err.Error())
		}
	}()

	fmt.Printf("   Server Started!")
	return openBrowser("http://localhost:8080/")
}

func stopServer(args []string) error {
	err := server.Shutdown(context.Background())
	server = &http.Server{Addr: ":8080"}
	serverSG = &sync.WaitGroup{}
	fmt.Printf("   Server Stopped!")
	return err
}

func openBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}
