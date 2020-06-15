package golf

import (
	"syscall/js"
)

//TODO change mouse dest location

// MouseBtn is a mouse key
type MouseBtn int

// MouseKey codes
const (
	LeftClick   = MouseBtn(0)
	MiddleClick = MouseBtn(1)
	RightClick  = MouseBtn(2)
)

type mouseListener struct {
	ram *[0xFFFF]byte
}

// Mouse Data addresses
const (
	mouseBase = uint16(0x3610)
	mouseX    = uint16(0x360E)
	mouseY    = uint16(0x360F)
)

func (e *Engine) initMouseListener(canvas js.Value) {
	mouseMove := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		height := canvas.Get("offsetHeight").Int()
		ratio := float64(ScreenHeight) / float64(height)
		e.RAM[mouseX] = byte(args[0].Get("offsetX").Float() * ratio)
		e.RAM[mouseY] = byte(args[0].Get("offsetY").Float() * ratio)

		return nil
	})
	mouseDown := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		k := MouseBtn(args[0].Get("button").Int())

		if k == LeftClick {
			btn := btnState(e.RAM[mouseBase] & 0b00000011)
			if btn == unpressed {
				e.RAM[mouseBase] |= byte(start)
			}
		}
		if k == MiddleClick {
			btn := btnState((e.RAM[mouseBase] & 0b00001100) >> 2)
			if btn == unpressed {
				e.RAM[mouseBase] |= (byte(start) << 2)
			}
		}
		if k == RightClick {
			btn := btnState((e.RAM[mouseBase] & 0b00110000) >> 4)
			if btn == unpressed {
				e.RAM[mouseBase] |= (byte(start) << 4)
			}
		}

		return nil
	})
	mouseUp := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		k := MouseBtn(args[0].Get("button").Int())

		if k == LeftClick {
			btn := btnState(e.RAM[mouseBase] & 0b00000011)
			if btn == pressed || btn == start {
				e.RAM[mouseBase] &= 0b11111100
				e.RAM[mouseBase] |= byte(end)
			}
		}
		if k == MiddleClick {
			btn := btnState((e.RAM[mouseBase] & 0b00001100) >> 2)
			if btn == pressed || btn == start {
				e.RAM[mouseBase] &= 0b11110011
				e.RAM[mouseBase] |= (byte(end) << 2)
			}
		}
		if k == RightClick {
			btn := btnState((e.RAM[mouseBase] & 0b00110000) >> 4)
			if btn == pressed || btn == start {
				e.RAM[mouseBase] &= 0b11001111
				e.RAM[mouseBase] |= (byte(end) << 4)
			}
		}

		return nil
	})
	canvas.Call("addEventListener", "mousemove", mouseMove)
	canvas.Call("addEventListener", "mousedown", mouseDown)
	canvas.Call("addEventListener", "mouseup", mouseUp)
}

func (e *Engine) tickMouse() {
	// Move btn from start to pressed
	if btnState(e.RAM[mouseBase]&0b00000011) == start {
		e.RAM[mouseBase] &= 0b11111100
		e.RAM[mouseBase] |= byte(pressed)
	}
	if btnState((e.RAM[mouseBase]&0b00001100)>>2) == start {
		e.RAM[mouseBase] &= 0b11110011
		e.RAM[mouseBase] |= (byte(pressed) << 2)
	}
	if btnState((e.RAM[mouseBase]&0b00110000)>>4) == start {
		e.RAM[mouseBase] &= 0b11001111
		e.RAM[mouseBase] |= (byte(pressed) << 4)
	}

	// Move from end to unpressed
	if btnState(e.RAM[mouseBase]&0b00000011) == end {
		e.RAM[mouseBase] &= 0b11111100
	}
	if btnState((e.RAM[mouseBase]&0b00001100)>>2) == end {
		e.RAM[mouseBase] &= 0b11110011
	}
	if btnState((e.RAM[mouseBase]&0b00110000)>>4) == end {
		e.RAM[mouseBase] &= 0b11001111
	}
}

// Mbtn returns true is the mouse key is being pressed
func (e *Engine) Mbtn(key MouseBtn) bool {
	btn := btnState(e.RAM[mouseBase] & 0b00000011)
	if key == MiddleClick {
		btn = btnState(e.RAM[mouseBase] & 0b00001100)
	}
	if key == RightClick {
		btn = btnState(e.RAM[mouseBase] & 0b00110000)
	}
	if btn == start || btn == pressed {
		return true
	}
	return false
}

// Mbtnp returns true if the mouse key was pressed this frame
func (e *Engine) Mbtnp(key MouseBtn) bool {
	btn := btnState(e.RAM[mouseBase] & 0b00000011)
	if key == MiddleClick {
		btn = btnState(e.RAM[mouseBase] & 0b00001100)
	}
	if key == RightClick {
		btn = btnState(e.RAM[mouseBase] & 0b00110000)
	}
	if btn == start {
		return true
	}
	return false
}

// Mbtnr returns true if the mouse key was released this frame
func (e *Engine) Mbtnr(key MouseBtn) bool {
	btn := btnState(e.RAM[mouseBase] & 0b00000011)
	if key == MiddleClick {
		btn = btnState(e.RAM[mouseBase] & 0b00001100)
	}
	if key == RightClick {
		btn = btnState(e.RAM[mouseBase] & 0b00110000)
	}
	if btn == end {
		return true
	}
	return false
}
