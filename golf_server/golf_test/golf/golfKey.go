package golf

import (
	"syscall/js"
)

// Golf Key Codes, Mirrors JS key code values
const (
	Backspace    = Key(8)
	Tab          = Key(9)
	Enter        = Key(13)
	Shift        = Key(16)
	Ctrl         = Key(17)
	Alt          = Key(18)
	Break        = Key(19)
	CapsLock     = Key(20)
	Esc          = Key(27)
	Space        = Key(32)
	PageUp       = Key(33)
	PageDown     = Key(34)
	End          = Key(35)
	Home         = Key(36)
	LeftArrow    = Key(37)
	UpArrow      = Key(38)
	RightArrow   = Key(39)
	DownArrow    = Key(40)
	Insert       = Key(45)
	Delete       = Key(46)
	ZeroKey      = Key(48)
	OneKey       = Key(49)
	TwoKey       = Key(50)
	ThreeKey     = Key(51)
	FourKey      = Key(52)
	FiveKey      = Key(53)
	SixKey       = Key(54)
	SevenKey     = Key(55)
	EightKey     = Key(56)
	NineKey      = Key(57)
	AKey         = Key(65)
	BKey         = Key(66)
	CKey         = Key(67)
	DKey         = Key(68)
	EKey         = Key(69)
	FKey         = Key(70)
	GKey         = Key(71)
	HKey         = Key(72)
	IKey         = Key(73)
	JKey         = Key(74)
	KKey         = Key(75)
	LKey         = Key(76)
	MKey         = Key(77)
	NKey         = Key(78)
	OKey         = Key(79)
	PKey         = Key(80)
	QKey         = Key(81)
	RKey         = Key(82)
	SKey         = Key(83)
	TKey         = Key(84)
	UKey         = Key(85)
	VKey         = Key(86)
	WKey         = Key(87)
	XKey         = Key(88)
	YKey         = Key(89)
	ZKey         = Key(90)
	LeftWinKey   = Key(91)
	RightWinKey  = Key(92)
	Select       = Key(93)
	NumPad0      = Key(96)
	NumPad1      = Key(97)
	NumPad2      = Key(98)
	NumPad3      = Key(99)
	NumPad4      = Key(100)
	NumPad5      = Key(101)
	NumPad6      = Key(102)
	NumPad7      = Key(103)
	NumPad8      = Key(104)
	NumPad9      = Key(105)
	NumPadMul    = Key(106)
	NumPadPlus   = Key(107)
	NumPadMinus  = Key(109)
	NumPadDot    = Key(110)
	NumPadDiv    = Key(111)
	F1           = Key(112)
	F2           = Key(113)
	F3           = Key(114)
	F4           = Key(115)
	F5           = Key(116)
	F6           = Key(117)
	F7           = Key(118)
	F8           = Key(119)
	F9           = Key(120)
	F10          = Key(121)
	F11          = Key(122)
	F12          = Key(123)
	NumLock      = Key(144)
	ScrollLock   = Key(145)
	SemiColon    = Key(186)
	Equals       = Key(187)
	Comma        = Key(188)
	Minus        = Key(189)
	Period       = Key(190)
	FSlash       = Key(191)
	Tilda        = Key(192)
	OpenBracket  = Key(219)
	BSlash       = Key(220)
	CloseBracket = Key(221)
	Quotes       = Key(222)
)

// Key is a Golf Key
type Key int

// btnState is the current state of the button press
type btnState byte

// KeyBtn and MouseBtn key states
const (
	unpressed = btnState(0)
	start     = btnState(1)
	end       = btnState(2)
	pressed   = btnState(3)
)

func (e *Engine) initKeyListener(doc js.Value) {
	keyDown := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Get KeyCode starting at 0
		k := Key(args[0].Get("keyCode").Int()) - Backspace
		masks := []byte{
			0b00000011,
			0b00001100,
			0b00110000,
			0b11000000,
		}

		addr := keyBase + uint16(k/4)
		b := e.RAM[addr]
		m := masks[k%4]
		shift := (k % 4) * 2
		btn := btnState(b & m)
		if btn == unpressed {
			e.RAM[addr] |= (byte(start) << shift)
		}

		return nil
	})
	keyUp := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		k := Key(args[0].Get("keyCode").Int()) - Backspace
		masks := []byte{
			0b00000011,
			0b00001100,
			0b00110000,
			0b11000000,
		}

		addr := keyBase + uint16(k/4)
		b := e.RAM[addr]
		m := masks[k%4]
		shift := (k % 4) * 2
		btn := btnState(b & m >> shift)
		if btn == pressed || btn == start {
			e.RAM[addr] &= (m ^ 0b11111111)
			e.RAM[addr] |= (byte(end) << shift)
		}

		return nil
	})
	doc.Call("addEventListener", "keydown", keyDown)
	doc.Call("addEventListener", "keyup", keyUp)
}

func (e *Engine) tickKeyboard() {
	masks := []byte{
		0b00000011,
		0b00001100,
		0b00110000,
		0b11000000,
	}
	for i := Backspace; i < Quotes; i++ {
		k := byte(i - Backspace)
		addr := keyBase + uint16(k/4)
		b := e.RAM[addr]
		m := masks[k%4]
		shift := (k % 4) * 2
		btn := btnState(b & m >> shift)
		if btn == start {
			e.RAM[addr] &= (m ^ 0b11111111)
			e.RAM[addr] |= (byte(pressed) << shift)
		}

		if btn == end {
			e.RAM[addr] &= (m ^ 0b11111111)
			e.RAM[addr] |= (byte(unpressed) << shift)
		}
	}
}

// Btn returns true if the given key was pressed
func (e *Engine) Btn(key Key) bool {
	if e.Btnp(key) {
		return true
	}

	masks := []byte{
		0b00000011,
		0b00001100,
		0b00110000,
		0b11000000,
	}

	k := key - Backspace
	addr := keyBase + uint16(k/4)
	b := e.RAM[addr]
	m := masks[k%4]
	shift := (k % 4) * 2
	if btnState(b&m>>shift) == pressed {
		return true
	}

	return false
}

// Btnp returns true if the given key was pressed this frame
func (e *Engine) Btnp(key Key) bool {
	masks := []byte{
		0b00000011,
		0b00001100,
		0b00110000,
		0b11000000,
	}

	k := key - Backspace
	addr := keyBase + uint16(k/4)
	b := e.RAM[addr]
	m := masks[k%4]
	shift := (k % 4) * 2
	if btnState(b&m>>shift) == start {
		return true
	}

	return false
}

// Btnr returns true if the given key was released this frame
func (e *Engine) Btnr(key Key) bool {
	masks := []byte{
		0b00000011,
		0b00001100,
		0b00110000,
		0b11000000,
	}

	k := key - Backspace
	addr := keyBase + uint16(k/4)
	b := e.RAM[addr]
	m := masks[k%4]
	shift := (k % 4) * 2
	if btnState(b&m>>shift) == end {
		return true
	}

	return false
}
