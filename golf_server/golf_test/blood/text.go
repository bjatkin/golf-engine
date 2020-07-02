package main

import (
	"fantasyConsole/golf_server/golf_test/golf"
)

type textBlrb struct {
	text         []string
	letter, line int
	x, y         int
	px, py       int
	tickCounter  int
	end          bool
	opts         golf.TextOpts
}

func (t *textBlrb) draw() {
	if t.end {
		return
	}
	print := t.text[t.line][:t.letter]
	g.Text(t.x, t.y, print, t.opts)
	if len(print) == len(t.text[t.line]) {
		g.Text(t.px, t.py, "press (x)", t.opts)
	}
}

func (t *textBlrb) tick() {
	if t.end {
		return
	}
	t.tickCounter++
	if t.tickCounter > 5 {
		t.tickCounter = 0
		if t.letter < len(t.text[t.line]) {
			t.letter++
		}
	}
}

func (t *textBlrb) next() {
	t.letter = 0
	t.line++
	t.tickCounter = 0
	if t.line >= len(t.text) {
		t.end = true
	}
}
