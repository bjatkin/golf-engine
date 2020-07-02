package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

type pixelArray struct {
	pixels     []pixel
	pal1, pal2 int
}

func newPixelArray() pixelArray {
	return pixelArray{
		pal1: -1,
		pal2: -1,
	}
}

func (p *pixelArray) addPixel(pxl pixel) error {
	if p.pal2 == -1 && p.pal1 != -1 && p.pal1 != pxl.pal {
		p.pal2 = pxl.pal
		if p.pal2 < p.pal1 {
			p.pal1, p.pal2 = p.pal2, p.pal1
		}
	}
	if p.pal1 == -1 {
		p.pal1 = pxl.pal
	}
	if pxl.pal != p.pal1 && pxl.pal != p.pal2 {
		return fmt.Errorf("more than 2 pallets detected, Pal1: %d, Pal2: %d, NewPal: %d, NewCol: %v",
			p.pal1, p.pal2, pxl.pal, pxl.col)
	}
	p.pixels = append(p.pixels, pxl)
	return nil
}

func dist(r1, g1, b1, r2, g2, b2 int) float64 {
	a := float64(r2 - r1)
	b := float64(g2 - g1)
	c := float64(b2 - b1)
	return math.Sqrt(a*a + b*b + c*c)
}

var test = 0

func nearistPixel(r, g, b int) pixel {
	minDist := 65535.0
	bestPixel := pixel{}
	for p, pal := range pallets {
		for c, col := range pal {
			dist := dist(r, g, b, int(col.r), int(col.g), int(col.b))
			if dist < minDist {
				bestPixel.pal = p
				bestPixel.col = c
				minDist = dist
			}
		}
	}
	return bestPixel
}

func convertSpriteSheet(inputFile, outputFile string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}

	img, err := png.Decode(file)
	if err != nil {
		return err
	}

	minX := img.Bounds().Min.X
	maxX := img.Bounds().Max.X
	minY := img.Bounds().Min.Y
	maxY := img.Bounds().Max.Y
	image := newPixelArray()
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			pxl := nearistPixel(int(r)/256, int(g)/256, int(b)/256)
			err := image.addPixel(pxl)
			if err != nil {
				return err
			}
		}
	}

	colorBuff := []byte{}
	for i, pxl := range image.pixels {
		shift := (i % 4) * 2
		index := i / 4
		if shift == 0 {
			colorBuff = append(colorBuff, 0)
		}

		col := pxl.col
		colorBuff[index] |= (byte(col) << shift)
	}

	palBuff := []byte{}
	for i, pxl := range image.pixels {
		shift := (i % 8)
		index := i / 8
		if shift == 0 {
			palBuff = append(palBuff, 0)
		}

		p := byte(0)
		if pxl.pal == image.pal2 {
			p = byte(1)
		}
		palBuff[index] |= p << shift
	}

	bytes := []string{}

	j := 0
	for i := 0; i < len(colorBuff); i++ {
		bytes = append(bytes, printByte(colorBuff[i]))
		if (i+1)%2 == 0 {
			bytes = append(bytes, printByte(palBuff[j]))
			j++
		}
	}

	content := fmt.Sprintf("package main\n\nvar spriteSheet = [0x%X]byte {\n", len(bytes))
	content += strings.Join(bytes, ",")
	content += ",\n}"

	err = ioutil.WriteFile(outputFile, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}
