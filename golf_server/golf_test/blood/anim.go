package main

type anim struct {
	frames []int
	speed  int
	i      int
	frame  int
	width  int
}

func (a *anim) tick() {
	a.i++
	if a.i > a.speed {
		a.i = 0
		a.frame++
	}
	if a.frame >= len(a.frames) {
		a.frame = 0
	}
}

func (a *anim) f() int {
	if a.frames[a.frame] < 0 {
		return -1 * a.frames[a.frame]
	}
	return a.frames[a.frame]
}

func (a *anim) flip() bool {
	return a.frames[a.frame] < 0
}
