package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/septemhill/fion"
)

type coordinateMap map[int]map[int]struct{}

type Coordinate struct {
	X     int
	Y     int
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

type Map struct {
	cords  []*Coordinate
	Width  int
	Height int
}

func (m *Map) exist(x, y int, cords []*Coordinate) bool {
	for i := 0; i < len(cords); i++ {
		if x == cords[i].X && y == cords[i].Y {
			return true
		}
	}

	return false
}

func (m *Map) find(x, y int, cords []*Coordinate) *Coordinate {
	for i := 0; i < len(cords); i++ {
		if x == cords[i].X && y == cords[i].Y {
			return cords[i]
		}
	}

	return nil
}

func (m *Map) drawMapDashline(width, height, space int, coord []*Coordinate) {
	dash := "-"

	for i := 0; i < width; i++ {
		if m.exist(i, height-1, coord) && m.exist(i, height, coord) {
			if m.exist(i, height, coord) && m.exist(i+1, height, coord) {
				dash += fion.BRed(strings.Repeat(" ", space+1))

			} else {
				dash += fion.BRed(strings.Repeat(" ", space))
				dash += "-"
			}
		} else {
			dash += strings.Repeat("-", space+1)
		}
		//dash += strings.Repeat("-", space+1)
	}
	fmt.Println(dash)
}

func (m *Map) drawMapGridColumn(width, height, space int, coord []*Coordinate) {
	grid := "|"

	for i := 0; i < width; i++ {
		if m.exist(i, height, coord) {
			grid += fion.BRed(strings.Repeat(" ", space))

			if m.exist(i+1, height, coord) {
				grid += fion.BRed(" ")
			} else {
				grid += "|"
			}
		} else {
			grid += strings.Repeat(" ", space) + "|"
		}
	}
	fmt.Println(grid)
}

func (m *Map) DrawMap() {
	dashLineCount := m.Height + 1
	space := 3
	h := 0

	for i := 0; i < dashLineCount+m.Height; i++ {
		if i%2 == 0 {
			m.drawMapDashline(m.Width, h, space, m.cords)
		} else {
			m.drawMapGridColumn(m.Width, h, space, m.cords)
			h++
		}
	}
}

func (m *Map) drawPathDashline(width, height, space int, cords []*Coordinate) {
	dash := "-"

	for i := 0; i < width; i++ {
		if m.exist(i, height-1, cords) && m.exist(i, height, cords) {
			up, down := m.find(i, height-1, cords), m.find(i, height, cords)
			if up.Down || down.Up {
				dash += fion.BRed(strings.Repeat(" ", space)) + "-"
			} else {
				dash += strings.Repeat("-", space+1)
			}
		} else {
			dash += strings.Repeat("-", space+1)
		}
	}
	fmt.Println(dash)
}

func (m *Map) drawPathGridColumn(width, height, space int, cords []*Coordinate) {
	grid := "|"

	for i := 0; i < width; i++ {
		if m.exist(i, height, cords) && m.exist(i+1, height, cords) {
			left, right := m.find(i, height, cords), m.find(i+1, height, cords)
			if left.Right || right.Left {
				grid += fion.BRed(strings.Repeat(" ", space+1))
			} else {
				grid += fion.BRed(strings.Repeat(" ", space)) + "|"
			}
		} else if m.exist(i, height, cords) {
			grid += fion.BRed(strings.Repeat(" ", space)) + "|"
		} else {
			grid += strings.Repeat(" ", space) + "|"
		}
	}
	fmt.Println(grid)
}

func (m *Map) DrawPath() {
	dashLineCount := m.Height + 1
	space := 3
	h := 0

	for i := 0; i < dashLineCount+m.Height; i++ {
		if i%2 == 0 {
			m.drawPathDashline(m.Width, h, space, m.cords)
		} else {
			m.drawPathGridColumn(m.Width, h, space, m.cords)
			h++
		}
	}
}

func mapSize(m coordinateMap) int {
	c := 0
	for i := 0; i < len(m); i++ {
		c += len(m[i])
	}

	return c
}

func entryWay(coord *[]*Coordinate) {
	for i := 0; i < len(*coord)-1; i++ {
		if (*coord)[i].X-(*coord)[i+1].X > 0 {
			(*coord)[i].Left = true
		}
		if (*coord)[i].X-(*coord)[i+1].X < 0 {
			(*coord)[i].Right = true
		}
		if (*coord)[i].Y-(*coord)[i+1].Y > 0 {
			(*coord)[i].Up = true
		}
		if (*coord)[i].Y-(*coord)[i+1].Y < 0 {
			(*coord)[i].Down = true
		}
	}
}

//func CreateMap(width, height int) []*Coordinate {
func CreateMap(width, height int) *Map {
	inMap, outMap := make(coordinateMap), make(coordinateMap)
	trace := make([]*Coordinate, 0)
	size := width * height * 40 / 100
	grids := 0

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < width; i++ {
		inMap[i], outMap[i] = make(map[int]struct{}), make(map[int]struct{})
		for j := 0; j < height; j++ {
			outMap[i][j] = struct{}{}
		}
	}

	startX, startY := 0, 0
	delete(outMap[0], 0)
	inMap[0][0] = struct{}{}

	for {
		ccont := make([]*Coordinate, 0)

		_, ok := outMap[startX][startY-1]
		if ok {
			ccont = append(ccont, &Coordinate{X: startX, Y: startY - 1})
		}
		_, ok = outMap[startX][startY+1]
		if ok {
			ccont = append(ccont, &Coordinate{X: startX, Y: startY + 1})
		}
		_, ok = outMap[startX+1][startY]
		if ok {
			ccont = append(ccont, &Coordinate{X: startX + 1, Y: startY})
		}
		_, ok = outMap[startX-1][startY]
		if ok {
			ccont = append(ccont, &Coordinate{X: startX - 1, Y: startY})
		}

		if len(ccont) == 0 {
			if mapSize(outMap) == 0 {
				break
			}

			trace = trace[:len(trace)-1]
			last := trace[len(trace)-1]
			startX, startY = last.X, last.Y
			grids--
			continue
		}

		rand.Shuffle(len(ccont), func(i, j int) {
			ccont[i], ccont[j] = ccont[j], ccont[i]
		})

		next := ccont[0]

		delete(outMap[next.X], next.Y)
		inMap[next.X][next.Y] = struct{}{}
		trace = append(trace, &Coordinate{X: next.X, Y: next.Y})

		startX, startY = next.X, next.Y

		grids++
		if grids == size {
			break
		}
	}

	trace = append([]*Coordinate{&Coordinate{X: 0, Y: 0}}, trace...)
	entryWay(&trace)

	return &Map{cords: trace, Width: width, Height: height}
	//return trace
}
