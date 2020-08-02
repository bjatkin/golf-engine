package main

import (
	"fmt"
	"image/png"
	"os"
)

type rgb struct {
	r, g, b uint8
}

func (rgb *rgb) JSString() string {
	return fmt.Sprintf("[%d, %d, %d]", rgb.r, rgb.g, rgb.b)
}

func (rgb *rgb) GoString() string {
	return fmt.Sprintf("{%d, %d, %d}", rgb.r, rgb.g, rgb.b)
}

type pal struct {
	p [4]rgb
}

func (p *pal) JSString() string {
	return fmt.Sprintf("[%s, %s, %s, %s]",
		p.p[0].JSString(),
		p.p[1].JSString(),
		p.p[2].JSString(),
		p.p[3].JSString())
}

func (p *pal) GoString() string {
	return fmt.Sprintf("pallet{%s, %s, %s, %s}",
		p.p[0].GoString(),
		p.p[1].GoString(),
		p.p[2].GoString(),
		p.p[3].GoString())
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	minX := img.Bounds().Min.X
	maxX := img.Bounds().Max.X
	minY := img.Bounds().Min.Y
	maxY := img.Bounds().Max.Y

	pal := [16]pal{}
	// add colors to the atlas
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			m := uint8(0b11111111)
			pal[y].p[x] = rgb{uint8(r) & m, uint8(g) & m, uint8(b) & m}
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	// JS Code
	js := "let pallets = ["
	for i, p := range pal {
		js += p.JSString()
		if i < 15 {
			js += ","
		}
	}
	js += "]"
	fmt.Println(js)

	// Go Code
	goCode := "var pallets = [16]pallet{"
	for i, p := range pal {
		goCode += p.GoString()
		if i < 15 {
			goCode += ","
		}
	}
	goCode += "}"
	fmt.Println(goCode)
}
