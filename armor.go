package main

import (
	"fmt"
)

type Armor struct {
	Name   string
	Price  int
	Action ActionFunc
}

type Armors []Armor

func (a Armors) List() {
	for key, value := range a {
		fmt.Printf("[%d]. %s(%d)\n", key+1, value.Name, value.Price)
	}
}

func (a Armors) Action(i int) {
	a[i-1].Action()
}

func (a Armors) Len() int {
	return len(a)
}

func (a Armors) IsShopping() bool {
	return true
}

type ArmorStore struct {
	armors Armors
	aopts  WelcomeOptions
}

func (a *ArmorStore) Welcome() {
	taskCh <- Task{Paragraph: "[ARMOR] What do u want today", Options: a.aopts}
}

func (a *ArmorStore) SellingList() {
	taskCh <- Task{Paragraph: "[WEAPON] Our selling", Options: a.armors}
}

func (a *ArmorStore) BeSoldList() {}

func NewArmorStore() *ArmorStore {
	a := &ArmorStore{}

	a.armors = append(a.armors, Armor{Name: "A Armor"})
	a.armors = append(a.armors, Armor{Name: "B Armor"})
	a.armors = append(a.armors, Armor{Name: "C Armor"})
	a.armors = append(a.armors, Armor{Name: "D Armor"})
	a.armors = append(a.armors, Armor{Name: "E Armor"})
	a.armors = append(a.armors, Armor{Name: "F Armor"})
	a.armors = append(a.armors, Armor{Name: "G Armor"})

	a.aopts = append(a.aopts, WelcomeOption{Description: "[ARMOR] Buy something", Action: a.SellingList})
	a.aopts = append(a.aopts, WelcomeOption{Description: "[ARMOR] Buy back", Action: a.BeSoldList})

	return a
}
