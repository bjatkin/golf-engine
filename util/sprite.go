package main

import (
	"errors"
	"fmt"
	"image"
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
	colors     [8]rgb
	cindex     int
	palMap     [8]int
	colMap     [8]int
	pal1, pal2 int
	imgArray   []int
}

func newColorAtlas(img image.Image) (colorAtlas, error) {
	minX := img.Bounds().Min.X
	maxX := img.Bounds().Max.X
	minY := img.Bounds().Min.Y
	maxY := img.Bounds().Max.Y

	ret := colorAtlas{}
	// add colors to the atlas
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			err := ret.addColor(rgb{r, g, b})
			if err != nil {
				return ret, err
			}
		}
	}
	ret.sortColor()

	// build the image array
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			i, ok := ret.cIndex(rgb{r, g, b})
			if !ok {
				return ret, errors.New("unknown color error")
			}
			ret.imgArray = append(ret.imgArray, i)
		}
	}

	// Find the best pallets for this sprite sheet
	minDist := ^uint32(0)
	var palMap, colMap [8]int
	var found bool
	for i := 0; i < 16; i++ {
		for j := i; j < 16; j++ {
			errDist := uint32(0)
			for atlasI, c := range ret.colors {
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
			if errDist <= minDist {
				ret.pal1, ret.pal2 = i, j
				ret.palMap = palMap
				ret.colMap = colMap
				minDist = errDist
			}
			if errDist == 0 && i == j {
				ret.pal2 = 0
				found = true
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}
	return ret, nil
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

func (a *colorAtlas) cIndex(color rgb) (int, bool) {
	for i := 0; i < 8; i++ {
		if a.colors[i].Equal(color) {
			return i, true
		}
	}
	return 0, false
}

func (a *colorAtlas) sortColor() {
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 8; j++ {
			o, t := a.colors[i], a.colors[j]
			if o.r|(o.g<<8)|(o.b<<16) < t.r|(t.g<<8)|(t.b<<16) {
				a.colors[i], a.colors[j] = a.colors[j], a.colors[i]
			}
		}
	}
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

	atlas, err := newColorAtlas(img)
	if err != nil {
		return err
	}

	colArray, palArray := []byte{}, []byte{}
	for _, i := range atlas.imgArray {
		colArray = append(colArray, byte(atlas.colMap[i]))
		palArray = append(palArray, byte(atlas.palMap[i]))
	}

	// pack the color information. 4 pixels to a byte
	colBuff := []byte{}
	for i, col := range colArray {
		shift := (i % 4) * 2
		index := i / 4
		if shift == 0 {
			colBuff = append(colBuff, 0)
		}

		col := col
		colBuff[index] |= (col << shift)
	}

	// pack the pallet informaiotn. 8 pixels to a byte
	palBuff := []byte{}
	for i, pal := range palArray {
		shift := (i % 8)
		index := i / 8
		if shift == 0 {
			palBuff = append(palBuff, 0)
		}

		p := byte(0)
		if pal == byte(atlas.pal2) {
			p = byte(1)
		}
		palBuff[index] |= p << shift
	}

	// interlace and write the pallet and color data to a string
	bytes := []string{}
	j := 0
	for i := 0; i < len(colBuff); i++ {
		bytes = append(bytes, fmt.Sprintf("0x%X", colBuff[i]))
		if (i+1)%2 == 0 {
			bytes = append(bytes, fmt.Sprintf("0x%X", palBuff[j]))
			j++
		}
	}

	// write the final string to a package file
	content := fmt.Sprintf("package main\n\nvar spriteSheet = [0x%x]byte {\n", len(bytes))
	content += strings.Join(bytes, ",")
	content += ",\n}"

	return ioutil.WriteFile(outputFile, []byte(content), 0644)
}
