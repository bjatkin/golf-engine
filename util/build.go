package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func buildProject(args []string) error {
	config, err := ioutil.ReadFile("golf_config")
	if err != nil {
		return errors.New(err.Error() + " perhaps golf_config is missing or corupted try running init again")
	}

	confData := toGolfConfig(string(config))

	// pack in the sprite sheet
	err = convertSpriteSheet(confData.spriteFile, confData.outputSpriteFile)
	if err != nil {
		return err
	}
	// fmt.Printf("   Converting " + confData.spriteFile + " \n")

	// pack in the map file
	mapFileType := strings.Split(confData.mapFile, ".")[1]
	if mapFileType == "png" {
		err = convertMap(confData.mapFile, confData.spriteFile, confData.outputMapFile)
	} else {
		err = convertCSVMap(confData.mapFile, confData.outputMapFile)
	}
	// fmt.Printf("   Converting " + confData.mapFile + " \n")

	// TODO: pack in the sprite flags file

	return runBuild()
}

func runBuild() error {
	err := exec.Command(appDir + "/build.sh").Run()
	if err != nil {
		return fmt.Errorf("%s, there may be a problem with your go code or your build.sh file, fix your code or try running/re-running init", err)
	}

	return nil
}
