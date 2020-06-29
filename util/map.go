package main

import (
	"errors"
	"fmt"
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

	lines := strings.Split(string(file), "\n")
	csvFile := [][]string{}
	for i, line := range lines {
		cols := strings.Split(string(line), ",")
		if len(cols) > 128 {
			return fmt.Errorf("row %d, has %d columns, no more than 128 columns permitted", i, len(lines))
		}
		for len(cols) < 128 {
			cols = append(cols, "0")
		}
		csvFile = append(csvFile, cols)
	}

	count, linesct := 0, 0
	for _, line := range csvFile {
		linesct++
		for _, tile := range line {
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
			count++
		}
	}
	fmt.Printf("\n\nCount: %d(%d) Lines: %d\n\n", count, len(high), linesct)

	conv := []byte{}
	for i := 0; i < count; i++ {
		if (i+1)%8 == 0 {
			top := i + 1
			mashedHigh, err := packHighBytes(high[top-8 : top])
			if err != nil {
				return err
			}
			conv = append(conv, mashedHigh)
		}
		conv = append(conv, low[i])
	}

	//TODO: this should print the file out
	return nil
}

func packHighBytes(bytes []byte) (byte, error) {
	if len(bytes) != 8 {
		return 0, errors.New("bytes array must be exactly 8 bytes")
	}
	ret := byte(0)
	for i := 0; i < 8; i++ {
		ret |= bytes[i] << i
	}
	return ret, nil
}
