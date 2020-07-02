package main

type player struct {
	x, y  int
	state int
	a     *anim
}

var mainPlayer = player{
	a: wait,
}

func (p *player) tick() {
	p.a.tick()
}
