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

//TODO replace with mouseBase
const mouseBase = uint16(0x3610)

func newMouseListener(canvas js.Value, ram *[0xFFFF]byte) *mouseListener {
	ret := mouseListener{ram: ram}
	mouseMove := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		height := canvas.Get("offsetHeight").Int()
		ratio := float64(ScreenHeight) / float64(height)
		ram[0x360E] = byte(args[0].Get("offsetX").Float() * ratio)
		ram[0x360F] = byte(args[0].Get("offsetY").Float() * ratio)

		return nil
	})
	mouseDown := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := MouseBtn(args[0].Get("button").Int())

		if e == LeftClick {
			btn := btnState(ret.ram[mouseBase] & 0b00000011)
			if btn == unpressed {
				ret.ram[mouseBase] |= byte(start)
			}
		}
		if e == MiddleClick {
			btn := btnState((ret.ram[mouseBase] & 0b00001100) >> 2)
			if btn == unpressed {
				ret.ram[mouseBase] |= (byte(start) << 2)
			}
		}
		if e == RightClick {
			btn := btnState((ret.ram[mouseBase] & 0b00110000) >> 4)
			if btn == unpressed {
				ret.ram[mouseBase] |= (byte(start) << 4)
			}
		}

		return nil
	})
	mouseUp := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := MouseBtn(args[0].Get("button").Int())

		if e == LeftClick {
			btn := btnState(ret.ram[mouseBase] & 0b00000011)
			if btn == pressed || btn == start {
				ret.ram[mouseBase] &= 0b11111100
				ret.ram[mouseBase] |= byte(end)
			}
		}
		if e == MiddleClick {
			btn := btnState((ret.ram[mouseBase] & 0b00001100) >> 2)
			if btn == pressed || btn == start {
				ret.ram[mouseBase] &= 0b11110011
				ret.ram[mouseBase] |= (byte(end) << 2)
			}
		}
		if e == RightClick {
			btn := btnState((ret.ram[mouseBase] & 0b00110000) >> 4)
			if btn == pressed || btn == start {
				ret.ram[mouseBase] &= 0b11001111
				ret.ram[mouseBase] |= (byte(end) << 4)
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
	// Move btn from start to pressed
	if btnState(ml.ram[mouseBase]&0b00000011) == start {
		ml.ram[mouseBase] &= 0b11111100
		ml.ram[mouseBase] |= byte(pressed)
	}
	if btnState((ml.ram[mouseBase]&0b00001100)>>2) == start {
		ml.ram[mouseBase] &= 0b11110011
		ml.ram[mouseBase] |= (byte(pressed) << 2)
	}
	if btnState((ml.ram[mouseBase]&0b00110000)>>4) == start {
		ml.ram[mouseBase] &= 0b11001111
		ml.ram[mouseBase] |= (byte(pressed) << 4)
	}

	// Move from end to unpressed
	if btnState(ml.ram[mouseBase]&0b00000011) == end {
		ml.ram[mouseBase] &= 0b11111100
	}
	if btnState((ml.ram[mouseBase]&0b00001100)>>2) == end {
		ml.ram[mouseBase] &= 0b11110011
	}
	if btnState((ml.ram[mouseBase]&0b00110000)>>4) == end {
		ml.ram[mouseBase] &= 0b11001111
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
