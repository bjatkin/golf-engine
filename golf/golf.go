package golf

import (
	"syscall/js"
)

// Engine Screen Width and Height
const (
	ScreenHeight = 192
	ScreenWidth  = 192
)

// Engine is the golf engine
type Engine struct {
	RAM           *[0xFFFF]byte
	screenBufHook js.Value
	Draw          func()
	Update        func()
}

//go:generate ../generate/genTemplates packedTemplates.go templates golf

// NewEngine creates a new golf engine
func NewEngine(update func(), draw func()) *Engine {
	ret := Engine{
		RAM:    &[0xFFFF]byte{},
		Draw:   draw,
		Update: update,
	}

	doc := js.Global().Get("document")
	ret.initKeyListener(doc)
	ret.initMouseListener(js.Global().Get("golfcanvas"))

	ret.RClip() // Reset the cliping box

	// Set internal resources
	base := internalSpriteBase
	for i := 0; i < 0x0900; i++ {
		ret.RAM[i+base] = internalSpriteSheet[i]
	}

	ret.RAM[startAnim] = 255
	ret.PalA(0)
	ret.PalB(1)

	// Inject the nessisary JS
	script := doc.Call("createElement", "script")
	script.Set("innerHTML", string(drawTemplate[:]))
	doc.Get("body").Call("appendChild", script)

	// Hook into the injected js
	ret.screenBufHook = js.Global().Get("screenBuff")

	return &ret
}

// Run starts the game engine running
func (e *Engine) Run() {
	var renderFrame js.Func

	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e.addFrame()

		if e.Frames() < int(e.RAM[startAnim]) {
			e.startupAnim()
		} else {
			e.Update()
			e.Draw()

			e.drawMouse()

			e.tickKeyboard()
			e.tickMouse()
		}

		js.CopyBytesToJS(e.screenBufHook, e.RAM[:screenPalSet+1])
		js.Global().Call("drawScreen")
		js.Global().Call("requestAnimationFrame", renderFrame)

		return nil
	})

	done := make(chan struct{}, 0)

	js.Global().Call("requestAnimationFrame", renderFrame)
	<-done
}

// Frames is the number of frames since the engine was started
func (e *Engine) Frames() int {
	return toInt(e.RAM[frames:frames+3], false)
}

func (e *Engine) addFrame() {
	f := toInt(e.RAM[frames:frames+3], false)
	f++
	b := toBytes(f, 3, false)
	e.RAM[frames] = b[0]
	e.RAM[frames+1] = b[1]
	e.RAM[frames+2] = b[2]
}

// DrawMouse sets the draw style
// 0 = none
// 1 = mouse
// 2 = hand
// 3 = cross
func (e *Engine) DrawMouse(style int) {
	e.RAM[mouseBase] &= 0b00111111
	e.RAM[mouseBase] |= byte(style << 6)
}

// drawMouse draws the mouse on the screen
func (e *Engine) drawMouse() {
	e.setActiveSpriteBuff(internalSpriteBase)

	cursor := e.RAM[mouseBase] >> 6
	opt := SOp{Fixed: true, TCol: Col7}

	if cursor == 1 {
		e.Spr(18, float64(e.RAM[mouseX]), float64(e.RAM[mouseY]), opt)
	}
	if cursor == 2 {
		e.Spr(50, float64(e.RAM[mouseX]), float64(e.RAM[mouseY]), opt)
	}
	if cursor == 3 {
		e.Spr(82, float64(e.RAM[mouseX]), float64(e.RAM[mouseY]), opt)
	}

	e.setActiveSpriteBuff(spriteBase)
}

// Mouse returns the X, Y coords of the mouse
func (e *Engine) Mouse() (int, int) {
	return int(e.RAM[mouseX]), int(e.RAM[mouseY])
}

// Cls fills the screen with col and resets TextL and TextR
func (e *Engine) Cls(col Col) {
	textLline, textRline = 0, 0

	c := col & 0b00000011
	colBG := byte((c << 6) | (c << 4) | (c << 2) | c)
	palBG := byte(0)
	if col>>2 == 0b00100001 {
		palBG = 0b11111111
	}

	for i := 0; i < screenPalSet; i++ {
		e.RAM[i] = colBG
		if (i+1)%3 == 0 {
			e.RAM[i] = palBG
		}
	}
}

// Camera moves the camera which modifies all draw functions
func (e *Engine) Camera(x, y int) {
	xb := toBytes(x, 2, true)
	yb := toBytes(y, 2, true)
	e.RAM[cameraX] = xb[0]
	e.RAM[cameraX+1] = xb[1]
	e.RAM[cameraY] = yb[0]
	e.RAM[cameraY+1] = yb[1]
}

// Rect draws a rectangle border on the screen
func (e *Engine) Rect(x, y, w, h float64, col Col, fixed ...bool) {
	f := false
	if len(fixed) > 0 {
		f = fixed[0]
	}
	if !f {
		x -= toFloat(e.RAM[cameraX:cameraX+2], true)
		y -= toFloat(e.RAM[cameraY:cameraY+2], true)
	}
	for r := 0.0; r < w; r++ {
		e.Pset(x+r, y, col)
		e.Pset(x+r, y+(h-1), col)
	}
	for c := 0.0; c < h; c++ {
		e.Pset(x, y+c, col)
		e.Pset(x+(w-1), y+(c), col)
	}
}

// RectFill draws a filled rectangle one the screen
func (e *Engine) RectFill(x, y, w, h float64, col Col, fixed ...bool) {
	f := false
	if len(fixed) > 0 {
		f = fixed[0]
	}
	if !f {
		x -= toFloat(e.RAM[cameraX:cameraX+2], true)
		y -= toFloat(e.RAM[cameraY:cameraY+2], true)
	}
	for r := 0.0; r < w; r++ {
		for c := 0.0; c < h; c++ {
			e.Pset(x+r, y+c, col)
		}
	}
}

// Line draws a colored line
func (e *Engine) Line(x1, y1, x2, y2 float64, col Col, fixed ...bool) {
	f := false
	if len(fixed) > 0 {
		f = fixed[0]
	}
	if !f {
		x1 -= toFloat(e.RAM[cameraX:cameraX+2], true)
		x2 -= toFloat(e.RAM[cameraX:cameraX+2], true)
		y1 -= toFloat(e.RAM[cameraY:cameraY+2], true)
		y2 -= toFloat(e.RAM[cameraY:cameraY+2], true)
	}
	if x2 < x1 {
		x2, x1 = x1, x2
	}
	w := x2 - x1
	dh := (float64(y2) - float64(y1)) / float64(w)
	if w > 0 {
		for x := x1; x < x2; x++ {
			e.Pset(x, y1+dh*(x-x1), col)
		}
		return
	}
	if y2 < y1 {
		y2, y1 = y1, y2
	}
	h := y2 - y1
	dw := (float64(x2) - float64(x1)) / float64(h)
	if h > 0 {
		for y := y1; y < y2; y++ {
			e.Pset(x1+(dw*(y-y1)), y, col)
		}
	}
}

// Circ draws a circle using Bresenham's algorithm
func (e *Engine) Circ(xc, yc, r float64, col Col, fixed ...bool) {
	f := false
	if len(fixed) > 0 {
		f = fixed[0]
	}
	e.circ(xc, yc, r, col, false, f)
}

// CircFill draws a filled circle using Bresenham's algorithm
func (e *Engine) CircFill(xc, yc, r float64, col Col, fixed ...bool) {
	f := false
	if len(fixed) > 0 {
		f = fixed[0]
	}
	e.circ(xc, yc, r, col, true, f)
}

// drawCirc8 draws 8 points on a circle
func (e *Engine) drawCirc8(xc, yc, x, y float64, c Col, filled bool, fixed bool) {
	if filled {
		e.Line(xc+x, yc+y, xc+x, yc-y, c, fixed)
		e.Line(xc-x, yc+y, xc-x, yc-y, c, fixed)
		e.Line(xc+y, yc+x, xc+y, yc-x, c, fixed)
		e.Line(xc-y, yc+x, xc-y, yc-x, c, fixed)
	}
	if !fixed {
		xc -= toFloat(e.RAM[cameraX:cameraX+2], true)
		yc -= toFloat(e.RAM[cameraY:cameraY+2], true)
	}
	e.Pset(xc+x, yc+y, c)
	e.Pset(xc-x, yc+y, c)
	e.Pset(xc+x, yc-y, c)
	e.Pset(xc-x, yc-y, c)
	e.Pset(xc+y, yc+x, c)
	e.Pset(xc-y, yc+x, c)
	e.Pset(xc+y, yc-x, c)
	e.Pset(xc-y, yc-x, c)
}

func (e *Engine) circ(xc, yc, r float64, c Col, filled bool, fixed bool) {
	if r == 0 {
		return
	}
	x, y := 0.0, r
	d := 3 - 2*r
	e.drawCirc8(xc, yc, x, y, c, filled, fixed)
	for y >= x {
		x++
		if d > 0 {
			y--
			d = d + 4*(x-y) + 10
		} else {
			d = d + 4*x + 6
		}
		e.drawCirc8(xc, yc, x, y, c, filled, fixed)
	}
}

// Clip clips all functions that draw to the screen
func (e *Engine) Clip(x, y, w, h int) {
	e.RAM[clipX] = byte(x)
	e.RAM[clipY] = byte(y)
	e.RAM[clipW] = byte(w)
	e.RAM[clipH] = byte(h)
}

// RClip resets the screen cliping
func (e *Engine) RClip() {
	e.RAM[clipX] = 0
	e.RAM[clipY] = 0
	e.RAM[clipW] = 192
	e.RAM[clipH] = 192
}

// sets a pixel in abitrary memory
// buffBase is the start of the pixel buffer in memory
// pxlWidth is the width of the pixel buffer in pixels
func (e *Engine) pset(x, y float64, col Col, buffBase, pxlWidth int) {
	ix, iy := int(x), int(y)
	i := ix + iy*pxlWidth
	index := int(float64(i/4) / 2 * 3)
	pIndex := index + (2 - index%3)
	cshift := (ix % 4) * 2
	pshift := ix % 8
	color := byte(col&0b00000011) << cshift
	pallet := byte(col&0b00000100) >> 2 << pshift

	e.RAM[buffBase+index] &= (0b00000011 << cshift) ^ 0b11111111
	e.RAM[buffBase+index] |= color
	e.RAM[buffBase+pIndex] &= (0b00000001 << pshift) ^ 0b11111111
	e.RAM[buffBase+pIndex] |= pallet
}

// Pset sets a pixel on the screen
func (e *Engine) Pset(x, y float64, col Col) {
	if x < float64(e.RAM[clipX]) || x >= float64(e.RAM[clipX]+e.RAM[clipW]) ||
		y < float64(e.RAM[clipY]) || y >= float64(e.RAM[clipY]+e.RAM[clipH]) {
		return
	}
	e.pset(x, y, col, 0, 192)
}

// pget gets a pixel from abitrary memory
// buffBase is the start of the memory buffer
// pxlWidth is the width of the buffer in pixels
func (e *Engine) pget(x, y float64, buffBase, pxlWidth int) Col {
	ix, iy := int(x), int(y)
	i := ix + iy*pxlWidth
	index := int(float64(i/4) / 2 * 3)
	pIndex := index + (2 - index%3)
	cshift := (ix % 4) * 2
	pshift := ix % 8
	color := (e.RAM[buffBase+index] >> cshift) & 0b00000011
	pallet := (e.RAM[buffBase+pIndex] >> pshift) & 0b00000001

	return Col((color | (pallet << 2)) | 0b10000000)
}

// Pget gets the color of a pixel on the screen
func (e *Engine) Pget(x, y float64) Col {
	if x < 0 || x > 192 || y < 0 || y >= 192 {
		return Col0
	}
	return e.pget(x, y, 0, 192)
}

// PalA sets pallet A
func (e *Engine) PalA(pallet Pal) {
	e.RAM[screenPalSet] &= 0b00001111
	e.RAM[screenPalSet] |= byte(pallet << 4)
}

// PalB sets pallet B
func (e *Engine) PalB(pallet Pal) {
	e.RAM[screenPalSet] &= 0b11110000
	e.RAM[screenPalSet] |= byte(pallet)
}

// PalGet gets the currently set pallets
func (e *Engine) PalGet() (Pal, Pal) {
	return Pal(e.RAM[screenPalSet] >> 4), Pal(e.RAM[screenPalSet] & 0b00001111)
}

// Col is a screen color
type Col byte

// These are the pallet and color constants
const (
	Col0 = Col(0b10000000)
	Col1 = Col(0b10000001)
	Col2 = Col(0b10000010)
	Col3 = Col(0b10000011)
	Col4 = Col(0b10000100)
	Col5 = Col(0b10000101)
	Col6 = Col(0b10000110)
	Col7 = Col(0b10000111)
)

func color(pixel byte) Col {
	pixel &= 0b00000011
	pixel |= 0b10000000
	return Col(pixel)
}

// Pal is a screen pallet
type Pal byte

// The list of all pallets
const (
	Pal0  = Pal(0b00000000)
	Pal1  = Pal(0b00000001)
	Pal2  = Pal(0b00000010)
	Pal3  = Pal(0b00000011)
	Pal4  = Pal(0b00000100)
	Pal5  = Pal(0b00000101)
	Pal6  = Pal(0b00000110)
	Pal7  = Pal(0b00000111)
	Pal8  = Pal(0b00001000)
	Pal9  = Pal(0b00001001)
	Pal10 = Pal(0b00001010)
	Pal11 = Pal(0b00001011)
	Pal12 = Pal(0b00001100)
	Pal13 = Pal(0b00001101)
	Pal14 = Pal(0b00001110)
	Pal15 = Pal(0b00001111)
)

// toInt converts a byte array to an integer
// byte arrays from len 1 to 4 are supported
func toInt(b []byte, signed bool) int {
	pad := []byte{0, 0, 0, 0}
	l := len(b)
	neg := false
	if signed && b[0] > 127 {
		neg = true
	}

	for i := 0; i < 4; i++ {
		if l-i-1 > -1 {
			pad[3-i] = b[l-i-1]
		}
	}

	if neg {
		pad[4-len(b)] &= 0b01111111
	}
	ret := int(pad[0])<<24 | int(pad[1])<<16 | int(pad[2])<<8 | int(pad[3])
	if neg {
		return ret * -1
	}
	return ret
}

// toFloat converts a byte arra to float64
// byte arrays from len 1 to 4 are supported
func toFloat(b []byte, signed bool) float64 {
	return float64(toInt(b, signed))
}

// toBytes converts an integer into a byte array
// signed bytes do not use twos complement
func toBytes(i int, l int, signed bool) []byte {
	neg := false
	if signed && i < 0 {
		i *= -1
		neg = true
	}

	ret := []byte{}
	if l == 1 {
		ret = []byte{byte(i)}
	}
	if l == 2 {
		ret = []byte{byte(i >> 8), byte(i)}
	}
	if l == 3 {
		ret = []byte{byte(i >> 16), byte(i >> 8), byte(i)}
	}
	if l == 4 {
		ret = []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
	}

	if neg {
		ret[0] |= 0b10000000
	}
	return ret
}
