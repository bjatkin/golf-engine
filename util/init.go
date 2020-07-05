package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

type golfConfig struct {
	name             string
	spriteFile       string
	mapFile          string
	flagFile         string
	outputSpriteFile string
	outputMapFile    string
	outputFlagFile   string
}

func (g golfConfig) String() string {
	return "name=" + g.name + "\n" +
		"spriteFile=" + g.spriteFile + "\n" +
		"mapFile=" + g.mapFile + "\n" +
		"flagFile=" + g.flagFile + "\n" +
		"outputSpriteFile=" + g.outputSpriteFile + "\n" +
		"outputMapFile=" + g.outputMapFile + "\n" +
		"outputFlagFile=" + g.outputFlagFile
}

func toGolfConfig(data string) golfConfig {
	pairs := strings.Split(data, "\n")
	ret := golfConfig{}
	for _, pair := range pairs {
		if pair == "" {
			continue
		}
		kv := strings.Split(pair, "=")
		k, v := kv[0], kv[1]

		switch k {
		case "name":
			ret.name = v
		case "spriteFile":
			ret.spriteFile = v
		case "mapFile":
			ret.mapFile = v
		case "flagFile":
			ret.flagFile = v
		case "outputSpriteFile":
			ret.outputSpriteFile = v
		case "outputMapFile":
			ret.outputMapFile = v
		case "outputFlagFile":
			ret.outputFlagFile = v
		}
	}
	return ret
}

func initProject(args []string) error {
	// do this first to make sure we can get the wasmjs file
	// before we start adding other files
	wasm, err := getWASMjs()
	if err != nil {
		return err
	}

	buildTemplate = []byte(strings.Replace(string(buildTemplate), "main.wasm", args[0]+".wasm", -1))
	err = addExecFile("build.sh", buildTemplate, true)
	if err != nil {
		return err
	}

	err = createNewDir("assets")
	if err != nil {
		return err
	}

	err = addFile("assets/spritesheet.png", spritesheetTemplate, false)
	if err != nil {
		return err
	}

	err = addFile("assets/map.png", mapTemplate, false)
	if err != nil {
		return err
	}

	err = addFile("main.go", mainTemplate, false)
	if err != nil {
		return err
	}

	err = createNewDir("web")
	if err != nil {
		return err
	}

	indexTemplate = []byte(strings.Replace(string(indexTemplate), "main.wasm", args[0]+".wasm", -1))
	err = addFile("web/index.html", indexTemplate, false)
	if err != nil {
		return err
	}

	err = addFile("web/draw.js", drawTemplate, false)
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
		flagFile:         "assets/flag.csv",
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
