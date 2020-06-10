package golf

import (
	"syscall/js"
)

// Addresses
// ScreenBuff: 0x0000 - 0x3601
//  Col Buff: 0x0000 - 0x2400
// 	Pal Buff: 0x2401 - 0x3600
//  Pal Set: 0x3601
// BG Color: 0x3602 - high 3 bits

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
	Update        func(float64)

	//TODO move these to the ram
	kl *keyListener
	ml *mouseListener
}

// NewEngine creates a new golf engine
func NewEngine(updateFunc func(float64), draw func()) *Engine {
	ret := Engine{
		RAM:           &[0xFFFF]byte{},
		Draw:          draw,
		Update:        updateFunc,
		screenBufHook: js.Global().Get("screenBuff"),
	}

	ret.kl = newKeyListener(js.Global().Get("document"), ret.RAM)
	ret.ml = newMouseListener(js.Global().Get("golfcanvas"), ret.RAM)

	return &ret
}

// Run starts the game engine running
func (e *Engine) Run() {
	var renderFrame js.Func

	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		// TODO update this to include tdiff
		// now := args[0].Float()
		e.Update(0.0)
		e.Draw()
		e.kl.tick()
		e.ml.tick()

		js.CopyBytesToJS(e.screenBufHook, e.RAM[:0x3601])
		js.Global().Call("drawScreen")
		js.Global().Call("requestAnimationFrame", renderFrame)

		return nil
	})

	done := make(chan struct{}, 0)

	js.Global().Call("requestAnimationFrame", renderFrame)
	<-done
}

// Time is the number of seconds since the engine was started
func (e *Engine) Time() float64 {
	return 0.0
}

// Mouse returns the X, Y coords of the mouse
func (e *Engine) Mouse() (int, int) {
	return 0, 0
}

// BG sets the bg color of the engine
func (e *Engine) BG(col Col) {
	e.RAM[0x3602] &= 0b00011111
	e.RAM[0x3602] |= byte(col << 5)
}

// Cls clears the screen and resets TextL and TextR
func (e *Engine) Cls() {
	c := e.RAM[0x3602] >> 5
	colBG := (c << 6) | (c << 4) | (c << 2) | c
	palBG := byte(0)
	if e.RAM[0x3602]>>7 == 1 {
		palBG = 0b11111111
	}

	for i := 0; i < 0x2400; i++ {
		e.RAM[i] = colBG
	}
	for i := 0x2401; i < 0x3600; i++ {
		e.RAM[i] = palBG
	}
}

// Camera moves the camera which modifies all draw functions
func (e *Engine) Camera(x, y int) {}

// Rect draws a rectangle border on the screen
func (e *Engine) Rect(x, y, w, h int, col Col) {}

// RectFill draws a filled rectangle one the screen
func (e *Engine) RectFill(x, y, w, h int, col Col) {}

// Line draws a colored line
func (e *Engine) Line(x1, y1, x2, y2 int, col Col) {}

// Clip clips all functions that draw to the screen
func (e *Engine) Clip(x, y, w, h int) {}

// RClip resets the screen cliping
func (e *Engine) RClip() {}

// Pset sets a pixel on the screen
func (e *Engine) Pset(x, y int, col Col) {
	cshift := x % 4
	pshift := x % 8
	i := (x / 4) + y*48
	j := (x / 8) + y*24
	pixel := e.RAM[i]

	masks := []byte{
		0b11111100,
		0b11110011,
		0b11001111,
		0b00111111,
	}

	newCol := byte(col&0b00000011) << (cshift * 2)
	newPix := (pixel & masks[cshift]) | newCol
	e.RAM[i] = newPix

	newPal := byte((col&0b00000100)>>2) << pshift
	e.RAM[j+0x2401] &= (0b00000001 << pshift)
	e.RAM[j+0x2401] |= newPal
}

// Pget gets the color of a pixel on the screen
func (e *Engine) Pget(x, y int) Col {
	cshift := x % 4
	pshift := x % 8
	i := (x / 4) + y*48
	j := (x / 8) + y*24
	pixel := e.RAM[i]
	pal := (e.RAM[j+0x2401] >> pshift) & 0b00000001

	masks := []byte{
		0b00000011,
		0b00001100,
		0b00110000,
		0b11000000,
	}
	pixel &= masks[cshift]
	pixel >>= (cshift * 2)

	return Col(pixel & (pal << 2))
}

// SprOpts additional options for drawing sprites
type SprOpts struct {
	FlipH         bool
	FlipV         bool
	Transparent   Col
	PalFrom       []Col
	PalTo         []Col
	Width, Height int
}

// Spr draws 8x8 sprite n from the sprite sheet to the
// screen at x, y.
func (e *Engine) Spr(n, x, y int, opts ...SprOpts) {}

// SSpr draw a rect from the sprite sheet to the screen
// sx, sy, sw, and sh define the rect on the sprite sheet
// dx, dy is the location to draw on the screen
func (e *Engine) SSpr(sx, sy, sw, sh, dx, dy int, opts ...SprOpts) {}

// subPixels is used to swap pixels based on a pallet swap
func subPixels(palFrom, palTo []Col, col Col) Col {
	return Col0
}

// TextOpts additional options for drawing text
type TextOpts struct {
	Transparent Col
	Col         Col
	Relative    bool
}

// TextL prints text at the top left of the screen
// the cursor moves to a new line each time TextL is called
func (e *Engine) TextL(text string, opts ...TextOpts) {}

// TextR prints text at the top right of the screen
// the cursor moves to a new line each time TextR is called
func (e *Engine) TextR(text string, opts ...TextOpts) {}

// Text prints text at the x, y coords on the screen
func (e *Engine) Text(text string, x, y int, opts ...TextOpts) {}

// Pal1 sets pallet one
func (e *Engine) Pal1(pallet Pal) {}

// Pal2 sets pallet one
func (e *Engine) Pal2(pallet Pal) {}

// PalGet gets the currently set pallets
func (e *Engine) PalGet() (Pal, Pal) {
	return Pal0, Pal0
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