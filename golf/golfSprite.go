package golf

import (
	"math"
)

// LoadSprs loads the sprite sheet into memory
func (e *Engine) LoadSprs(sheet [0x3000]byte) {
	base := spriteBase
	for i, b := range sheet {
		e.RAM[i+base] = b
	}
}

// LoadFlags load the sprite flags into memory
func (e *Engine) LoadFlags(flags [0x200]byte) {
	base := spriteFlags
	for i, b := range flags {
		e.RAM[i+base] = b
	}
}

func (e *Engine) setActiveSpriteBuff(colAddr int) {
	c := toBytes(colAddr, 2, false)
	e.RAM[activeSpriteBuff] = c[0]
	e.RAM[activeSpriteBuff+1] = c[1]
}

// SOp additional options for drawing sprites
type SOp struct {
	FH, FV bool
	TCol   Col
	PFrom  []Col
	PTo    []Col
	W, H   int
	SW, SH float64
	Fixed  bool
}

// Spr draws 8x8 sprite n from the sprite sheet to the
// screen at x, y.
func (e *Engine) Spr(n int, x, y float64, opts ...SOp) {
	sx := n % 32
	sy := n / 32
	opt := SOp{}
	if len(opts) > 0 {
		opt = opts[0]
	}
	w := 1
	h := 1
	if opt.H > 1 {
		h = opt.H
	}
	if opt.W > 1 {
		w = opt.W
	}
	e.SSpr(sx*8, sy*8, w*8, h*8, x, y, opt)
}

// SSpr draw a rect from the sprite sheet to the screen
// sx, sy, sw, and sh define the rect on the sprite sheet
// dx, dy is the location to draw on the screen
func (e *Engine) SSpr(sx, sy, sw, sh int, dx, dy float64, opts ...SOp) {
	opt := SOp{}
	if len(opts) > 0 {
		opt = opts[0]
	}
	if !opt.Fixed {
		dx -= toFloat(e.RAM[cameraX:cameraX+2], true)
		dy -= toFloat(e.RAM[cameraY:cameraY+2], true)
	}
	if opt.SH == 0 {
		opt.SH = 1
	}
	if opt.SW == 0 {
		opt.SW = 1
	}
	buffBase := toInt(e.RAM[activeSpriteBuff:activeSpriteBuff+2], false)

	for x := 0; x < sw; x++ {
		for y := 0; y < sh; y++ {
			pxl := e.pget(float64(sx+x), float64(sy+y), buffBase, 256)
			if pxl != opt.TCol {
				pxl = subPixels(opt.PFrom, opt.PTo, pxl)
				fx := 0
				if opt.FH {
					fx = int(float64(sw) * opt.SW)
				}
				fy := 0
				if opt.FV {
					fy = int(float64(sh) * opt.SH)
				}
				for scaleX := int(float64(x) * opt.SW); scaleX < int(float64(x+1)*opt.SW); scaleX++ {
					for scaleY := int(float64(y) * opt.SH); scaleY < int(float64(y+1)*opt.SH); scaleY++ {
						e.Pset(dx+math.Abs(float64(fx-scaleX)), dy+math.Abs(float64(fy-scaleY)), pxl)
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
	return e.RAM[spriteFlags+n]&(0b10000000>>f) > 0
}

// FgetByte gets the byte flag on the nth sprite
func (e *Engine) FgetByte(n int) byte {
	return e.RAM[spriteFlags+n]
}

// Fset sets the fth flag on the nth sprite to s
func (e *Engine) Fset(n, f int, s bool) {
	e.RAM[spriteFlags+n] &= (0b10000001>>f ^ 0b11111111)
	if s {
		e.RAM[spriteFlags+n] |= (0b10000000 >> f)
	}
}

// FsetByte sets the byte flag on the nth sprite
func (e *Engine) FsetByte(n int, b byte) {
	e.RAM[spriteFlags+n] = b
}
