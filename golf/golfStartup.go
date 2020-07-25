package golf

import "math"

// startupAnim runs the startup animation for golf engine
func (e *Engine) startupAnim() {
	frame := e.Frames()
	startLogo, startText := 0, 80
	endLogo, endIntro := 180, 180
	fadeSpeed := 5
	fadeLen := fadeSpeed*8 - 1
	fadeFrom := []Col{Col0, Col1, Col2, Col3, Col4, Col5, Col6, Col7}
	fadeTo := [][]Col{
		[]Col{Col7, Col7, Col7, Col7, Col7, Col7, Col7, Col7},
		[]Col{Col6, Col7, Col7, Col7, Col7, Col7, Col7, Col7},
		[]Col{Col5, Col6, Col7, Col7, Col7, Col7, Col7, Col7},
		[]Col{Col4, Col5, Col6, Col7, Col7, Col7, Col7, Col7},
		[]Col{Col3, Col4, Col5, Col6, Col7, Col7, Col7, Col7},
		[]Col{Col2, Col3, Col4, Col5, Col6, Col7, Col7, Col7},
		[]Col{Col1, Col2, Col3, Col4, Col5, Col6, Col7, Col7},
		[]Col{Col0, Col1, Col2, Col3, Col4, Col5, Col6, Col7},
		[]Col{Col0, Col0, Col1, Col2, Col3, Col4, Col5, Col6},
		[]Col{Col0, Col0, Col0, Col1, Col2, Col3, Col4, Col5},
		[]Col{Col0, Col0, Col0, Col0, Col1, Col2, Col3, Col4},
		[]Col{Col0, Col0, Col0, Col0, Col0, Col1, Col2, Col3},
		[]Col{Col0, Col0, Col0, Col0, Col0, Col0, Col1, Col2},
		[]Col{Col0, Col0, Col0, Col0, Col0, Col0, Col0, Col1},
		[]Col{Col0, Col0, Col0, Col0, Col0, Col0, Col0, Col0},
	}

	// Clear the BG
	bgCol := Col7
	if frame > endIntro {
		bgCol = fadeTo[(frame-endIntro)/fadeSpeed][0]
	}
	if frame > endIntro+fadeLen {
		bgCol = Col0
	}
	tmpBG := e.BG()
	e.SetBG(bgCol)
	e.Cls()
	e.SetBG(tmpBG)

	// Draw "made with"
	txtf := 0
	if frame > startText {
		txtf = (frame - startText) / fadeSpeed
	}
	if frame > startText+fadeLen {
		txtf = 7
	}
	tCol := TOp{Col: fadeTo[txtf][0], SH: 2, SW: 2}
	e.Text(10, 50, "made with", tCol)

	// Draw golf logo
	sprf := 0
	s := 0.0
	if frame > startLogo {
		s = math.Sin(float64(frame-startLogo)/30) + 1
		sprf = (frame - startLogo) / fadeSpeed
	}
	if frame > startLogo+fadeLen {
		s = 2
		sprf = 7
	}
	if frame > endLogo {
		sprf = (frame-endLogo)/fadeSpeed + 7
	}
	if frame > endLogo+fadeLen {
		sprf = 14
	}
	width := 64.0 * s

	// Change to the internal sprite sheet
	e.RAM[activeSpriteBuff] = internalSpriteBase >> 8
	e.RAM[activeSpriteBuff+1] = internalSpriteBase & 0b0000000011111111

	e.SSpr(152, 0, 64, 24, float64(96-(width/2)), 64.0, SOp{TCol: Col1, SW: s, SH: s, PFrom: fadeFrom, PTo: fadeTo[sprf]})

	// Change back to the main sprite sheet
	e.RAM[activeSpriteBuff] = spriteBase >> 8
	e.RAM[activeSpriteBuff+1] = spriteBase & 0b0000000011111111
}
