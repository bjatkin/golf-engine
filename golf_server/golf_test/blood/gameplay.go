package main

import (
	"fantasyConsole/golf_server/golf_test/golf"
	"fmt"
)

var whiteText = golf.TextOpts{Col: golf.Col3}
var openingText = textBlrb{
	x: 30, y: 75,
	px: 120, py: 95,
	text: []string{"I wander alone...           \nunwanted...   ",
		"A nomad of \ntwo worlds...   ",
		"born of evil...   ",
		"seeking redemption...   ",
		"perhaps this time will           \nbe different...   "},
	opts: whiteText,
}

func updateIntro() {
	openingText.tick()
	if g.Btnp(golf.XKey) {
		openingText.next()
	}
	if openingText.end {
		g.Update = updateGame
		g.Draw = drawGame
	}
}

func drawIntro() {
	g.Cls()

	openingText.draw()
}

var wait = &anim{
	frames: []int{2, 3},
	speed:  60,
	width:  1,
}
var iwait = &anim{
	frames: []int{-2, -3},
	speed:  60,
	width:  1,
}
var sideWalk = &anim{
	frames: []int{6, 8, 10, 12},
	speed:  5,
	width:  2,
}
var isideWalk = &anim{
	frames: []int{-6, -8, -10, -12},
	speed:  5,
	width:  2,
}
var upWalk = &anim{
	frames: []int{14, 16, 18, 20, -14, -16, -18, -20},
	speed:  5,
	width:  2,
}

var speed = 2
var faceRight = false

func updateGame() {
	move := false
	if g.Btn(golf.WKey) {
		mainPlayer.a = upWalk
		mainPlayer.y -= speed
		move = true
	}
	if g.Btn(golf.SKey) {
		mainPlayer.a = upWalk
		mainPlayer.y += speed
		move = true
	}
	if g.Btn(golf.AKey) {
		mainPlayer.a = sideWalk
		mainPlayer.x -= speed
		move = true
		faceRight = false
	}
	if g.Btn(golf.DKey) {
		mainPlayer.a = isideWalk
		mainPlayer.x += speed
		move = true
		faceRight = true
	}
	if !move && !faceRight {
		mainPlayer.a = wait
	}
	if !move && faceRight {
		mainPlayer.a = iwait
	}
	mainPlayer.tick()
}

func drawGame() {
	g.Cls()

	g.TextL(fmt.Sprintf("%d, %d", mainPlayer.x, mainPlayer.y))
	g.CircFill(99, 97, 3, golf.Col1)
	g.CircFill(102, 99, 2, golf.Col1)
	g.CircFill(105, 100, 1, golf.Col1)

	g.Spr(wait.f(), 96, 96, golf.SprOpts{Height: 2, Transparent: golf.Col7})
	g.Spr(sideWalk.f(), 96, 50, golf.SprOpts{Width: 2, Height: 2, Transparent: golf.Col7})
	g.Spr(upWalk.f(), 96, 70, golf.SprOpts{Width: 2, Height: 2, Transparent: golf.Col7, FlipH: upWalk.flip()})

	g.Spr(mainPlayer.a.f(), mainPlayer.x, mainPlayer.y,
		golf.SprOpts{Width: mainPlayer.a.width, Height: 2, Transparent: golf.Col7, FlipH: mainPlayer.a.flip()})
}
