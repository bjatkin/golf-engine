package main

import gob "fantasyConsole/terminal/GoB"

type uiHandler struct {
	elements []ui
}

type ui struct {
	name  string
	x, y  int
	w, h  int
	draw  func()
	click func(ui, gob.MouseBtn)
	hover func(ui)
}

func newUIHandler() {

}

func (h *uiHandler) add(elements ...ui) {
	for _, e := range elements {
		h.elements = append(h.elements, e)
	}
}

func (h *uiHandler) remove(names ...string) {
	for _, name := range names {
		for i := 0; i < len(h.elements); i++ {
			if h.elements[i].name == name {
				h.elements[i] = h.elements[len(h.elements)-1]
				h.elements = h.elements[:len(h.elements)-1]
			}
		}
	}
}

func (h *uiHandler) draw() {
	for _, e := range h.elements {
		e.draw()
	}
}

func (h *uiHandler) click(x, y int, btn gob.MouseBtn) {
	for _, e := range h.elements {
		if x >= e.x && y > e.y && x < e.x+e.w && y < e.y+e.h {
			e.click(e, btn)
		}
	}
}

func (h *uiHandler) hover(x, y int, btn gob.MouseBtn) {
	for _, e := range h.elements {
		if x >= e.x && y > e.y && x < e.x+e.w && y < e.y+e.h {
			e.hover(e)
		}
	}
}
