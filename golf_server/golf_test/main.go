package main

import (
	"fantasyConsole/golf_server/golf_test/golf"
)

var g *golf.Engine

func main() {
	g = golf.NewEngine(update, draw)
	g.BG(golf.Col1)
	g.Run()
}

var x = 0

func update() {
	x++
	x %= 192
}

func draw() {
	g.Cls()
	g.Pset(x, 10, golf.Col3)
	for i := 0; i < 192; i++ {
		g.Pset(i, i, golf.Col0)
	}
	g.Rect(10, 10, 10, 10, golf.Col0)
	g.RectFill(50, 15, 10, 20, golf.Col0)
	x, y := g.Mouse()
	g.Pset(x, y, golf.Col3)
}
