package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	openBrowser("http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func openBrowser(url string) {
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
	if err != nil {
		log.Fatal(err)
	}
}
