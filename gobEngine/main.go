// +build js

package main

import (
	gob "fantasyConsole/terminal/GoB"
	"fmt"
	"time"
)

var engine gob.Engine
var uih uiHandler

func main() {
	engine = gob.NewEngine(terminalUpdate, terminalDraw)
	uih = buildEditorUI()
	engine.Load(&Resource)
	engine.Run()
}

var showFPS = false
var draw = [][]int{}

func editorUpdate(tdiff float64) {
	engine.DrawMouse = true
	if engine.Btnp(gob.EscKey) {
		engine.Update = terminalUpdate
		engine.Draw = terminalDraw
	}

	x, y := engine.Mouse()
	if engine.Mbtnp(gob.LeftClick) {
		uih.click(x, y, gob.LeftClick)
	}
	if engine.Mbtn(gob.LeftClick) && x < 256 && y-8 < 256 {
		engine.Resource.Pset((x/16)+selSprX*16, ((y-8)/16)+selSprY*16, selCol1)
	}
	if engine.Mbtnp(gob.RightClick) {
		uih.click(x, y, gob.RightClick)
	}
	if engine.Mbtn(gob.RightClick) && x < 256 && y-8 < 256 {
		engine.Resource.Pset((x/16)+selSprX*16, ((y-8)/16)+selSprY*16, selCol2)
	}
}

func editorDraw() {
	engine.Cls()

	// Editor Area
	engine.Rect(0, 8, 256, 256, gob.Col7)
	// Scrol Bars
	engine.Rect(256, 8, 8, 256, gob.Col7)
	engine.Rect(0, 264, 256, 8, gob.Col7)

	// Sprite Select
	engine.Rect(288, 8, 32, 8, gob.Col7)
	engine.Rect(288, 16, 32, 8, gob.Col7)
	engine.Rect(288, 24, 32, 256, gob.Col7)

	// Color Picker
	uih.draw()

	// Selected Colors
	engine.RectFill(272, 123, 16, 16, selCol2)
	engine.RectFill(264, 115, 16, 16, selCol1)

	for _, p := range draw {
		engine.Pset(p[0], p[1], gob.Col15)
	}

	// Draw Resource File
	for x := 0; x < 32; x++ {
		for y := 0; y < 256; y++ {
			engine.Pset(x+288, y+32, engine.Resource.Pget(x, y))
		}
	}

	// Header and Footer
	engine.RectFill(0, 0, gob.ScreenWidth, 8, gob.Col15)
	engine.Text("Sprites", 2, 1, gob.Col0)
	engine.RectFill(0, gob.ScreenHeight-8, gob.ScreenWidth, 8, gob.Col15)
	engine.Line(0, 8, gob.ScreenWidth, 8, gob.Col1)
	x, y := engine.Mouse()
	engine.Text(fmt.Sprintf("(%d, %d)", x, y), 1, 280, gob.Col0)

	//SelSpr
	engine.Rect(288+selSprX*16-1, 32+selSprY*16-1, 18, 18, gob.Col7)

	drawEditor()
	drawFPS()
}

func drawEditor() {
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			col := engine.Resource.Pget(x+selSprX*16, y+selSprY*16)
			engine.RectFill(x*16, 8+y*16, 16, 16, col)
		}
	}
}

var frames, tMark, lastCheck, fps = 0, 0.0, int64(0), 0.0

// Add this to the gob engine
func drawFPS() {
	tMark++
	now := time.Now().UnixNano()
	if now >= lastCheck+250000000 {
		lastCheck = time.Now().UnixNano()
		fps = tMark / 0.25
		tMark = 0
	}

	if !showFPS {
		return
	}
	x, y := engine.Mouse()
	if x > 256 && y < 11 && engine.DrawMouse {
		return
	}

	engine.RectFill(257, 0, 63, 10, gob.Col15)
	engine.Rect(257, 0, 63, 10, gob.Col7)
	engine.TextR(fmt.Sprintf("FPS: %.2f", fps), gob.Col0)
}

var selCol1, selCol2 = gob.Col15, gob.Col7
var selSprX, selSprY = 0, 0

func buildEditorUI() uiHandler {
	ret := uiHandler{}

	c1, c2 := gob.Col15, gob.Col7
	size := 12
	for r := 0; r < 8; r++ {
		x, y := 264, 8+12*r
		// we need to cpy the values or draw will not work correctly
		nc1, nc2 := c1, c2
		ret.add(
			ui{
				name: fmt.Sprintf("color select %d", nc1),
				x:    x,
				y:    y,
				w:    size,
				h:    size,
				draw: func() {
					engine.RectFill(x, y, size, size, nc1)
					engine.Rect(x, y, size, size, gob.Col4)
				},
				click: func(self ui, btn gob.MouseBtn) {
					if btn == gob.LeftClick {
						selCol1 = nc1
					}
					if btn == gob.RightClick {
						selCol2 = nc1
					}
				},
			},
			ui{
				name: fmt.Sprintf("color select %d", nc2),
				x:    x + size,
				y:    y,
				w:    size,
				h:    size,
				draw: func() {
					engine.RectFill(x+size, y, size, size, nc2)
					engine.Rect(x+size, y, size, size, gob.Col12)
				},
				click: func(self ui, btn gob.MouseBtn) {
					if btn == gob.LeftClick {
						selCol1 = nc2
					}
					if btn == gob.RightClick {
						selCol2 = nc2
					}
				},
			},
		)
		c1--
		c2--
	}

	for x := 0; x < 2; x++ {
		for y := 0; y < 16; y++ {
			nx, ny := x, y
			ret.add(
				ui{
					name: fmt.Sprintf("sprite select %d, %d", x, y),
					x:    288 + x*16,
					y:    32 + y*16,
					w:    16,
					h:    16,
					draw: func() {},
					click: func(ui, gob.MouseBtn) {
						selSprX = nx
						selSprY = ny
					},
				},
			)
		}
	}

	return ret
}
