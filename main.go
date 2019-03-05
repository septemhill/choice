package main

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/sys/unix"
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

func getTerminalSize() (int, int) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)

	if err != nil {
		return -1, -1
	}

	return int(ws.Col), int(ws.Row)
}

func main() {
	//	m := CreateMap(6, 6)
	//
	//	t := &Team{
	//		Characters: []Character{
	//			Character{Name: "Septem"},
	//			Character{Name: "Nicole"},
	//			Character{Name: "Asolia"},
	//		},
	//	}
	//
	//	m.Enter(t)
	//	EraseDisplay(CLR_ENTIRE_ALL)
	//	m.Walk()

	//tbox := &Box{bytes.NewBuffer(nil), 1, 1, 20, 20}
	//fmt.Fprintf(tbox, fion.BRed("Hi, Septem 科科科科科我不是天才喝科八"))
	//tbox.Draw()

	//str := fion.BRed("Septem吃飯囉") + ("Nicole go home囉") + fion.BBlue("Seednia so coooool!!")
	//es := stringParse(str)
	//fmt.Println(es.SubstringByWidth(4, 30))

	_, height := getTerminalSize()

	twbox := &Box{bytes.NewBuffer(nil), 1, 1, 20, 10}
	fmt.Fprintf(twbox, "QQ哭枯喔")
	twbox.Draw()

	moveToPaint(2, 5, "A")
	moveToPaint(height, 1, "A")
}
