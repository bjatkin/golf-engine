package main

import (
	"fantasyConsole/golf_server/golf_test/golf"
	"fmt"
	"time"
)

var g *golf.Engine

func main() {
	lastFrameTime = time.Now().UnixNano()
	g = golf.NewEngine(update, draw)
	g.BG(golf.Col3)
	g.DrawMouse(1)
	g.LoadSprs(spriteSheet)
	g.PalA(5)
	g.PalB(4)
	g.Run()
}

var cx, cy int
var clipped, palDiff bool

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
		if palDiff {
			g.PalA(1)
			g.PalB(0)
		} else {
			g.PalA(5)
			g.PalB(4)
		}
		palDiff = !palDiff
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

	g.Mset(0, 0, 1)
	g.Mset(1, 0, 2)
	g.Mset(0, 1, 33)
	g.Mset(1, 1, 34)
	g.Map(0, 0, 256, 64, 0, 0, golf.SprOpts{Transparent: golf.Col2})
	g.Camera(cx, cy)
	g.Rect(1, 10, 190, 181, golf.Col0)
	g.TextR(fmt.Sprintf("FPS: %d", frameDiff), golf.TextOpts{Col: golf.Col0, Fixed: true})
	g.RectFill(64, 90, 64, 5, golf.Col2)
	g.Line(74, 97, 118, 97, golf.Col2)
	g.Line(90, 99, 102, 99, golf.Col2)
	g.SSpr(152, 0, 64, 24, 64, 64)

	g.SSpr(0, 0, 256, 64, 0, 64, golf.SprOpts{Transparent: golf.Col2})
	for i := 0; i < 16; i++ {
		g.Spr(15, i*16, 40, golf.SprOpts{Width: 2, Height: 6})
	}
	g.Spr((g.Frames()/30%2)*2+1, 50, 27, golf.SprOpts{Width: 2, Height: 2, Transparent: golf.Col2, Fixed: true})
	frames := []int{65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 67, 69, 67, 69, 67}
	g.Spr(frames[(g.Frames()/5%20)], 70, 10, golf.SprOpts{Width: 2, Height: 2, Transparent: golf.Col2})
	g.Text("BiBi Duck!", 60, 180, golf.TextOpts{Col: golf.Col0})
}
