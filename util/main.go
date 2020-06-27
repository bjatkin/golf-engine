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

type pallet [4]color

var pallets = []pallet{
	pallet{{0, 0, 0},
		{224, 60, 40},
		{255, 255, 255},
		{215, 215, 215}},
	pallet{{168, 168, 168},
		{123, 123, 123},
		{52, 52, 52},
		{21, 21, 21}},
	pallet{{13, 32, 48},
		{65, 93, 102},
		{113, 166, 161},
		{189, 255, 202}},
	pallet{{37, 226, 205},
		{10, 152, 172},
		{0, 82, 128},
		{0, 96, 75}},
	pallet{{32, 181, 98},
		{88, 211, 50},
		{19, 157, 8},
		{0, 78, 0}},
	pallet{{23, 40, 8},
		{55, 109, 3},
		{106, 180, 23},
		{140, 214, 18}},
	pallet{{190, 235, 113},
		{238, 255, 169},
		{182, 193, 33},
		{147, 151, 23}},
	pallet{{204, 143, 21},
		{255, 187, 49},
		{255, 231, 55},
		{246, 143, 55}},
	pallet{{173, 78, 26},
		{35, 23, 18},
		{92, 60, 13},
		{174, 108, 55}},
	pallet{{197, 151, 130},
		{226, 215, 181},
		{79, 21, 7},
		{130, 60, 61}},
	pallet{{218, 101, 94},
		{225, 130, 137},
		{245, 183, 132},
		{255, 233, 197}},
	pallet{{255, 130, 206},
		{207, 60, 113},
		{135, 22, 70},
		{163, 40, 179}},
	pallet{{204, 105, 228},
		{213, 156, 252},
		{254, 201, 237},
		{226, 201, 255}},
	pallet{{166, 117, 254},
		{106, 49, 202},
		{90, 25, 145},
		{33, 22, 64}},
	pallet{{61, 52, 165},
		{98, 100, 220},
		{155, 160, 239},
		{152, 220, 255}},
	pallet{{91, 168, 255},
		{10, 137, 255},
		{2, 74, 202},
		{0, 23, 125}},
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
		if p.pal2 < p.pal1 {
			p.pal1, p.pal2 = p.pal2, p.pal1
		}
		fmt.Printf("Pal1: %v Pal2: %v\n", p.pal1, p.pal2)
	}
	if p.pal1 == -1 {
		p.pal1 = pxl.pal
		fmt.Printf("Pal1: %v\n", p.pal1)
	}
	if pxl.pal != p.pal1 && pxl.pal != p.pal2 {
		return fmt.Errorf("more than 2 pallets detected, Pal1: %d, Pal2: %d, NewPal: %d, NewCol: %v", p.pal1, p.pal2, pxl.pal, pxl.col)
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
	fmt.Printf("\nminX: %d minY: %d maxX: %d maxY: %d\n", minX, minY, maxX, maxY)
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			pxl := nearistPixel(int(r)/256, int(g)/256, int(b)/256)
			err := image.addPixel(pxl)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				fmt.Printf("R: %d, G: %d, B: %d\n", int(r)/256, int(g)/256, int(b)/256)
				fmt.Printf("X: %d, Y: %d", x, y)
				return
			}
			if x == 20 && y == 2 {
				fmt.Printf("Len: %d, Col %d\n", len(image.pixels), image.pixels[len(image.pixels)-1])
			}
			// if x == 20 && y == 1 {
			// 	// fmt.Printf("rgb: (%d, %d, %d)\n", int(r)/256, int(g)/256, int(b)/256)
			// 	fmt.Printf("Pal 5 = %d, Col 3 = %d\n", pxl.pal, pxl.col)
			// 	minDist := 99999999999.0
			// 	bestPixel := pixel{}
			// 	for p, pal := range pallets {
			// 		for c, col := range pal {
			// 			dist := dist(int(r)/256, int(g)/256, int(b)/256, int(col.r), int(col.g), int(col.b))
			// 			if dist < minDist {
			// 				fmt.Printf("Dist: %.2f Col: %d, Pal: %d\n", dist, c, p)
			// 				fmt.Printf("rgb: (%d, %d, %d) == (%d, %d, %d)\n", int(r)/256, int(g)/256, int(b)/256, col.r, col.g, col.b)
			// 				bestPixel.pal = p
			// 				bestPixel.col = c
			// 				minDist = dist
			// 			}
			// 		}
			// 	}
			// }
		}
	}
	fmt.Printf("Image Size 256x128: %d\n", len(image.pixels))

	colorBuff := []byte{}
	for i, pxl := range image.pixels {
		shift := (i % 4) * 2
		index := i / 4
		if shift == 0 {
			colorBuff = append(colorBuff, 0)
		}

		col := pxl.col
		// if i == 530 {
		// 	fmt.Printf("COl %v \n", pxl.col)
		// 	fmt.Printf("Before: colorBuff: %b, Col: %b, Shift: %d\n", colorBuff[index], col, shift)
		// }
		colorBuff[index] |= (byte(col) << shift)
		// if i == 530 {
		// 	fmt.Printf("After: colorBuff: %b, Col: %b, Shift: %d\n", colorBuff[index], col, shift)
		// }
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
		if i == 533 {
			fmt.Printf("pal %v \n", pxl.pal)
			fmt.Printf("Before: palBuff: %b, pal: %b, Shift: %d\n", palBuff[index], p, shift)
		}
		palBuff[index] |= p << shift
		if i == 533 {
			fmt.Printf("After: palBuff: %b, pal: %b, Shift: %d\n", palBuff[index], p, shift)
		}
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

	fmt.Printf("colBuff: %d, palBuff: %d,total bytes: %d\n", len(colorBuff), len(palBuff), len(bytes))
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
