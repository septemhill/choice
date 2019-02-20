package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/septemhill/fion"
)

type Listable interface {
	List()
	Len() int
	IsShopping() bool
	Action(int)
}

//type Menuable interface {
//	Listable
//	IsShopping() bool
//	Action(int)
//}

type ActionFunc func()

var taskCh = make(chan Task, 10)
var menuStack = list.New()

func Menu(paragraph string, options Listable) {
	var choice int

	fmt.Println(paragraph)

	options.List()

	choice = readUserChoice()

	if (choice > 0) && (choice <= options.Len()) {
		if options.IsShopping() {
			taskCh <- Task{Paragraph: "Thank you, anything else ?", Options: options}
		} else {
			menuStack.PushBack(Task{Paragraph: paragraph, Options: options})
			options.Action(choice)
		}
	} else if choice == 0 {
		elem := menuStack.Back()
		if elem != nil {
			task := elem.Value.(Task) //menuStack.Back().Value.(Task)
			menuStack.Remove(elem)
			taskCh <- task
		} else {
			taskCh <- Task{Paragraph: paragraph, Options: options}
		}
	} else {
		taskCh <- Task{Paragraph: "Sorry, we don't have that one", Options: options}
	}
}

func TaskRoutine() {
	for {
		select {
		case task := <-taskCh:
			Menu(task.Paragraph, task.Options)
		}
	}
}

func readUserChoice() int {
	r := bufio.NewReader(os.Stdin)
	b, _ := r.ReadBytes('\n')

	if b[0] == '\n' {
		return int(^uint(0) >> 1)
	} else {
		numStr := string(b[:len(b)-1])
		num, err := strconv.Atoi(numStr)

		if err != nil {
			return int(^uint(0) >> 1)
		}

		return num
	}
}

//func main() {
//	//Create a team
//	team := &Team{
//		Characters: []Character{
//			Character{Name: "Septem"},
//			Character{Name: "Nicole"},
//		},
//		Money: 10000000,
//	}
//
//	//v := NewVilleage("DQ Town")
//	//v.Enter(team)
//
//	rndMap := CreateRandomMap()
//	rndMap.Enter(team)
//
//	rndMap.Walk()
//
//	//	go TaskRoutine()
//	//	select {}
//}

func drawDashLine(width, space int) {
	dash := "-"

	for i := 0; i < width; i++ {
		dash += strings.Repeat("-", 3+1)
	}

	fmt.Println(dash)
}

func drawGridColumn(width, space int) {
	grid := "|"

	for i := 0; i < width; i++ {
		grid += fion.BRed(strings.Repeat(" ", 3)) + "|"
	}

	fmt.Println(grid)
}

func drawSmallMap(width, height int) {
	dashLineCount := height + 1
	space := 3
	//gridWidth := space + 2

	for i := 0; i < dashLineCount+height; i++ {
		if i%2 == 0 {
			drawDashLine(width, space)
		} else {
			drawGridColumn(width, space)
		}
	}
}

func exist(x, y int, cords []*Coordinate) bool {
	for i := 0; i < len(cords); i++ {
		//fmt.Println(x, y, cords[i][0], cords[i][1])
		if x == cords[i].X && y == cords[i].Y {
			return true
		}
	}

	return false
}

func drawPathGridColumn(width, height, space int, coord []*Coordinate) {
	grid := "|"

	for i := 0; i < width; i++ {
		if exist(i, height, coord) {
			grid += fion.BRed(strings.Repeat(" ", space)) + "|"
		} else {
			grid += strings.Repeat(" ", space) + "|"
		}
	}
	fmt.Println(grid)
}

func drawPath(width, height int, coord []*Coordinate) {
	dashLineCount := height + 1
	space := 3
	h := 0

	for i := 0; i < dashLineCount+height; i++ {
		if i%2 == 0 {
			drawDashLine(width, space)
		} else {
			drawPathGridColumn(width, h, space, coord)
			h++
		}
	}
}

func randomWay(width, height int, cords *[][]int) {
	rand.Seed(time.Now().UnixNano())
	ways := []uint{RIGHT_WAY, LEFT_WAY, DOWN_WAY, UP_WAY}

	startX, startY := 0, 0
	way := uint(0)
	lastWay := uint(0)

	fmt.Println(startX, startY)
	for i := 0; i < 250; i++ {
	ENDLOOP:
		for {
			rand.Shuffle(len(ways), func(i, j int) {
				ways[i], ways[j] = ways[j], ways[i]
			})

			way = ways[0]

			if way == lastWay {
				continue
			}

			switch way {
			case UP_WAY:
				if (startY - 1) >= 0 {
					startY--
					break ENDLOOP
				}
			case DOWN_WAY:
				if (startY + 1) < height {
					startY++
					break ENDLOOP
				}
			case RIGHT_WAY:
				if (startX + 1) < width {
					startX++
					break ENDLOOP
				}
			case LEFT_WAY:
				if (startX - 1) >= 0 {
					startX--
					break ENDLOOP
				}
			}
		}
		lastWay = ^uint(way)
		*cords = append(*cords, []int{startX, startY})
		//fmt.Println(startX, startY)
	}
}

func main() {
	//Create2DRandomMap(7, 7)
	//cords := make([][]int, 0)
	//randomWay(40, 40, &cords)
	cords := CreateMap(20, 20)
	drawPath(20, 20, cords)
	//drawPath(40, 40, cords)

}
