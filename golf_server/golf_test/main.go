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

	g.Camera(cx, cy)
	g.Rect(1, 10, 190, 181, golf.Col0)
	g.TextR(fmt.Sprintf("FPS: %d", frameDiff), golf.TextOpts{Col: golf.Col0, Fixed: true})
	g.RectFill(64, 90, 64, 5, golf.Col2)
	g.Line(74, 97, 118, 97, golf.Col2)
	g.Line(90, 99, 102, 99, golf.Col2)
	g.SSpr(152, 0, 64, 24, 64, 64)
}
