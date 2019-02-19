package main

import "fmt"

type Misc struct {
	Name   string
	Price  int
	Action ActionFunc
}

type Miscs []Misc

type MiscStore struct {
	miscs Miscs
	mopts WelcomeOptions
}

func (m Miscs) List() {
	for key, value := range m {
		fmt.Printf("[%d]. %s(%d)\n", key+1, value.Name, value.Price)
	}
}

func (m Miscs) Action(i int) {
	m[i-1].Action()
}

func (m Miscs) Len() int {
	return len(m)
}

func (m Miscs) IsShopping() bool {
	return true
}

func (m *MiscStore) Welcome() {
	taskCh <- Task{Paragraph: "[MISC] What do u want today?", Options: m.mopts}
}

func (m *MiscStore) SellingList() {
	taskCh <- Task{Paragraph: "[MISC] Our selling", Options: m.miscs}
}

func (m *MiscStore) BeSoldList() {
}

func NewMiscStore() *MiscStore {
	m := &MiscStore{}

	m.miscs = append(m.miscs, Misc{Name: "A misc"})
	m.miscs = append(m.miscs, Misc{Name: "B misc"})
	m.miscs = append(m.miscs, Misc{Name: "C misc"})
	m.miscs = append(m.miscs, Misc{Name: "D misc"})
	m.miscs = append(m.miscs, Misc{Name: "E misc"})

	m.mopts = append(m.mopts, WelcomeOption{Description: "Buy something", Action: m.SellingList})
	m.mopts = append(m.mopts, WelcomeOption{Description: "Buy something", Action: m.BeSoldList})

	return m
}
