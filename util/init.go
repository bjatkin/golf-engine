package main

import "errors"

const (
	pngFile = iota
	textFile
	goFile
	jsFile
	configFile
	htmlFile
)

func addFiles() error {
	var err error
	err = addFileIfNew("spritesheet.png", []byte("test"), pngFile)
	if err != nil {
		return err
	}
	err = addFileIfNew("map.png", []byte("test"), pngFile)
	if err != nil {
		return err
	}
	err = addFileIfNew("main.go", []byte("test"), goFile)
	if err != nil {
		return err
	}
	err = addFileIfNew("index.html", []byte("test"), htmlFile)
	if err != nil {
		return err
	}

	return nil
}

func addFileIfNew(fileName string, content []byte, fileType int) error {
	return errors.New("Not yet implemented")
}
