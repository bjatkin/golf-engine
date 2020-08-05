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
var dir string

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

	// Shut down any currently running server
	err := server.Shutdown(context.Background())
	if err != nil {
		return err
	}

	return startServer()
}

func startServer() error {
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
