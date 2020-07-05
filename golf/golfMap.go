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
func (e *Engine) Map(mx, my, mw, mh, dx, dy int, opts ...SprOpts) {
	opt := SprOpts{}
	if len(opts) > 0 {
		opt = opts[0]
		if opt.Width == 0 {
			opt.Width = 1
		}
		if opt.Height == 0 {
			opt.Height = 1
		}
		if opt.ScaleH == 0 {
			opt.ScaleH = 1
		}
		if opt.ScaleW == 0 {
			opt.ScaleW = 1
		}
	}

	for x := 0; x < mw; x++ {
		sprX := int(float64((x+dx)*8*opt.Width) * opt.ScaleW)
		if !tileInboundsX(sprX, opt) {
			continue
		}
		for y := 0; y < mh; y++ {
			sprY := int(float64((y+dy)*8*opt.Height) * opt.ScaleW)
			if !tileInboundsY(sprY, opt) {
				continue
			}
			s := e.Mget(x+mx, y+my)
			if s == 0 {
				continue
			}
			e.Spr(s, sprX, sprY, opt)
		}
	}
}

func tileInboundsX(x int, opt SprOpts) bool {
	w := int(float64(8*opt.Width) * opt.ScaleW)
	return tileInbounds(x, 96, w, 0)
}

func tileInboundsY(y int, opt SprOpts) bool {
	h := int(float64(8*opt.Height) * opt.ScaleH)
	return tileInbounds(96, y, 0, h)
}

func tileInbounds(x, y, w, h int) bool {
	if x < -w || x > 192+w {
		return false
	}
	if x < -h || y > 192+h {
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
