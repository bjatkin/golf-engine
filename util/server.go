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

var serverSG = &sync.WaitGroup{}
var server = &http.Server{Addr: ":8080"}
var fs http.Handler
var serverRunning bool

func initServer() {
	// Make the server serve the file server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			fmt.Println("build")
			err := buildProject(nil)
			if err != nil {
				printErrorLine(err.Error())
			}
			printCommandLine()
		}
		fs.ServeHTTP(w, r)
	})
}

func startDevServer(args []string) error {
	fs = http.FileServer(http.Dir("./web"))

	if serverRunning {
		fmt.Println("   swaping to server the web directory")
		return nil
	}
	err := startServer()
	if err == nil {
		fmt.Println("   now serving the web directory")
	}
	serverRunning = true

	return err
}

func startServer() error {
	serverSG.Add(1)

	go func() {
		defer serverSG.Done()
		fmt.Println("Starting the Server listenting and serving")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf(err.Error())
		}
	}()

	return openBrowser("http://localhost:8080/")
}

func stopServer(args []string) error {
	err := server.Shutdown(context.Background())
	server = &http.Server{Addr: ":8080"}
	serverSG = &sync.WaitGroup{}
	fmt.Printf("   server stopped")
	serverRunning = false

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
