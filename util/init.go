package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

func initProject(args []string) error {
	// do this first to make sure we can get the wasmjs file
	// before we start adding other files
	wasm, err := getWASMjs()
	if err != nil {
		return err
	}

	newBuildTemplate := []byte(strings.Replace(string(buildTemplate[:]), "main.wasm", args[0]+".wasm", -1))
	err = addExecFile("build.sh", newBuildTemplate, true)
	if err != nil {
		return err
	}

	err = createNewDir("assets")
	if err != nil {
		return err
	}

	err = addFile("assets/spritesheet.png", spritesheetTemplate[:], false)
	if err != nil {
		return err
	}

	err = addFile("assets/map.png", mapTemplate[:], false)
	if err != nil {
		return err
	}

	err = addFile("assets/spriteflags.csv", spriteflagsTemplate[:], false)
	if err != nil {
		return err
	}

	err = addFile("main.go", mainTemplate[:], false)
	if err != nil {
		return err
	}

	err = createNewDir("web")
	if err != nil {
		return err
	}

	newIndexTemplate := []byte(strings.Replace(string(indexTemplate[:]), "main.wasm", args[0]+".wasm", -1))
	err = addFile("web/index.html", newIndexTemplate, false)
	if err != nil {
		return err
	}

	err = addFile("web/wasm_exec.js", wasm, true)
	if err != nil {
		return err
	}

	config := golfConfig{
		name:             args[0],
		spriteFile:       "assets/spritesheet.png",
		mapFile:          "assets/map.png",
		flagFile:         "assets/spriteflags.csv",
		outputSpriteFile: "spritesheet.go",
		outputMapFile:    "map.go",
		outputFlagFile:   "flag.go",
	}

	err = addFile("golf_config", []byte(config.String()), true)
	if err != nil {
		return err
	}

	fmt.Printf("   Success!\n")
	return nil
}

func addFile(fileName string, content []byte, overwrite bool) error {
	if !overwrite {
		if _, err := os.Stat(fileName); err == nil {
			return nil
		} else if !os.IsNotExist(err) {
			return err
		}
	}
	fmt.Printf("   Adding %s\n", fileName)
	return ioutil.WriteFile(fileName, content, 0666)
}

func addExecFile(fileName string, content []byte, overwrite bool) error {
	if !overwrite {
		if _, err := os.Stat(fileName); err == nil {
			return nil
		} else if !os.IsNotExist(err) {
			return err
		}
	}
	fmt.Printf("   Adding %s\n", fileName)
	return ioutil.WriteFile(fileName, content, 0777)
}

func getWASMjs() ([]byte, error) {
	// Grab the wasm_exec file from the go/misc/wasm folder
	dir := "/usr/local/go"
	if runtime.GOOS == "windows" {
		dir = "C:\\Go"
	}
	if os.Getenv("GOROOT") != "" {
		dir = os.Getenv("GOROOT")
	}

	if runtime.GOOS == "windows" {
		dir += "\\misc\\wasm\\wasm_exec.js"
	} else {
		dir += "/misc/wasm/wasm_exec.js"
	}

	return ioutil.ReadFile(dir)
}

func createNewDir(dir string) error {
	if _, err := os.Stat(dir); err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}
	return os.MkdirAll(dir, 0777)
}
