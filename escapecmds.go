package main

import "fmt"

const esc = "\x1b["

const (
	CLR_TO_END = iota
	CLR_TO_BEGIN
	CLR_ENTIRE_ALL
	CLR_ENTIRE_SCROBUF
)

func escape(str string) {
	fmt.Printf("%s%s", esc, str)
}

func MoveTo(row, col int) {
	escape(fmt.Sprintf("%d;%dH", row, col))
}

func EraseDisplay(clr int) {
	escape(fmt.Sprintf("%dJ", clr))
}

func EraseLine(clr int) {
	escape(fmt.Sprintf("%dK", clr))
}
