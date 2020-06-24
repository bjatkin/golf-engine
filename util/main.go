package main

import (
	"errors"
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

type pallet [4]color

var pallets = []pallet{
	pallet{
		color{0, 0, 0},       // black
		color{84, 84, 84},    // dark gray
		color{168, 168, 168}, // light gray
		color{255, 255, 255}, // white
	},
	pallet{
		color{54, 36, 26},    // Brown
		color{119, 97, 65},   // Cream
		color{181, 165, 125}, // Light Cream
		color{218, 207, 190}, // Pale
	},
	pallet{
		color{21, 25, 18},  // GB 0
		color{41, 50, 36},  // GB 1
		color{62, 75, 54},  // GB 2
		color{82, 100, 71}, // GB 3
	},
	pallet{
		color{103, 125, 89},  // GB 4
		color{124, 149, 107}, // GB 5
		color{145, 175, 125}, // GB 6
		color{165, 199, 142}, // GB 7
	},
}

type pixel struct {
	pal, col int
}

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
	}
	if p.pal1 == -1 {
		p.pal1 = pxl.pal
	}
	if pxl.pal != p.pal1 && pxl.pal != p.pal2 {
		return errors.New("More than 2 pallets detected")
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

func main() {
	filePath := os.Args[1]
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
	image := newPixelArray()
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			pxl := nearistPixel(int(r)/256, int(g)/256, int(b)/256)
			image.addPixel(pxl)
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

	outputFile := os.Args[2]
	bytes := []string{}

	j := 0
	for i := 0; i < len(colorBuff); i++ {
		bytes = append(bytes, printByte(colorBuff[i]))
		if (i+1)%2 == 0 {
			bytes = append(bytes, printByte(palBuff[j]))
			j++
		}
	}

	content := fmt.Sprintf("package main\n\nvar spriteSheet = [%d]byte {\n", len(bytes))
	content += strings.Join(bytes, ",")
	content += ",\n}"

	err = ioutil.WriteFile(outputFile, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error 3: %s\n", err.Error())
		return
	}
}

func printByte(b byte) string {
	ret := fmt.Sprintf("%b", b)
	for i := len(ret); i < 8; i++ {
		ret = "0" + ret
	}
	return "0b" + ret
}
