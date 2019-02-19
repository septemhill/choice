package main

import "fmt"

type Shop interface {
	Welcome()
	SellingList()
	BeSoldList()
}

type WelcomeOption struct {
	Description string
	Action      ActionFunc
}

type WelcomeOptions []WelcomeOption

func (w WelcomeOptions) List() {
	for key, value := range w {
		fmt.Printf("[%d]. %s\n", key+1, value.Description)
	}
}

func (w WelcomeOptions) Action(i int) {
	w[i-1].Action()
}

func (w WelcomeOptions) Len() int {
	return len(w)
}

func (w WelcomeOptions) IsShopping() bool {
	return false
}

//type ActionFunc func()
//
//type Listable interface {
//	List()
//	Len() int
//	IsShopping() bool
//	Action(int)
//}
