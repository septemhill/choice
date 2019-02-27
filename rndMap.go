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
	X      int
	Y      int
	Up     bool
	Down   bool
	Left   bool
	Right  bool
	Events []CoordinateEvent
}

type Map struct {
	cords    []*Coordinate
	cordsMap map[string]*Coordinate
	Width    int
	Height   int
	curr     *Coordinate
	team     *Team
}

func (m *Map) exist(x, y int) bool {
	str := fmt.Sprintf("%d:%d", x, y)
	_, ok := m.cordsMap[str]

	return ok
}

func (m *Map) find(x, y int) *Coordinate {
	str := fmt.Sprintf("%d:%d", x, y)
	d, ok := m.cordsMap[str]

	if ok {
		return d
	}

	return nil
}

func (m *Map) setCurrent(x, y int) {
	m.curr = m.find(x, y)

	//Trigger events
	events := m.curr.Events

	for i := 0; i < len(events); i++ {
		events[i].Trigger(m.team)
	}
}

func (m *Map) drawMapDashline(width, height, space int, coord []*Coordinate) {
	dash := "-"

	for i := 0; i < width; i++ {
		if m.exist(i, height-1) && m.exist(i, height) {
			if m.exist(i, height) && m.exist(i+1, height) && m.exist(i+1, height-1) {
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
		if m.exist(i, height) {
			if m.curr == m.find(i, height) {
				grid += fion.BYellow(strings.Repeat(" ", space))
			} else {
				grid += fion.BRed(strings.Repeat(" ", space))

			}

			if m.exist(i+1, height) {
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
		if m.exist(i, height-1) && m.exist(i, height) {
			up, down := m.find(i, height-1), m.find(i, height)
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
		if m.exist(i, height) && m.exist(i+1, height) {
			left, right := m.find(i, height), m.find(i+1, height)
			if left.Right || right.Left {
				if m.curr == left {
					//grid += fion.BYellow(strings.Repeat(" ", space+1))
					grid += fion.BYellow(strings.Repeat(" ", space))
					grid += fion.BRed(" ")
				} else {
					grid += fion.BRed(strings.Repeat(" ", space+1))
				}
			} else {
				if m.curr == left {
					grid += fion.BYellow(strings.Repeat(" ", space)) + "|"
				} else {
					grid += fion.BRed(strings.Repeat(" ", space)) + "|"
				}
			}
		} else if m.exist(i, height) {
			if m.curr == m.find(i, height) {
				grid += fion.BYellow(strings.Repeat(" ", space)) + "|"

			} else {
				grid += fion.BRed(strings.Repeat(" ", space)) + "|"

			}
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

func (m *Map) Enter(t *Team) {
	m.curr = m.cords[0]
	m.team = t
}

func (m *Map) Walk() {
	for {
		var input int

		EraseDisplay(CLR_ENTIRE_ALL)
		MoveTo(1, 1)

		m.DrawMap()

		fmt.Println("[1] UP")
		fmt.Println("[2] DOWN")
		fmt.Println("[3] LEFT")
		fmt.Println("[4] RIGHT")

		fmt.Scanf("%d", &input)

		if input < 1 || input > 4 {
			continue
		}

		if input == 1 && m.exist(m.curr.X, m.curr.Y-1) {
			m.setCurrent(m.curr.X, m.curr.Y-1)
		} else if input == 2 && m.exist(m.curr.X, m.curr.Y+1) {
			m.setCurrent(m.curr.X, m.curr.Y+1)
		} else if input == 3 && m.exist(m.curr.X-1, m.curr.Y) {
			m.setCurrent(m.curr.X-1, m.curr.Y)
		} else if input == 4 && m.exist(m.curr.X+1, m.curr.Y) {
			m.setCurrent(m.curr.X+1, m.curr.Y)
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

func setupEvents(cords *[]*Coordinate) {
	rndpercent := rand.Intn(5) + 5
	eventCount := len(*cords) * rndpercent / 100
	rand.Shuffle(len(*cords), func(i, j int) {
		(*cords)[i], (*cords)[j] = (*cords)[j], (*cords)[i]
	})

	for i := 0; i < eventCount; i++ {
		(*cords)[i].Events = append((*cords)[i].Events, eventList[i%EVT_MAX])
	}
}

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

	cordsmap := make(map[string]*Coordinate)
	for i := 0; i < len(trace); i++ {
		str := fmt.Sprintf("%d:%d", trace[i].X, trace[i].Y)
		cordsmap[str] = trace[i]
	}

	return &Map{cords: trace, cordsMap: cordsmap, Width: width, Height: height}
}
