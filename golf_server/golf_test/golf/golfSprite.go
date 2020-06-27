package golf

import (
	"fmt"
	"math"
)

// LoadSprs loads the sprite sheet into memory
func (e *Engine) LoadSprs(sheet [0x3000]byte) {
	base := spriteColBase
	for i, b := range sheet {
		e.RAM[i+base] = b
	}
}

func (e *Engine) setActiveSpriteBuff(colAddr, palAddr int) {
	c, p := toBytes(colAddr, 2, false), toBytes(palAddr, 2, false)
	e.RAM[activeSpriteColBuff] = c[0]
	e.RAM[activeSpriteColBuff+1] = c[1]
	e.RAM[activeSpritePalBuff] = p[0]
	e.RAM[activeSpritePalBuff+1] = p[1]
}

// SprOpts additional options for drawing sprites
type SprOpts struct {
	FlipH, FlipV   bool
	Transparent    Col
	PalFrom        []Col
	PalTo          []Col
	Width, Height  int
	ScaleW, ScaleH float64
	Fixed          bool
}

// Spr draws 8x8 sprite n from the sprite sheet to the
// screen at x, y.
func (e *Engine) Spr(n, x, y int, opts ...SprOpts) {
	sx := n % 32
	sy := n / 32
	opt := SprOpts{}
	if len(opts) > 0 {
		opt = opts[0]
	}
	w := 1
	h := 1
	if opt.Height > 1 {
		h = opt.Height
	}
	if opt.Width > 1 {
		w = opt.Width
	}
	e.SSpr(sx*8, sy*8, w*8, h*8, x, y, opt)
}

// SSpr draw a rect from the sprite sheet to the screen
// sx, sy, sw, and sh define the rect on the sprite sheet
// dx, dy is the location to draw on the screen
func (e *Engine) SSpr(sx, sy, sw, sh, dx, dy int, opts ...SprOpts) {
	opt := SprOpts{}
	if len(opts) > 0 {
		opt = opts[0]
	}
	if !opt.Fixed {
		dx -= toInt(e.RAM[cameraX:cameraX+2], true)
		dy -= toInt(e.RAM[cameraY:cameraY+2], true)
	}
	if opt.ScaleH == 0 {
		opt.ScaleH = 1
	}
	if opt.ScaleW == 0 {
		opt.ScaleW = 1
	}

	buffBase := toInt(e.RAM[activeSpriteColBuff:activeSpriteColBuff+2], false)

	for x := 0; x < sw; x++ {
		for y := 0; y < sh; y++ {
			pxl := e.pget(sx+x, sy+y, buffBase, 256)
			if sx+x == 0 && sy+y == 0 {
				fmt.Printf("pget: %b\n", pxl)
			}
			if pxl != opt.Transparent {
				pxl = subPixels(opt.PalFrom, opt.PalTo, pxl)
				fx := 0
				if opt.FlipH {
					fx = int(float64(sw) * opt.ScaleW)
				}
				fy := 0
				if opt.FlipV {
					fy = int(float64(sh) * opt.ScaleH)
				}
				for scaleX := int(float64(x) * opt.ScaleW); scaleX < int(float64(x+1)*opt.ScaleW); scaleX++ {
					for scaleY := int(float64(y) * opt.ScaleH); scaleY < int(float64(y+1)*opt.ScaleH); scaleY++ {
						e.Pset(dx+int(math.Abs(float64(fx-scaleX))), dy+int(math.Abs(float64(fy-scaleY))), pxl)
					}
				}
			}
		}
	}
}

// subPixels is used to swap pixels based on a pallet swap
func subPixels(palFrom, palTo []Col, col Col) Col {
	if len(palFrom) == 0 {
		return col
	}
	for i, p := range palFrom {
		if p == col {
			return palTo[i]
		}
	}
	return col
}

// Fget gets the fth flag on the nth sprite
func (e *Engine) Fget(n, f int) bool {
	return e.RAM[spriteFlags+n]&(0b00000001<<f) > 0
}

// FgetByte gets the byte flag on the nth sprite
func (e *Engine) FgetByte(n int) byte {
	return e.RAM[spriteFlags+n]
}

// Fset sets the fth flag on the nth sprite to s
func (e *Engine) Fset(n, f int, s bool) {
	e.RAM[spriteFlags+n] &= (0b00000001<<f ^ 0b11111111)
	if s {
		e.RAM[spriteFlags+n] |= (0b00000001 << f)
	}
}

// FsetByte sets the byte flag on the nth sprite
func (e *Engine) FsetByte(n int, b byte) {
	e.RAM[spriteFlags+n] = b
}
