package golf

// LoadMap loads the sprite sheet into memory
func (e *Engine) LoadMap(mapData [0x4800]byte) {
	for i, b := range mapData {
		e.RAM[mapBase-i] = b
	}
}

// Map draws the map on the screen starting from tile
// mx, my with a size of mw and mh. The map is draw
// at screen coordinate dx, dy
func (e *Engine) Map(mx, my, mw, mh int, dx, dy float64, opts ...SOp) {
	opt := SOp{}
	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt.W == 0 {
		opt.W = 1
	}
	if opt.H == 0 {
		opt.H = 1
	}
	if opt.SH == 0 {
		opt.SH = 1
	}
	if opt.SW == 0 {
		opt.SW = 1
	}
	cx := toInt(e.RAM[cameraX:cameraX+2], true)
	cy := toInt(e.RAM[cameraY:cameraY+2], true)
	if opt.Fixed {
		cx, cy = 0, 0
	}

	for x := 0; x < mw; x++ {
		sprX := int(float64((x+int(dx))*8*opt.W) * roundPxl(opt.SW, float64(8*opt.W)))
		if !tileInboundsX(sprX-cx, opt) {
			continue
		}
		for y := 0; y < mh; y++ {
			sprY := int(float64((y+int(dy))*8*opt.H) * roundPxl(opt.SH, float64(8*opt.H)))
			if !tileInboundsY(sprY-cy, opt) {
				continue
			}
			s := e.Mget(x+mx, y+my)
			if s == 0 {
				continue
			}
			e.Spr(s, float64(sprX), float64(sprY), opt)
		}
	}
}

// roundPxl rounds to the nearist pixel rather than the nearist number
// number is the number to be rounded, size is the number of pixels
func roundPxl(number, size float64) float64 {
	inc := 1.0 / size
	num := int(number)
	frac := number - float64(num)
	for i := 0.0; i < 8.0; i += inc {
		if i > frac {
			return float64(num) + (i - inc)
		}
	}
	return float64(num)
}

// tileInboundsX checks if the x coordiante is in screen bounds
// sprite opts are taken into consideration
func tileInboundsX(x int, opt SOp) bool {
	w := int(float64(8*opt.W) * opt.SW)
	return tileInbounds(x, 96, w, 0)
}

// tileInboundsX checks if the y coordiante is in screen bounds
// sprite opts are taken into consideration
func tileInboundsY(y int, opt SOp) bool {
	h := int(float64(8*opt.H) * opt.SH)
	return tileInbounds(96, y, 0, h)
}

// tileInbounds checks if the x and y coordinates are in screen bounds
// w and h ensure the tiles that are partially off screen are still drawn
func tileInbounds(x, y, w, h int) bool {
	if x < -w || x > 192+w {
		return false
	}
	if y < -h || y > 192+h {
		return false
	}
	return true
}

// Mget gets the tile at the x, y coordinate on the map
func (e *Engine) Mget(x, y int) int {
	dex := x + y*128
	shift := dex % 8
	i := ((dex / 8) * 9) + shift
	j := ((dex / 8) * 9) + 8
	return int(e.RAM[mapBase-j])<<(shift+1)&0b100000000 | int(e.RAM[mapBase-i])
}

// Mset sets the tile at the x, y coordinate on the map
func (e *Engine) Mset(x, y, t int) {
	dex := x + y*128
	shift := dex % 8
	i := ((dex / 8) * 9) + shift
	j := ((i / 8) * 9) + 8

	e.RAM[mapBase-i] = byte(t)
	e.RAM[mapBase-j] &= (0b00000001 << (7 - shift)) ^ 0b11111111
	e.RAM[mapBase-j] |= byte(t>>1&0b10000000) >> shift
}
