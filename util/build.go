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
	fmt.Println("   reading golf config file")

	confData := toGolfConfig(string(config))

	// pack in the sprite sheet
	fmt.Println("   converting sprite sheet")
	err = convertSpriteSheet(confData.spriteFile, confData.outputSpriteFile)
	if err != nil {
		return fmt.Errorf("spritesheet err %s", err.Error())
	}

	// pack in the map file
	fmt.Println("   converting the map file")
	mapFileType := strings.Split(confData.mapFile, ".")[1]
	if mapFileType == "png" {
		err = convertMap(confData.mapFile, confData.spriteFile, confData.outputMapFile)
	} else {
		err = convertCSVMap(confData.mapFile, confData.outputMapFile)
	}
	if err != nil {
		return fmt.Errorf("mapfile err %s", err.Error())
	}

	// TODO: pack in the sprite flags file
	fmt.Println("   converting the sprite flags file")
	err = convertFlag(confData.flagFile, confData.outputFlagFile)
	if err != nil {
		return fmt.Errorf("sprite flags err %s", err.Error())
	}

	fmt.Println("   building the project")
	return runBuild()
}

func runBuild() error {
	out, err := exec.Command(appDir + "/build.sh").CombinedOutput()
	if err != nil {
		fmt.Println("\n" + string(out))
		return fmt.Errorf("%s, there may be a problem with your go code or your build.sh file, fix your code or try running/re-running init", err)
	}

	fmt.Println("   building was successful")
	return nil
}
