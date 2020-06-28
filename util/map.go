package main

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

func convertMap(inputFile, spriteFile, outputFile string) error {
	return errors.New("This function is not implemented yet")
}

func convertCSVMap(inputFile, outputFile string) error {
	low := []byte{}
	high := []byte{}

	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	csvFile := strings.Split(string(file), ",")
	if len(csvFile) > 512 {
		return errors.New("only the first 512 sprites can have sprite tiles")
	}

	for _, tile := range csvFile {
		id, err := strconv.Atoi(tile)
		if err != nil {
			return err
		}
		h := byte(0)
		if id > 512 {
			return errors.New("tile indexes above 512 are not supported in the map")
		}
		if id > 255 {
			h = 1
		}
		low = append(low, byte(id))
		high = append(high, h)
	}

	return nil
}
