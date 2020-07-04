package main

import (
	gob "fantasyConsole/terminal/GoB"
	"fmt"
	"strconv"
	"strings"
)

var lines = []string{""}
var inputLines = []bool{true}
var aLine, histLine = 0, 0
var history = []string{}

func terminalUpdate(tdiff float64) {
	engine.DrawMouse = false
	lineTxt := lines[aLine]
	switch {
	case engine.Btnp(gob.AKey):
		lineTxt += "A"
	case engine.Btnp(gob.BKey):
		lineTxt += "B"
	case engine.Btnp(gob.CKey):
		lineTxt += "C"
	case engine.Btnp(gob.DKey):
		lineTxt += "D"
	case engine.Btnp(gob.EKey):
		lineTxt += "E"
	case engine.Btnp(gob.FKey):
		lineTxt += "F"
	case engine.Btnp(gob.GKey):
		lineTxt += "G"
	case engine.Btnp(gob.HKey):
		lineTxt += "H"
	case engine.Btnp(gob.IKey):
		lineTxt += "I"
	case engine.Btnp(gob.JKey):
		lineTxt += "J"
	case engine.Btnp(gob.KKey):
		lineTxt += "K"
	case engine.Btnp(gob.LKey):
		lineTxt += "L"
	case engine.Btnp(gob.MKey):
		lineTxt += "M"
	case engine.Btnp(gob.NKey):
		lineTxt += "N"
	case engine.Btnp(gob.OKey):
		lineTxt += "O"
	case engine.Btnp(gob.PKey):
		lineTxt += "P"
	case engine.Btnp(gob.QKey):
		lineTxt += "Q"
	case engine.Btnp(gob.RKey):
		lineTxt += "R"
	case engine.Btnp(gob.SKey):
		lineTxt += "S"
	case engine.Btnp(gob.TKey):
		lineTxt += "T"
	case engine.Btnp(gob.UKey):
		lineTxt += "U"
	case engine.Btnp(gob.VKey):
		lineTxt += "V"
	case engine.Btnp(gob.WKey):
		lineTxt += "W"
	case engine.Btnp(gob.XKey):
		lineTxt += "X"
	case engine.Btnp(gob.YKey):
		lineTxt += "Y"
	case engine.Btnp(gob.ZKey):
		lineTxt += "Z"
	case engine.Btnp(gob.ZeroKey):
		lineTxt += "0"
	case engine.Btnp(gob.OneKey):
		lineTxt += "1"
	case engine.Btnp(gob.TwoKey):
		lineTxt += "2"
	case engine.Btnp(gob.ThreeKey):
		lineTxt += "3"
	case engine.Btn(gob.ShiftKey) && engine.Btnp(gob.FourKey):
		lineTxt += "$"
	case engine.Btnp(gob.FourKey):
		lineTxt += "4"
	case engine.Btn(gob.ShiftKey) && engine.Btnp(gob.FiveKey):
		lineTxt += "%"
	case engine.Btnp(gob.FiveKey):
		lineTxt += "5"
	case engine.Btnp(gob.SixKey):
		lineTxt += "6"
	case engine.Btnp(gob.SevenKey):
		lineTxt += "7"
	case engine.Btnp(gob.EightKey):
		lineTxt += "8"
	case engine.Btnp(gob.NineKey):
		lineTxt += "9"
	case engine.Btnp(gob.SpaceKey):
		lineTxt += " "
	case engine.Btnp(gob.DelKey):
		if len(lineTxt) == 0 {
			break
		}
		lineTxt = lineTxt[:len(lineTxt)-1]
	case engine.Btnp(gob.CommaKey):
		lineTxt += ","
	case engine.Btnp(gob.DotKey):
		lineTxt += "."
	case engine.Btn(gob.ShiftKey) && engine.Btnp(gob.FSlashKey):
		lineTxt += "?"
	case engine.Btnp(gob.FSlashKey):
		lineTxt += "/"
	case engine.Btnp(gob.EnterKey):
		lines = append(lines, "")
		inputLines = append(inputLines, true)
		aLine++
		histLine = 0
		runCommand(lineTxt)
		return
	case engine.Btnp(gob.UpKey):
		if len(history) == 0 {
			break
		}
		histLine++
		l := len(history)
		if histLine > l {
			histLine--
		}
		lineTxt = history[l-histLine]
	case engine.Btnp(gob.DownKey):
		histLine--
		if histLine <= 0 {
			histLine++
			lineTxt = ""
		} else {
			lineTxt = history[len(history)-histLine]
		}
	}

	lines[aLine] = lineTxt
}

func terminalDraw() {
	engine.Cls()

	frames++
	frames %= 60

	for j, line := range lines {
		if len(lines) > 27 {
			offset := len(lines) - 27
			if j < offset {
				continue
			}
			j -= offset
		}
		x := 6
		if inputLines[j] {
			engine.Text("%", 5, 3+10*j, gob.Col12)
			x = 16
		}
		if j == len(lines)-1 && frames < 35 {
			engine.Spr8(0, x+len(line)*6, 3+j*10, gob.SprOpts{Transparent: gob.TCol0})
		}
		engine.Text(line, x, 3+j*10)
	}

	drawFPS()
}

func runCommand(command string) {
	if command != "" {
		history = append(history, command)
	}
	if len(history) > 500 {
		history = history[:499]
	}
	if command == "" {
		return
	}
	com := strings.Split(command, " ")
	key := strings.ToLower(com[0])
	args := []string{}
	if len(com) > 1 {
		args = com[1:]
	}

	ret := ""
	switch key {
	case "pala":
		ret = swapPal(0, args)
	case "palb":
		ret = swapPal(1, args)
	case "clear":
		clearScreen()
	case "palget":
		p1, p2 := engine.PalGet()
		ret = fmt.Sprintf("PALA: %d, PALB: %d", p1, p2)
	case "editor":
		engine.Update = editorUpdate
		engine.Draw = editorDraw
	case "palr":
		engine.Pal1(gob.Pal0)
		engine.Pal2(gob.Pal1)
		ret = "PALA: 0, PALB: 1"
	case "fps":
		showFPS = !showFPS
	}

	if ret != "" {
		lines[aLine] = ret
		inputLines[aLine] = false
		lines = append(lines, "")
		inputLines = append(inputLines, true)
		aLine++
	}
}

func swapPal(pal int, args []string) string {
	errMsg := "USAGE: PALB [0-7]"
	if pal == 0 {
		errMsg = "USAGE: PALA [0-7]"
	}
	if len(args) == 0 {
		return errMsg
	}
	newPal, err := strconv.Atoi(args[0])
	if err == nil && newPal < 8 && newPal >= 0 {
		if pal == 0 {
			engine.Pal1(gob.Pal(newPal))
		} else {
			engine.Pal2(gob.Pal(newPal))
		}
		return ""
	}
	return errMsg
}

func clearScreen() {
	lines = []string{""}
	inputLines = []bool{true}
	aLine = 0
}
