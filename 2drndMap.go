package main

import (
	"math/rand"
	"time"
)

type Coordinate struct {
	X     int
	Y     int
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

type CoordinateMap map[int]map[int]struct{}

func mapSize(m CoordinateMap) int {
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

func CreateMap(width, height int) []*Coordinate {
	inMap, outMap := make(CoordinateMap), make(CoordinateMap)
	trace := make([]*Coordinate, 0)
	size := width * height * 50 / 100
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

	return trace
}
