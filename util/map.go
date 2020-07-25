package main

import (
	"errors"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func convertMap(mapFile, spriteFile, outputFile string) error {
	sprites, err := os.Open(spriteFile)
	if err != nil {
		return err
	}

	sprimg, err := png.Decode(sprites)
	if err != nil {
		return err
	}

	sprAtlas, err := newColorAtlas(sprimg)
	if err != nil {
		return err
	}

	spriteKey := map[[64]byte]int{}
	spriteIndex := 0
	minX := sprimg.Bounds().Min.X
	maxX := sprimg.Bounds().Max.X
	minY := sprimg.Bounds().Min.Y
	maxY := sprimg.Bounds().Max.Y
	width := maxX - minX
	for cy := minY; cy < maxY; cy += 8 {
		for cx := minX; cx < maxX; cx += 8 {
			key := [64]byte{}
			i := 0
			for x := 0; x < 8; x++ {
				for y := 0; y < 8; y++ {
					key[i] = byte(sprAtlas.imgArray[cx+x+(cy+y)*width])
					i++
				}
			}
			spriteKey[key] = spriteIndex
			spriteIndex++
		}
	}

	mapdata, err := os.Open(mapFile)
	if err != nil {
		return err
	}

	mapimg, err := png.Decode(mapdata)
	if err != nil {
		return err
	}

	mapAtlas, err := newColorAtlas(mapimg)
	if err != nil {
		return err
	}

	if mapAtlas.pal1 != sprAtlas.pal1 || mapAtlas.pal2 != sprAtlas.pal2 {
		return fmt.Errorf("map and sprite sheet pallets do not match, map(%d, %d), spr(%d, %d)",
			mapAtlas.pal1, mapAtlas.pal2, sprAtlas.pal1, sprAtlas.pal2)
	}

	tiles := []int{}
	minX = mapimg.Bounds().Min.X
	maxX = mapimg.Bounds().Max.X
	minY = mapimg.Bounds().Min.Y
	maxY = mapimg.Bounds().Max.Y
	width = maxX - minX
	for cy := minY; cy < maxY; cy += 8 {
		for cx := minX; cx < maxX; cx += 8 {
			key := [64]byte{}
			i := 0
			for x := 0; x < 8; x++ {
				for y := 0; y < 8; y++ {
					key[i] = byte(mapAtlas.imgArray[cx+x+(cy+y)*width])
					i++
				}
			}
			tile, ok := spriteKey[key]
			if !ok {
				tile = 0
			}

			tiles = append(tiles, tile)
		}
	}

	low := []byte{}
	high := []byte{}
	for _, tile := range tiles {
		h := byte(0)
		if tile > 511 {
			return fmt.Errorf("tile with value %d found tile indexes above 512 are not supported in the map", tile)
		}
		if tile > 255 {
			h = 1
		}
		low = append(low, byte(tile))
		high = append(high, h)
	}

	return writeMapData(low, high, outputFile)
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

	for _, line := range csvFile {
		for _, tile := range line {
			id, err := strconv.Atoi(tile)
			if err != nil {
				return err
			}
			h := byte(0)
			if id > 511 {
				return errors.New("tile indexes above 512 are not supported in the map")
			}
			if id > 255 {
				h = 1
			}
			low = append(low, byte(id))
			high = append(high, h)
		}
	}
	return writeMapData(low, high, outputFile)
}

func writeMapData(low, high []byte, outputFile string) error {
	conv := []byte{}
	for i := 0; i < len(high); i++ {
		if i%8 == 0 && i != 0 {
			top := i
			mashedHigh, err := packHighBytes(high[top-8 : top])
			if err != nil {
				return err
			}
			conv = append(conv, mashedHigh)
		}
		conv = append(conv, low[i])
	}

	content := "package main\n\nvar mapData = [0x4800]byte{\n"
	for _, b := range conv {
		content += fmt.Sprintf("0x%X", b) + ","
	}
	content += "\n}"
	return ioutil.WriteFile(outputFile, []byte(content), 0666)
}

func packHighBytes(bytes []byte) (byte, error) {
	if len(bytes) != 8 {
		return 0, errors.New("bytes array must be exactly 8 bytes")
	}
	ret := byte(0)
	for i := 0; i < 8; i++ {
		ret |= bytes[7-i] << i
	}
	return ret, nil
}
