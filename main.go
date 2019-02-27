package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
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

func main() {
	m := CreateMap(6, 6)

	t := &Team{
		Characters: []Character{
			Character{Name: "Septem"},
			Character{Name: "Nicole"},
			Character{Name: "Asolia"},
		},
	}

	m.Enter(t)
	m.Walk()
}

//func main() {
//	router := gin.Default()
//	go RestfulService(router)
//}
