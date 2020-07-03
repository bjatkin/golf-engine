package main

import (
	"errors"
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
	err = convertSpriteSheet(confData.spriteFile, spriteFileName+".go")
	if err != nil {
		return err
	}

	// TODO: pack in the sprite sheet
	// TODO: pack in the sprite flags file

	return runBuild()
}

func runBuild() error {
	err := exec.Command(appDir + "/build.sh").Run()
	if err != nil {
		return errors.New(err.Error() + " perhaps try running init again")
	}

	return nil
}
