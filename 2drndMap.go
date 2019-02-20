package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	UP_WAY    = 0
	DOWN_WAY  = ^uint(0) //18446744073709551615
	LEFT_WAY  = 1
	RIGHT_WAY = ^uint(1) //18446744073709551614
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

func CreateMap(width, height int) []*Coordinate {
	outMap := make(CoordinateMap)
	inMap := make(CoordinateMap)
	trace := make([]*Coordinate, 0)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < width; i++ {
		outMap[i] = make(map[int]struct{})
		inMap[i] = make(map[int]struct{})
		for j := 0; j < height; j++ {
			outMap[i][j] = struct{}{}
		}
	}

	startX, startY := 0, 0

	delete(outMap[0], 0)
	inMap[0][0] = struct{}{}

	size := width * height * 60 / 100

	fmt.Println(0, 0)
	for grids := 0; grids < size; {
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
				trace = trace[:len(trace)-1]
				last := trace[len(trace)-1]
				//fmt.Println(last.X, last.Y)
				startX, startY = last.X, last.Y
				grids--
				continue
			}

			rand.Shuffle(len(ccont), func(i, j int) {
				ccont[i], ccont[j] = ccont[j], ccont[i]
			})

			s := ccont[0]

			startX, startY = s.X, s.Y

			delete(outMap[s.X], s.Y)
			inMap[s.X][s.Y] = struct{}{}

			trace = append(trace, s)

			break
		}
		grids++
	}

	trace = append([]*Coordinate{&Coordinate{X: 0, Y: 0}}, trace[0:]...)
	//for i := 0; i < len(trace); i++ {
	//	fmt.Println("T", trace[i])
	//}
	//fmt.Println("TRACE", len(trace))
	return trace
}
