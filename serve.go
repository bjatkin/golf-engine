package main

// Serve static files from current working directory
// it will start at port 8080 if port is being used it will try next one

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

//Add a url to save byte data to the correct location so we can save resource files
// func main() {
// 	port := 8080
// 	for {
// 		addr := fmt.Sprintf(":%d", port)
// 		listener, err := net.Listen("tcp", addr)
// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, "err opening port", err)
// 			port++
// 			continue
// 		}
// 		fmt.Printf("Listening at %s\n", addr)
// 		openbrowser("http://localhost:8080/")
// 		log.Fatal(http.Serve(listener, logger(http.FileServer(http.Dir(".")))))
// 	}
// }

func main() {
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/saveresource", saveResource)
	openBrowser("http://localhost:8080/")
	log.Fatal(http.ListenAndServe(addr, nil))
}

// func logger(next http.Handler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(r.Method, r.URL.Path)
// 		next.ServeHTTP(w, r)
// 	}
// }

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

func saveResource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We're HERE in save resources!")
}
