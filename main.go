package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"

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
		if x == cords[i].X && y == cords[i].Y {
			return true
		}
	}

	return false
}

func find(x, y int, cords []*Coordinate) *Coordinate {
	for i := 0; i < len(cords); i++ {
		if x == cords[i].X && y == cords[i].Y {
			return cords[i]
		}
	}

	return nil
}

func drawMapDashline(width, height, space int, coord []*Coordinate) {
	dash := "-"

	for i := 0; i < width; i++ {
		if exist(i, height-1, coord) && exist(i, height, coord) {
			if exist(i, height, coord) && exist(i+1, height, coord) {
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

func drawMapGridColumn(width, height, space int, coord []*Coordinate) {
	grid := "|"

	for i := 0; i < width; i++ {
		if exist(i, height, coord) {
			grid += fion.BRed(strings.Repeat(" ", space))

			if exist(i+1, height, coord) {
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

func drawMap(width, height int, coord []*Coordinate) {
	dashLineCount := height + 1
	space := 3
	h := 0

	for i := 0; i < dashLineCount+height; i++ {
		if i%2 == 0 {
			drawMapDashline(width, h, space, coord)
		} else {
			drawMapGridColumn(width, h, space, coord)
			h++
		}
	}
}

func drawPathDashline(width, height, space int, cords []*Coordinate) {
	dash := "-"

	for i := 0; i < width; i++ {
		if exist(i, height-1, cords) && exist(i, height, cords) {
			up, down := find(i, height-1, cords), find(i, height, cords)
			if up.Down || down.Up {
				dash += fion.BRed(strings.Repeat(" ", space)) + "-"
			} else {
				dash += strings.Repeat("-", space+1)
			}
		} else {
			dash += strings.Repeat("-", space+1)
		}
		//		if exist(i, height-1, cords) {
		//			cord := find(i, height-1, cords)
		//			if cord.Down {
		//				dash += fion.BRed(strings.Repeat(" ", space)) + "-"
		//			} else {
		//				dash += strings.Repeat("-", space+1)
		//			}
		//		} else if exist(i, height, cords) {
		//			cord := find(i, height, cords)
		//			//			fmt.Println(i, height, cord)
		//			if cord.Up {
		//				dash += fion.BRed(strings.Repeat(" ", space)) + "-"
		//			} else {
		//				dash += strings.Repeat("-", space+1)
		//			}
		//		} else {
		//			//fmt.Println("FDSA")
		//			dash += strings.Repeat("-", space+1)
		//		}
	}
	fmt.Println(dash)
}

func drawPathGridColumn(width, height, space int, cords []*Coordinate) {
	drawMapGridColumn(width, height, space, cords)
}

func drawPath(width, height int, coord []*Coordinate) {
	dashLineCount := height + 1
	space := 3
	h := 0

	for i := 0; i < dashLineCount+height; i++ {
		if i%2 == 0 {
			drawPathDashline(width, h, space, coord)
		} else {
			drawPathGridColumn(width, h, space, coord)
			h++
		}
	}
}

func main() {
	width, height := 20, 20

	cords := CreateMap(width, height)
	drawPath(width, height, cords)

	for i := 0; i < len(cords); i++ {
		fmt.Println(cords[i])
	}
}
