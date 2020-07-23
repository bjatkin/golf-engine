package golf

import "math"

// startupAnim runs the startup animation for golf engine
func (e *Engine) startupAnim() {
	bgCol := Col3
	pto := []Col{Col0, Col1, Col2, Col3}
	pfrom := []Col{Col0, Col1, Col2, Col3}
	if e.Frames() > 180 {
		bgCol = Col2
		pto = []Col{Col0, Col0, Col1, Col2}
	}
	if e.Frames() > 190 {
		bgCol = Col1
		pto = []Col{Col0, Col0, Col0, Col1}
	}
	if e.Frames() > 200 {
		bgCol = Col0
		pto = []Col{Col0, Col0, Col0, Col0}
	}
	tmpBG := e.BG()
	e.SetBG(bgCol)
	e.Cls()
	e.SetBG(tmpBG)

	// Draw the logo
	tCol := TOp{Col: Col3, SH: 2, SW: 2}
	if e.Frames() > 40 {
		tCol = TOp{Col: Col2, SH: 1.8, SW: 1.8}
	}
	if e.Frames() > 50 {
		tCol = TOp{Col: Col1, SH: 1.9, SW: 1.9}
	}
	if e.Frames() > 60 {
		tCol = TOp{Col: Col0, SH: 2, SW: 2}
	}
	e.Text(10, 50, "made with", tCol)

	// Change to the internal sprite sheet
	e.RAM[activeSpriteBuff] = internalSpriteBase >> 8
	e.RAM[activeSpriteBuff+1] = internalSpriteBase & 0b0000000011111111

	s := math.Sin(float64(e.Frames())/30) + 1
	if e.Frames() > 30 {
		s = 2
	}
	width := 64.0 * s
	e.SSpr(152, 0, 64, 24, 96-(width/2), 64, SOp{TCol: Col7, SW: s, SH: s, PTo: pto, PFrom: pfrom})

	// Change back to the main sprite sheet
	e.RAM[activeSpriteBuff] = spriteBase >> 8
	e.RAM[activeSpriteBuff+1] = spriteBase & 0b0000000011111111
}
