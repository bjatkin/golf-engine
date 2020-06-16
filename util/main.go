package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

type color struct {
	r, g, b uint32
}

var pallet = []color{
	color{0, 0, 0},       //black
	color{84, 84, 84},    // dark gray
	color{168, 168, 168}, // light gray
	color{255, 255, 255}, // white
}

func main() {
	filePath := os.Args[1]
	fmt.Printf("File: %s\n", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error 1! %s", err.Error())
		return
	}

	img, err := png.Decode(file)
	if err != nil {
		fmt.Printf("Error 2! %s", err.Error())
		return
	}

	minX := img.Bounds().Min.X
	maxX := img.Bounds().Max.X
	minY := img.Bounds().Min.Y
	maxY := img.Bounds().Max.Y
	fmt.Printf("MinX: %d, MaxX: %d, MinY: %d, MaxY: %d\n", minX, maxX, minY, maxY)
	imgColor := []int{}
	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			if a == 0 {
				continue
			}
			avg := (r + g + b) / 3
			bestI := 0
			bestDist := 65536.0
			for i := 0; i < 4; i++ {
				dist := math.Abs(float64(avg - pallet[i].r))
				if dist < bestDist {
					bestI = i
					bestDist = dist
				}
			}
			imgColor = append(imgColor, bestI)
		}
	}

	colorBuff := []byte{}
	for i, col := range imgColor {
		shift := (i % 4) * 2
		index := i / 4
		if shift == 0 {
			colorBuff = append(colorBuff, 0)
		}

		colorBuff[index] = (byte(col) << shift)
	}
	fmt.Printf("%#v\n", colorBuff[:30])

	palBuff := []byte{}
	for i := 0; i < len(colorBuff)/2; i++ {
		palBuff = append(palBuff, 0)
	}
	fmt.Printf("Image: %d, ColorBuffer: %d, PalBuffer: %d\n", len(imgColor), len(colorBuff), len(palBuff))

	outputFile := os.Args[2]
	bytes := []string{}
	for _, b := range colorBuff {
		bStr := fmt.Sprintf("%b", b)
		for i := len(bStr); i < 8; i++ {
			bStr = "0" + bStr
		}
		bStr = "0b" + bStr
		bytes = append(bytes, bStr)
	}

	for _, b := range palBuff {
		bStr := fmt.Sprintf("%b", b)
		for i := len(bStr); i < 8; i++ {
			bStr = "0" + bStr
		}
		bStr = "0b" + bStr
		bytes = append(bytes, bStr)
	}

	content := fmt.Sprintf("[%d]byte {", len(bytes))
	content += strings.Join(bytes, ",")
	content += "}"

	err = ioutil.WriteFile(outputFile, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error 3: %s\n", err.Error())
		return
	}
}
