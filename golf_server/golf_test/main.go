package main

import (
	"fantasyConsole/golf_server/golf_test/golf"
	"fmt"
	"math"
	"time"
)

var g *golf.Engine

func main() {
	g = golf.NewEngine(demo2Update, demo2Draw)
	g.PalA(0)
	g.PalB(3)
	g.Run()

	// lastFrameTime = time.Now().UnixNano()
	// g = golf.NewEngine(update, draw)
	// g.LoadSprs(spriteSheet)

	// g.BG(golf.Col3)
	// g.PalA(0)
	// g.PalB(3)

	// g.DrawMouse(1)

	// g.Run()
}

var cx, cy int
var clipped bool

func update() {
	if g.Btn(golf.WKey) {
		cy -= 10
	}
	if g.Btn(golf.AKey) {
		cx -= 10
	}
	if g.Btn(golf.SKey) {
		cy += 10
	}
	if g.Btn(golf.DKey) {
		cx += 10
	}
	if g.Btnp(golf.Space) {
		if clipped {
			g.RClip()
		} else {
			g.Clip(10, 30, 172, 152)
		}
		clipped = !clipped
	}
}

var lastFrameTime int64
var lastFrame int
var frameDiff int

func draw() {
	g.Cls()
	now := time.Now().UnixNano()
	nanoSec := int64(1000000000)
	if now-lastFrameTime > nanoSec/4 {
		lastFrameTime = now
		diff := g.Frames() - lastFrame
		lastFrame = g.Frames()
		frameDiff = diff * 4
	}

	g.Camera(cx, cy)
	g.Rect(1, 10, 190, 181, golf.Col0)
	g.TextR(fmt.Sprintf("FPS: %d", frameDiff), golf.TextOpts{Col: golf.Col0, Fixed: true})

	// Shadow
	g.RectFill(65, 90, 62, 5, golf.Col2)
	g.RectFill(64, 91, 64, 3, golf.Col2)
	g.Line(74, 97, 118, 97, golf.Col2)
	g.Line(90, 99, 102, 99, golf.Col2)

	// Change to the internal sprite sheet
	g.RAM[0x6F49], g.RAM[0x6F4A] = 0x36, 0x47
	g.RAM[0x6F4B], g.RAM[0x6F4C] = 0x3C, 0x47

	// Draw the logo
	g.SSpr(152, 0, 64, 24, 64, 64, golf.SprOpts{Transparent: golf.Col2})

	// Draw a scaled sprite
	scale := (math.Sin(float64(g.Frames()) / 60)) + 1
	g.Spr(82, 4, 13, golf.SprOpts{
		ScaleW: scale * 2, ScaleH: scale * 2,
		Transparent: golf.Col2,
	})
	g.TextL(fmt.Sprintf("Scale: %.2f", scale))
}

var bgCol = golf.Col0

func demo2Update() {
	if g.Btnp(golf.Space) {
		bgCol++
		if bgCol > golf.Col7 {
			bgCol = golf.Col0
		}
		g.BG(bgCol)
	}
}

func demo2Draw() {
	g.Cls()
}
