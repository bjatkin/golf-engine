package golf

import (
	"strings"
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

// NewEngine creates a new golf engine
func NewEngine(updateFunc func(), draw func()) *Engine {
	ret := Engine{
		RAM:           &[0xFFFF]byte{},
		Draw:          draw,
		Update:        updateFunc,
		screenBufHook: js.Global().Get("screenBuff"),
	}

	ret.initKeyListener(js.Global().Get("document"))
	ret.initMouseListener(js.Global().Get("golfcanvas"))

	ret.RClip() // Reset the cliping box

	// Set internal resources
	base := internalSpriteColBase
	for i := 0; i < 0x0900; i++ {
		ret.RAM[i+base] = internalSpriteSheet[i]
	}

	//TODO inject the custom javascritp into the page here

	return &ret
}

// Run starts the game engine running
func (e *Engine) Run() {
	var renderFrame js.Func

	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e.addFrame()

		e.Update()
		e.Draw()

		e.drawMouse()

		e.tickKeyboard()
		e.tickMouse()

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
	e.setActiveSpriteBuff(internalSpriteColBase, internalSpritePalBase)

	cursor := e.RAM[mouseBase] >> 6
	opt := SprOpts{Fixed: true}

	if cursor == 1 {
		e.Spr(18, int(e.RAM[mouseX]), int(e.RAM[mouseY]), opt)
	}
	if cursor == 2 {
		e.Spr(50, int(e.RAM[mouseX]), int(e.RAM[mouseY]), opt)
	}
	if cursor == 3 {
		e.Spr(82, int(e.RAM[mouseX]), int(e.RAM[mouseY]), opt)
	}

	e.setActiveSpriteBuff(spriteColBase, spritePalBase)
}

// Mouse returns the X, Y coords of the mouse
func (e *Engine) Mouse() (int, int) {
	return int(e.RAM[mouseX]), int(e.RAM[mouseY])
}

// BG sets the bg color of the engine
func (e *Engine) BG(col Col) {
	e.RAM[bgColor] &= 0b00011111
	e.RAM[bgColor] |= byte(col << 5)
}

// Cls clears the screen and resets TextL and TextR
func (e *Engine) Cls() {
	textLline, textRline = 0, 0

	c := (e.RAM[bgColor] >> 5) & 0b00000011
	colBG := (c << 6) | (c << 4) | (c << 2) | c
	palBG := byte(0)
	if e.RAM[bgColor]>>7 == 1 {
		palBG = 0b11111111
	}

	for i := 0; i < 0x2400; i++ {
		e.RAM[i] = colBG
	}
	for i := screenPalBuffBase; i < screenPalSet; i++ {
		e.RAM[i] = palBG
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
func (e *Engine) Rect(x, y, w, h int, col Col) {
	x -= toInt(e.RAM[cameraX:cameraX+2], true)
	y -= toInt(e.RAM[cameraY:cameraY+2], true)
	for r := 0; r < w; r++ {
		e.Pset(x+r, y, col)
		e.Pset(x+r, y+(h-1), col)
	}
	for c := 0; c < h; c++ {
		e.Pset(x, y+c, col)
		e.Pset(x+(w-1), y+c, col)
	}
}

// RectFill draws a filled rectangle one the screen
func (e *Engine) RectFill(x, y, w, h int, col Col) {
	x -= toInt(e.RAM[cameraX:cameraX+2], true)
	y -= toInt(e.RAM[cameraY:cameraY+2], true)
	for r := 0; r < w; r++ {
		for c := 0; c < h; c++ {
			e.Pset(r+x, c+y, col)
		}
	}
}

// Line draws a colored line
func (e *Engine) Line(x1, y1, x2, y2 int, col Col) {
	x1 -= toInt(e.RAM[cameraX:cameraX+2], true)
	x2 -= toInt(e.RAM[cameraX:cameraX+2], true)
	y1 -= toInt(e.RAM[cameraY:cameraY+2], true)
	y2 -= toInt(e.RAM[cameraY:cameraY+2], true)
	if x2 < x1 {
		x2, x1 = x1, x2
	}
	w := x2 - x1
	dh := (float64(y2) - float64(y1)) / float64(w)
	if w > 0 {
		for x := x1; x < x2; x++ {
			e.Pset(x, y1+int(dh*float64(x-x1)), col)
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
			e.Pset(x1+int(dw*float64(y-y1)), y, col)
		}
	}
}

// drawCirc8 draws 8 points on a circle
func (e *Engine) drawCirc8(xc, yc, x, y int, c Col, filled bool) {
	if filled {
		e.Line(xc+x, yc+y, xc+x, yc-y, c)
		e.Line(xc-x, yc+y, xc-x, yc-y, c)
		e.Line(xc+y, yc+x, xc+y, yc-x, c)
		e.Line(xc-y, yc+x, xc-y, yc-x, c)
	} else {
		e.Pset(xc+x, yc+y, c)
		e.Pset(xc-x, yc+y, c)
		e.Pset(xc+x, yc-y, c)
		e.Pset(xc-x, yc-y, c)
		e.Pset(xc+y, yc+x, c)
		e.Pset(xc-y, yc+x, c)
		e.Pset(xc+y, yc-x, c)
		e.Pset(xc-y, yc-x, c)
	}
}

// Circ draws a circle using Bresenham's algorithm
func (e *Engine) Circ(xc, yc, r int, c Col) {
	e.circ(xc, yc, r, c, false)
}

// CircFill draws a filled circle using Bresenham's algorithm
func (e *Engine) CircFill(xc, yc, r int, c Col) {
	e.circ(xc, yc, r, c, true)
}

func (e *Engine) circ(xc, yc, r int, c Col, filled bool) {
	if r == 0 {
		return
	}
	x, y := 0, r
	d := 3 - 2*r
	e.drawCirc8(xc, yc, x, y, c, filled)
	for y >= x {
		x++
		if d > 0 {
			y--
			d = d + 4*(x-y) + 10
		} else {
			d = d + 4*x + 6
		}
		e.drawCirc8(xc, yc, x, y, c, filled)
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
// cBase is the start of the pallet memory buffer
// pBase is the start of the color memory buffer
func (e *Engine) pset(x, y int, col Col, cBase, pBase int) {
	cshift := x % 4
	pshift := x % 8
	i := (x / 4) + y*48
	j := (x / 8) + y*24
	pixel := e.RAM[cBase+i]

	masks := []byte{
		0b11111100,
		0b11110011,
		0b11001111,
		0b00111111,
	}

	newCol := byte(col&0b00000011) << (cshift * 2)
	newPix := (pixel & masks[cshift]) | newCol
	e.RAM[cBase+i] = newPix

	newPal := byte((col&0b00000100)>>2) << pshift
	e.RAM[pBase+j] &= ((0b00000001 << pshift) ^ 0b11111111)
	e.RAM[pBase+j] |= newPal
}

// Pset sets a pixel on the screen
func (e *Engine) Pset(x, y int, col Col) {
	if x < int(e.RAM[clipX]) || x >= int(e.RAM[clipX]+e.RAM[clipW]) ||
		y < int(e.RAM[clipY]) || y >= int(e.RAM[clipY]+e.RAM[clipH]) {
		return
	}
	e.pset(x, y, col, 0, screenPalBuffBase)
}

// gets a pixel from abitrary memory
// cBase is the start of the pallet memory buffer
// pBase is the start of the color memory buffer
func (e *Engine) pget(x, y, cBase, pBase int) Col {
	cshift := x % 4
	pshift := x % 8
	i := (x / 4) + y*64
	j := (x / 8) + y*32
	pixel := e.RAM[cBase+i]
	pal := (e.RAM[pBase+j] >> pshift) & 0b00000001

	masks := []byte{
		0b00000011,
		0b00001100,
		0b00110000,
		0b11000000,
	}
	pixel &= masks[cshift]
	pixel >>= (cshift * 2)

	return Col((pixel | (pal << 2)) | 0b10000000)
}

// Pget gets the color of a pixel on the screen
func (e *Engine) Pget(x, y int) Col {
	if x < 0 || x > 192 || y < 0 || y >= 192 {
		return Col0
	}
	return e.pget(x, y, 0, screenPalBuffBase)
}

// TextOpts additional options for drawing text
type TextOpts struct {
	Col   Col
	Solid bool
	Fixed bool
}

// the TextLine for TextL and TextR
var textLline, textRline = 0, 0

// TextL prints text at the top left of the screen
// the cursor moves to a new line each time TextL is called
func (e *Engine) TextL(text string, opts ...TextOpts) {
	splitText := strings.Split(text, "\n")
	for _, line := range splitText {
		if len(opts) > 0 {
			e.Text(line, 1, 1+6*textLline, opts[0])
		} else {
			e.Text(line, 1, 1+6*textLline)
		}
		textLline++
	}
}

// TextR prints text at the top right of the screen
// the cursor moves to a new line each time TextR is called
func (e *Engine) TextR(text string, opts ...TextOpts) {
	splitText := strings.Split(text, "\n")
	for _, line := range splitText {
		x := ScreenWidth - 1 - len(line)*6
		if len(opts) > 0 {
			e.Text(line, x, 1+6*textRline, opts[0])
		} else {
			e.Text(line, x, 1+6*textRline)
		}
		textRline++
	}
}

// the Text font reference
const textRef = "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_+={}[]\\:;\"<>,./?|"
const btnRef = "(<)(>)(^)(v)(x)(o)(l)(r)(+)(-)"
const specialRef = ":):(x(:|=[|^|v<-->$$oo<|<3<4+1-1pi()[]:;**"

// Text prints text at the x, y coords on the screen
func (e *Engine) Text(text string, x, y int, opts ...TextOpts) {
	text = strings.ToLower(text)
	px, py := x, y
	opt := TextOpts{}
	sopt := SprOpts{Transparent: Col3}
	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt.Solid {
		sopt.Transparent = 0
	}
	if opt.Col != 0 {
		sopt.PalFrom = []Col{Col0}
		sopt.PalTo = []Col{opt.Col}
	}
	sopt.Fixed = opt.Fixed

	e.setActiveSpriteBuff(internalSpriteColBase, internalSpritePalBase)

	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			px = x
			py += 6
			continue
		}
		if text[i] == ' ' {
			px += 6
			continue
		}
		if text[i] == '^' {
			i++
			dex := strings.Index(textRef, string(text[i]))
			e.drawChar(px, py, dex, sopt)
			px += 6
			continue
		}
		bdex := -1
		if i+2 < len(text) {
			bdex = strings.Index(btnRef, string(text[i:i+3]))
		}
		if bdex%3 == 0 {
			bdex = bdex/3 + 86
			e.drawChar(px, py, bdex, sopt)
			px += 6
			i += 2
			continue
		}
		sdex := -1
		if i+1 < len(text) {
			sdex = strings.Index(specialRef, string(text[i:i+2]))
		}
		if sdex%2 == 0 {
			sdex = sdex/2 + 65
			e.drawChar(px, py, sdex, sopt)
			px += 6
			i++
			continue
		}
		dex := -1
		if i < len(text) {
			dex = strings.Index(textRef, string(text[i]))
		}
		e.drawChar(px, py, dex, sopt)
		px += 6
	}

	e.setActiveSpriteBuff(spriteColBase, spritePalBase)
}

func (e *Engine) drawChar(x, y, i int, opt SprOpts) {
	if i < 0 {
		return
	}
	sx := i % 24
	sy := i / 24
	e.SSpr(sx*6, sy*6, 6, 6, x, y, opt)
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
