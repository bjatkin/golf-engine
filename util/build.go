package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func buildProject(args []string) error {
	config, err := ioutil.ReadFile(".golf_config")
	if err != nil {
		return errors.New(err.Error() + " perhaps try running init again")
	}

	confData := toGolfConfig(string(config))

	// pack in the sprite sheet
	spriteFileName := strings.Split(confData.spriteFile, ".")[0]
	err = convertSpriteSheet("assets/"+confData.spriteFile, spriteFileName+".go")
	if err != nil {
		return err
	}

	// pack in the map file
	mapFileName := strings.Split(confData.mapFile, ".")[0]
	err = convertMap("assets/"+confData.mapFile, "assets/"+confData.spriteFile, mapFileName+".go")

	// TODO: pack in the sprite flags file

	return runBuild()
}

func runBuild() error {
	err := exec.Command(appDir + "/build.sh").Run()
	if err != nil {
		return fmt.Errorf("%s, there may be a problem with your go code or you may need to run/re-run init", err)
	}

	return nil
}
