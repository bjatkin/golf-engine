package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

var cmdDone = false
var lastCmd = "help"
var serverSG = &sync.WaitGroup{}
var server = &http.Server{Addr: ":8080"}
var appDir = ""

//go:generate ../generate/genTemplates packedTemplates.go templates main
func main() {
	var err error
	appDir, err = os.Getwd()
	if err != nil {
		fmt.Printf("Could not get the working directory %s\n", err.Error())
		return
	}

	// Add help and bangbang command to prevent an initializaiton loop
	commands = append(commands, helpCommand)
	commands = append(commands, bangbangCommand)

	// Make the server serve the file server
	fs := http.FileServer(http.Dir("./web"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			fmt.Println("build")
			buildProject(nil)
			printCommandLine()
		}
		fs.ServeHTTP(w, r)
	})

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nWelcome to the Golf Toolkit!\ntype help for more info\n")
	for !cmdDone {
		printCommandLine()
		cmd, _ := reader.ReadString('\n')
		cmd = strings.Trim(cmd, " \t\n")
		runCmd(cmd)
		if cmd != bangbangCommand.command {
			lastCmd = cmd
		}
	}
}

func runCmd(cmd string) {
	cmdRun := false
	for _, command := range commands {
		args := strings.Split(cmd, " ")
		if args[0] == command.command {
			cmdRun = true
			if len(args) != command.argCount+1 {
				printErrorLine("Incorrect arg count, usage: " + command.usage)
				break
			}
			err := command.commandHandler(args[1:])
			if err != nil {
				printErrorLine(err.Error())
				break
			}
			break
		}
	}
	if !cmdRun {
		printErrorLine("Unrecognized command " + cmd + ", type help for more info")
	}
}

func printCommandLine() {
	fmt.Print("\n >")
}

func printErrorLine(err string) {
	fmt.Println("   Error: " + err)
}

func printBlockText(text string, maxLen int) {
	ln := 0
	p := "   "
	for i := 0; i < len(text); i++ {
		a := text[i]
		p += string(a)
		if a == '.' || a == ',' || a == ' ' {
			if i/maxLen > ln {
				ln = i / maxLen
				p += "\n   "
			}
		}
	}
	fmt.Print(p)
}

func printByte(b byte) string {
	ret := fmt.Sprintf("%b", b)
	for i := len(ret); i < 8; i++ {
		ret = "0" + ret
	}
	return "0b" + ret
}
