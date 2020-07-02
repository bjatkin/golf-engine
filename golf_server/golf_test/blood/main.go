package main

import (
	"fantasyConsole/golf_server/golf_test/golf"
)

var g *golf.Engine

func main() {
	g = golf.NewEngine(updateStart, drawStart)

	g.BG(golf.Col3)
	g.LoadSprs(spriteSheet)
	g.Run()
}

var fadeText = golf.TextOpts{Col: golf.Col3}
var fadeOutStart = 0
var bgCol = golf.Col3

func updateStart() {
	if (g.Frames()+1)%5 == 0 {
		fadeText.Col--
		if fadeText.Col < golf.Col0 {
			fadeText.Col = golf.Col0
		}
	}

	if g.Btnp(golf.XKey) && fadeOutStart == 0 {
		fadeOutStart = g.Frames()
	}

	if fadeOutStart > 0 && (g.Frames()-fadeOutStart+1)%5 == 0 {
		bgCol--
		g.BG(bgCol)
		if bgCol == golf.Col0 {
			g.Update = updateIntro
			g.Draw = drawIntro
		}
	}
}

func drawStart() {
	g.Cls()

	g.Text(75, 50, "oobl^oodoo")
	g.Text(50, 120, "press (x) to begin", fadeText)
}
