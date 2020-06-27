package main

import (
	"fantasyConsole/golf_server/golf_test/golf"
	"fmt"
	"math"
	"time"
)

var g *golf.Engine

func main() {
	// g = golf.NewEngine(bibiDuckUpdate, bibiDuckDraw)
	// g.BG(golf.Col5)
	// g.PalA(5)
	// g.PalB(6)
	// g.LoadSprs(spriteSheet)
	// g.Run()

	lastFrameTime = time.Now().UnixNano()
	g = golf.NewEngine(update, draw)

	g.LoadSprs(spriteSheet)
	g.BG(golf.Col3)
	g.DrawMouse(1)
	g.Run()
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
	if g.Btnp(golf.PKey) {
		a, _ := g.PalGet()
		a++
		if a > golf.Pal15 {
			a = golf.Pal0
		}
		g.PalA(a)
		g.PalB(a)
	}
	if g.Btnp(golf.RKey) {
		cx, cy = 0, 0
	}
	if g.Btnp(golf.Space) {
		if clipped {
			g.RClip()
		} else {
			g.Clip(10, 50, 172, 80)
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
	g.SSpr(152, 0, 64, 24, 64, 64, golf.SprOpts{Transparent: golf.Col7})

	// Draw a scaled sprite
	scale := (math.Sin(float64(g.Frames()) / 60)) + 1
	g.Spr(82, 4, 13, golf.SprOpts{
		ScaleW: scale * 2, ScaleH: scale * 2,
		Transparent: golf.Col2,
	})
	g.TextL(fmt.Sprintf("Scale: %.2f", scale))

	// Draw the current pallet
	pal, _ := g.PalGet()
	g.TextR(fmt.Sprintf("\nPal: %d ", pal))
}

var bgCol = golf.Col0

func demo2Update() {
	if g.Btnp(golf.Space) {
		bgCol++
		if bgCol > golf.Col7 {
			bgCol = golf.Col0
		}
		g.BG(bgCol)
		fmt.Printf("Swap BGCol: %v\n", bgCol)
	}
}

func demo2Draw() {
	g.Cls()
	for i := 0; i < 192; i++ {
		g.Pset(i, i, golf.Col0)
	}
}

func bibiDuckUpdate() {

}

func bibiDuckDraw() {
	g.Cls()
	g.Spr(1, 50, 50, golf.SprOpts{Width: 2, Height: 2, Transparent: golf.Col7})
}
