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
	}
	opt.Fixed = false //Fixed positions don't work with the map

	for x := 0; x < mw; x++ {
		sprX := (x + dx) * 8
		checkX := sprX - toInt(e.RAM[cameraX:cameraX+2], true)
		if checkX < -8 || checkX > 200 {
			continue
		}
		for y := 0; y < mh; y++ {
			sprY := (y + dy) * 8
			checkY := sprY - toInt(e.RAM[cameraY:cameraY+2], true)
			if checkY < -8 || checkY > 200 {
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

// Mget gets the tile at the x, y coordinate on the map
func (e *Engine) Mget(x, y int) int {
	i := x + y*128
	shift := i % 8
	i = ((i / 8) * 9) + shift
	j := ((i / 8) * 9) + 8
	return int(e.RAM[mapBase-j])<<(shift+1)&0b100000000 | int(e.RAM[mapBase-i])
}

// Mset sets the tile at the x, y coordinate on the map
func (e *Engine) Mset(x, y, t int) {
	i := x + y*128
	shift := i % 8
	i = ((i / 8) * 9) + shift
	j := ((i / 8) * 9) + 8

	e.RAM[mapBase-i] = byte(t)
	e.RAM[mapBase-j] &= (0b00000001 << (7 - shift)) ^ 0b11111111
	e.RAM[mapBase-j] |= byte(t>>1&0b10000000) >> shift
}
