package main

import (
	"fantasyConsole/golf_server/golf_test/golf"
)

var g *golf.Engine

func main() {
	g = golf.NewEngine(update, draw)

	g.BG(golf.Col3)
	g.Run()
}

func update() {}

func draw() {
	g.Cls()

	g.Text(60, 50, "hello world!")
}