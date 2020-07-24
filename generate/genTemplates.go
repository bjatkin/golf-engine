package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Error: got %d args but need 2 args <output> <template dir>", len(os.Args))
		return
	}

	pkg := os.Args[3]
	content := "package " + pkg + "\n//generated code do not edit\n"
	outputFile := os.Args[1]
	dir := os.Args[2]

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		data, err := ioutil.ReadFile(dir + "/" + file.Name())
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}
		filename := strings.Split(file.Name(), ".")[0] + "Template"
		l, i := len(data), 6+len(filename)+len(content)
		content += fmt.Sprintf("var %s=%#v\n", filename, data)
		content = content[:i] + fmt.Sprintf("0x%x", l) + content[i:]
	}

	ioutil.WriteFile(outputFile, []byte(content), 0666)
}
