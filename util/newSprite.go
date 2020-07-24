package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

type rgb struct {
	r, g, b uint32
}

func (a *rgb) Equal(b rgb) bool {
	return a.r == b.r && a.g == b.g && a.b == b.b
}

func (a *rgb) Dist(b rgb) uint32 {
	m := uint32(255)
	x := (a.r&m - b.r&m)
	y := (a.g&m - b.g&m)
	z := (a.b&m - b.b&m)
	return x*x + y*y + z*z
}

func (a *rgb) PalDist(pal int) (uint32, int) {
	minDist := ^uint32(0)
	index := 0
	for i, c := range pallets[pal] {
		d := a.Dist(c)
		if d < minDist {
			minDist = d
			index = i
		}
	}

	return minDist, index
}

func (a rgb) String() string {
	return fmt.Sprintf("(%d, %d, %d)", uint8(a.r), uint8(a.g), uint8(a.b))
}

type colorAtlas struct {
	colors [8]rgb
	cindex int
}

func (a *colorAtlas) addColor(color rgb) error {
	for i := 0; i < a.cindex; i++ {
		if color.Equal(a.colors[i]) {
			return nil
		}
	}

	if a.cindex > 7 {
		return fmt.Errorf("more than 8 colors detected %s", color)
	}
	a.colors[a.cindex] = color
	a.cindex++

	return nil
}

type pixelArray struct {
	pallets    []int
	colors     []int
	pal1, pal2 int
}

func (p *pixelArray) addPixel(pal, col int) error {
	if pal != p.pal1 && pal != p.pal2 {
		return fmt.Errorf("can't add a pixel with pal %d, only pal %d and %d are allowed on this sheet", pal, p.pal1, p.pal2)
	}
	p.pallets = append(p.pallets, pal)
	p.colors = append(p.colors, col)
	return nil
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

	// Build the color atlas
	atlas := colorAtlas{}
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			err := atlas.addColor(rgb{r, g, b})
			if err != nil {
				return err
			}
		}
	}

	// Find the best pallets for this sprite sheet
	minDist := ^uint32(0)
	var pal1, pal2 int
	var palMap, colMap [8]int
	var bestPalMap, bestColMap [8]int
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			errDist := uint32(0)
			for atlasI, c := range atlas.colors {
				a, ai := c.PalDist(i)
				b, bi := c.PalDist(j)
				palMap[atlasI] = i
				colMap[atlasI] = ai
				if b < a {
					a = b
					palMap[atlasI] = j
					colMap[atlasI] = bi
				}
				errDist += a
			}
			if errDist < minDist {
				pal1, pal2 = i, j
				bestPalMap = palMap
				bestColMap = colMap
			}
		}
	}

	// build the pixel array using the best pallets
	sheet := pixelArray{pal1: pal1, pal2: pal2}
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			col := rgb{r, g, b}
			aIndex := 0

			add := false
			//TODO iteragte only through the list of colors
			//but not beyond (e.g. if you have only 4 colors rather than 8)
			for i, c := range atlas.colors {
				if c.Equal(col) {
					err := sheet.addPixel(bestColMap[i], bestPalMap[i])
					if err != nil {
						return nil
					}
					add := true
					break
				}
			}
			if !add {
				return fmt.Errorf("there was an error adding the pixel at %d, %d from the image", x, y)
			}
		}
	}

	// pack the color information. 4 pixels to a byte
	colorBuff := []byte{}
	for i, col := range sheet.colors {
		shift := (i % 4) * 2
		index := i / 4
		if shift == 0 {
			colorBuff = append(colorBuff, 0)
		}

		col := col
		colorBuff[index] |= (byte(col) << shift)
	}

	// pack the pallet informaiotn. 8 pixels to a byte
	palBuff := []byte{}
	for i, pal := range sheet.pallets {
		shift := (i % 8)
		index := i / 8
		if shift == 0 {
			palBuff = append(palBuff, 0)
		}

		p := byte(0)
		if pal == sheet.pal2 {
			p = byte(1)
		}
		palBuff[index] |= p << shift
	}

	// interlace and write the pallet and color data to a string
	bytes := []string{}
	j := 0
	for i := 0; i < len(colorBuff); i++ {
		bytes = append(bytes, fmt.Sprintf("0x%X", colorBuff[i]))
		if (i+1)%2 == 0 {
			bytes = append(bytes, fmt.Sprintf("0x%X", palBuff[j]))
			j++
		}
	}

	// write the final string to a package file
	content := fmt.Sprintf("package main\n\nvar spritesheet = [0x%x]byte {\n", len(bytes))
	content += strings.Join(bytes, ",")
	content += ",\n}"

	return ioutil.WriteFile(outputFile, []byte(content), 0644)
}
