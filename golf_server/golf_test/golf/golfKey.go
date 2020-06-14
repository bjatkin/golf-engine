package golf

import (
	"syscall/js"
)

// GoB Keycodes
const (
	AKey      = Key(65)
	BKey      = Key(66)
	CKey      = Key(67)
	DKey      = Key(68)
	EKey      = Key(69)
	FKey      = Key(70)
	GKey      = Key(71)
	HKey      = Key(72)
	IKey      = Key(73)
	JKey      = Key(74)
	KKey      = Key(75)
	LKey      = Key(76)
	MKey      = Key(77)
	NKey      = Key(78)
	OKey      = Key(79)
	PKey      = Key(80)
	QKey      = Key(81)
	RKey      = Key(82)
	SKey      = Key(83)
	TKey      = Key(84)
	UKey      = Key(85)
	VKey      = Key(86)
	WKey      = Key(87)
	XKey      = Key(88)
	YKey      = Key(89)
	ZKey      = Key(90)
	SpaceKey  = Key(32)
	ShiftKey  = Key(16)
	CtrlKey   = Key(17)
	AltKey    = Key(18)
	TabKey    = Key(9)
	ZeroKey   = Key(48)
	OneKey    = Key(49)
	TwoKey    = Key(50)
	ThreeKey  = Key(51)
	FourKey   = Key(52)
	FiveKey   = Key(53)
	SixKey    = Key(54)
	SevenKey  = Key(55)
	EightKey  = Key(56)
	NineKey   = Key(57)
	DelKey    = Key(8)
	CommaKey  = Key(188)
	DotKey    = Key(190)
	FSlashKey = Key(191)
	EnterKey  = Key(13)
	EscKey    = Key(27)
	LeftKey   = Key(37)
	UpKey     = Key(38)
	RightKey  = Key(39)
	DownKey   = Key(40)
)

// Key is a GoB Key
type Key int

type keyListener struct {
	new      []Key
	old      []Key
	released []Key
	ram      *[0xFFFF]byte
}

func newKeyListener(doc js.Value, ram *[0xFFFF]byte) *keyListener {
	ret := keyListener{ram: ram}
	keyDown := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := Key(args[0].Get("keyCode").Int())
		// Add to new if this is a new key
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
	keyUp := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := Key(args[0].Get("keyCode").Int())
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
	doc.Call("addEventListener", "keydown", keyDown)
	doc.Call("addEventListener", "keyup", keyUp)

	return &ret
}

func (kl *keyListener) tick() {
	// Move Keys from new to old
	var add bool
	for _, key := range kl.new {
		add = true
		for _, Okey := range kl.old {
			if key == Okey {
				add = false
				break
			}
		}
		if add {
			kl.old = append(kl.old, key)
		}
	}

	kl.new = []Key{}
	kl.released = []Key{}
}

// Btn returns true if the given key was pressed
func (e *Engine) Btn(key Key) bool {
	if e.Btnp(key) {
		return true
	}
	for _, Okey := range e.kl.old {
		if Okey == key {
			return true
		}
	}
	return false
}

// Btnp returns true if the given key was pressed this frame
func (e *Engine) Btnp(key Key) bool {
	for _, Nkey := range e.kl.new {
		if Nkey == key {
			return true
		}
	}
	return false
}

// Btnr returns true if the given key was released this frame
func (e *Engine) Btnr(key Key) bool {
	for _, Rkey := range e.kl.released {
		if key == Rkey {
			return true
		}
	}
	return false
}
