package gob

import "syscall/js"

// MouseKey codes
const (
	LeftClick   = MouseBtn(0)
	MiddleClick = MouseBtn(1)
	RightClick  = MouseBtn(2)
)

type mouseListener struct {
	x, y               int
	new, old, released []MouseBtn
}

// MouseBtn is a mouse key
type MouseBtn int

func newMouseListener(canvas js.Value) *mouseListener {
	ret := mouseListener{}
	mouseMove := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		height := canvas.Get("offsetHeight").Int()
		ratio := float64(ScreenHeight) / float64(height)
		ret.x = int(args[0].Get("offsetX").Float() * ratio)
		ret.y = int(args[0].Get("offsetY").Float() * ratio)

		return nil
	})
	mouseDown := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := MouseBtn(args[0].Get("button").Int())

		// Add to new if this is a new mouse key
		add := true
		for _, key := range ret.new {
			if key == e {
				add = false
			}
		}

		for _, key := range ret.old {
			if key == e {
				add = false
			}
		}

		if add {
			ret.new = append(ret.new, e)
		}

		return nil
	})
	mouseUp := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := MouseBtn(args[0].Get("button").Int())
		ret.released = append(ret.new, e)

		// Remove from old and new
		for i, key := range ret.new {
			if key == e {
				ret.new[i] = ret.new[len(ret.new)-1]
				ret.new = ret.new[:len(ret.new)-1]
				break
			}
		}

		for i, key := range ret.old {
			if key == e {
				ret.old[i] = ret.old[len(ret.old)-1]
				ret.old = ret.old[:len(ret.old)-1]
				break
			}
		}
		return nil
	})
	canvas.Call("addEventListener", "mousemove", mouseMove)
	canvas.Call("addEventListener", "mousedown", mouseDown)
	canvas.Call("addEventListener", "mouseup", mouseUp)

	return &ret
}

func (ml *mouseListener) tick() {
	// Move Keys from new to old
	var add bool
	for _, key := range ml.new {
		add = true
		for _, Okey := range ml.old {
			if key == Okey {
				add = false
				break
			}
		}
		if add {
			ml.old = append(ml.old, key)
		}
	}

	ml.new = []MouseBtn{}
	ml.released = []MouseBtn{}
}

// Mbtn returns true is the mouse key is being pressed
func (e *Engine) Mbtn(key MouseBtn) bool {
	if e.Mbtnp(key) {
		return true
	}
	for _, Okey := range e.ml.old {
		if Okey == key {
			return true
		}
	}
	return false
}

// Mbtnp returns true if the mouse key was pressed this frame
func (e *Engine) Mbtnp(key MouseBtn) bool {
	for _, Nkey := range e.ml.new {
		if Nkey == key {
			return true
		}
	}
	return false
}

// Mbtnr returns true if the mouse key was released this frame
func (e *Engine) Mbtnr(key MouseBtn) bool {
	for _, Rkey := range e.ml.released {
		if key == Rkey {
			return true
		}
	}
	return false
}
