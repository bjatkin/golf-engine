package golf

import (
	"strings"
)

// TOp additional options for drawing text
type TOp struct {
	Col    Col
	Fixed  bool
	SW, SH float64
}

// the TextLine for TextL and TextR
var textLline, textRline = 0, 0

// TextL prints text at the top left of the screen
// the cursor moves to a new line each time TextL is called
func (e *Engine) TextL(text string, opts ...TOp) {
	splitText := strings.Split(text, "\n")
	for _, line := range splitText {
		if len(opts) > 0 {
			e.Text(1, float64(1+6*textLline), line, opts[0])
		} else {
			e.Text(1, float64(1+6*textLline), line)
		}
		textLline++
	}
}

// TextR prints text at the top right of the screen
// the cursor moves to a new line each time TextR is called
func (e *Engine) TextR(text string, opts ...TOp) {
	splitText := strings.Split(text, "\n")
	for _, line := range splitText {
		x := ScreenWidth - 1 - len(line)*6
		if len(opts) > 0 {
			e.Text(float64(x), float64(1+6*textRline), line, opts[0])
		} else {
			e.Text(float64(x), float64(1+6*textRline), line)
		}
		textRline++
	}
}

// the Text font reference
const textRef = "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_+={}[]\\:;\"<>,./?|"
const btnRef = "(<)(>)(^)(v)(x)(o)(l)(r)(+)(-)"
const specialRef = ":):(x(:|=[|^|v<-->$$@@<|<3<4+1-1~~()[]:;**"

// Text prints text at the x, y coords on the screen
func (e *Engine) Text(x, y float64, text string, opts ...TOp) {
	text = strings.ToLower(text)
	px, py := x, y
	opt := TOp{}
	sopt := SOp{TCol: Col1, SH: 1, SW: 1}
	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt.Col != 0 {
		sopt.PFrom = []Col{Col0}
		sopt.PTo = []Col{opt.Col}
	}
	if opt.SH != 0 {
		sopt.SH = opt.SH
	}
	if opt.SW != 0 {
		sopt.SW = opt.SW
	}
	sopt.Fixed = opt.Fixed
	width := 6 * sopt.SW
	height := 6 * sopt.SH

	e.setActiveSpriteBuff(internalSpriteBase)

	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			px = x
			py += height //6
			continue
		}
		if text[i] == ' ' {
			px += width //6
			continue
		}
		if text[i] == '^' {
			i++
			dex := strings.Index(textRef, string(text[i]))
			e.drawChar(px, py, dex, sopt)
			px += width //6
			continue
		}
		bdex := -1
		if i+2 < len(text) {
			bdex = strings.Index(btnRef, string(text[i:i+3]))
		}
		if bdex%3 == 0 {
			bdex = bdex/3 + 86
			e.drawChar(px, py, bdex, sopt)
			px += width //6
			i += 2
			continue
		}
		sdex := -1
		if i+1 < len(text) {
			sdex = strings.Index(specialRef, string(text[i:i+2]))
		}
		if sdex%2 == 0 {
			sdex = sdex/2 + 65
			e.drawChar(px, py, sdex, sopt)
			px += width //6
			i++
			continue
		}
		dex := -1
		if i < len(text) {
			dex = strings.Index(textRef, string(text[i]))
		}
		e.drawChar(px, py, dex, sopt)
		px += width //6
	}

	e.setActiveSpriteBuff(spriteBase)
}

func (e *Engine) drawChar(x, y float64, i int, opt SOp) {
	if i < 0 {
		return
	}
	sx := i % 24
	sy := i / 24
	e.SSpr(sx*6, sy*6, 6, 6, x, y, opt)
}
