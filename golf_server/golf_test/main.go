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

func update(tdiff float64) {
	x++
	x %= 192
}

func draw() {
	g.Cls()
	g.Pset(x, 10, golf.Col3)
}
