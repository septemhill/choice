package main

import (
	"regexp"
	"strconv"
)

type StringWithEscape struct {
	oriStr   []rune
	escStart string
	escEnd   string
	oriStart int
	oriEnd   int
}

func (s StringWithEscape) String() string {
	//return s.escStart + string(s.oriStr) + s.escEnd
	return "oriStart: " + strconv.FormatInt(int64(s.oriStart), 10) + ", oriEnd: " + strconv.FormatInt(int64(s.oriEnd), 10)
}

type EscapeString []StringWithEscape

func (e EscapeString) Len() int {
	var strLen int
	for i := 0; i < len(e); i++ {
		strLen += len(e[i].oriStr)
	}

	return strLen
}

func (e EscapeString) String() string {
	var str string
	for i := 0; i < len(e); i++ {
		str += e[i].escStart + string(e[i].oriStr) + e[i].escEnd
	}
	return str
}

func (e EscapeString) Element(ie int) string {
	for i := 0; i < len(e); i++ {
		if ie >= e[i].oriStart && ie < e[i].oriEnd {
			idx := ie - e[i].oriStart
			return e[i].escStart + string(e[i].oriStr[idx]) + e[i].escEnd
		}
	}

	return ""
}

func (e EscapeString) Substring(start, end int) string {
	var rstr string
	var es, ee int
	var esoff, eeoff int

	for i := 0; i < len(e); i++ {
		if start >= e[i].oriStart && start < e[i].oriEnd {
			es = i
			esoff = start - e[i].oriStart
		}
		if end >= e[i].oriStart && end <= e[i].oriEnd {
			ee = i
			eeoff = end - e[i].oriStart
		}
	}

	if es == ee {
		rstr += e[es].escStart + string(e[es].oriStr[esoff:eeoff]) + e[es].escEnd
	} else {
		for i := es; i <= ee; i++ {
			if i == es {
				rstr += e[i].escStart + string(e[i].oriStr[esoff:]) + e[i].escEnd
			} else if i == ee {
				rstr += e[i].escStart + string(e[i].oriStr[:eeoff]) + e[i].escEnd
			} else {
				rstr += e[i].escStart + string(e[i].oriStr) + e[i].escEnd
			}
		}
	}

	return rstr
}

func stringParse(str string) EscapeString {
	cont := make([]StringWithEscape, 0)
	re := regexp.MustCompile("\x1b[[0-9;]*m")
	idxs := re.FindAllStringIndex(str, -1)
	strIdx := 0

	if len(idxs) == 0 && len(str) != 0 {
		var cstr StringWithEscape
		cstr.oriStr = []rune(str[0:])
		cstr.oriStart = strIdx
		cstr.oriEnd = strIdx + len(cstr.oriStr)

		cont = append(cont, cstr)
	}

	for i := 0; i < len(idxs); i++ {
		var cstr StringWithEscape
		if i == 0 && idxs[i][0] > 0 {
			cstr.oriStr = []rune(str[0:idxs[0][0]])
			cstr.oriStart = strIdx
			cstr.oriEnd = strIdx + len(cstr.oriStr)

			strIdx += len(cstr.oriStr)
			cont = append(cont, cstr)

			cstr.oriStr = []rune(str[idxs[i][1]:idxs[i+1][0]])
			cstr.escStart = str[idxs[i][0]:idxs[i][1]]
			cstr.escEnd = str[idxs[i+1][0]:idxs[i+1][1]]
			cstr.oriStart = strIdx
			cstr.oriEnd = strIdx + len([]rune(str[idxs[i][1]:idxs[i+1][0]]))

			strIdx += len(cstr.oriStr)
			cont = append(cont, cstr)

		} else if i == len(idxs)-1 {
			cstr.oriStr = []rune(str[idxs[len(idxs)-1][1]:len(str)])
			cstr.oriStart = strIdx
			cstr.oriEnd = strIdx + len(cstr.oriStr)

			strIdx += len(cstr.oriStr)
			cont = append(cont, cstr)
		} else {
			cstr.oriStr = []rune(str[idxs[i][1]:idxs[i+1][0]])
			cstr.escStart = str[idxs[i][0]:idxs[i][1]]
			cstr.escEnd = str[idxs[i+1][0]:idxs[i+1][1]]
			cstr.oriStart = strIdx
			cstr.oriEnd = strIdx + len([]rune(str[idxs[i][1]:idxs[i+1][0]]))
			strIdx += len(cstr.oriStr)
			cont = append(cont, cstr)
		}
	}

	for i := 0; i < len(cont); i++ {
		if len(cont[i].oriStr) == 0 {
			cont = append(cont[:i], cont[i+1:]...)
		}
		//i--
	}

	return cont
}
