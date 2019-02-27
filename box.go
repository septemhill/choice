package main

import (
	"bytes"
	"fmt"
	"strings"
)

const DEFAULT_BORDER = "- │ ┌ ┐ └ ┘"

func MoveToPaint(x, y int, str string) {
	escape(fmt.Sprintf("%d;%dH%s", x, y, str))
}

type Box struct {
	Buffer *bytes.Buffer
	X      int
	Y      int
	Width  int
	Height int
}

func (b Box) Write(p []byte) (n int, err error) {
	return b.Buffer.Write(p)
}

func (b *Box) drawline(row, x, width int, data string) {
	runes := strings.Split(data, "")

	for i := 0; i < width; i++ {
		MoveToPaint(row, x+i, runes[i])
	}
}

func (b *Box) clearArea() {
	for i := b.Y; i <= b.Height+b.Y; i++ {
		for j := b.X; j < b.Width+b.X; j++ {
			MoveToPaint(i, j, " ")
		}
	}
}

func (b *Box) Draw() {
	var line string
	borders := strings.Split(DEFAULT_BORDER, " ")
	linecnt := 0
	lines := bytes.Split(b.Buffer.Bytes(), []byte("\n"))

	b.clearArea()
	for i := b.Y; i < b.Height+b.Y; i++ {
		if i == b.Y {
			line += borders[2] + strings.Repeat(borders[0], b.Width-2) + borders[3]
		} else if i == b.Height+b.Y-1 {
			line += borders[4] + strings.Repeat(borders[0], b.Width-2) + borders[5]
		} else if linecnt < len(lines) {
			strs := strings.Split(string(lines[linecnt]), "")
			if len(strs) > b.Width-2 {
				line += borders[1] + strings.Join(strs[:b.Width-2], "") + borders[1]
			} else {
				line += borders[1] + string(lines[linecnt]) + strings.Repeat(" ", b.Width-len(strs)-2) + borders[1]
			}
			linecnt++
		} else {
			line += borders[1] + strings.Repeat(" ", b.Width-2) + borders[1]
		}

		b.drawline(i, b.X, b.Width, line)
		line = ""
	}
}
