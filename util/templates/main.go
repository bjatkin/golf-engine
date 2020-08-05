package main

import (
	"github.com/bjatkin/golf-engine/golf"
)

var g *golf.Engine

func main() {
	g = golf.NewEngine(update, draw)

	// g.LoadSprs(spriteSheet)
	// g.LoadMap(mapData)
	// g.LoadFlags(flagData)

	g.Run()
}

func update() {}

func draw() {
	g.Cls(golf.Col7)

	g.Text(60, 50, "hello world!")
}
