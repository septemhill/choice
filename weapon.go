package main

import (
	"fmt"
)

type Weapon struct {
	Name   string
	Price  int
	Action ActionFunc
}

type Weapons []Weapon

func (w Weapons) List() {
	for key, value := range w {
		fmt.Printf("[%d]. %s(%d)\n", key+1, value.Name, value.Price)
	}
}

func (w Weapons) Action(i int) {
	w[i-1].Action()
}

func (w Weapons) Len() int {
	return len(w)
}

func (w Weapons) IsShopping() bool {
	return true
}

type WeaponStore struct {
	weapons Weapons
	wopts   WelcomeOptions
}

func (w *WeaponStore) Welcome() {
	taskCh <- Task{Paragraph: "[WEAPON] What do u want today?", Options: w.wopts}
}

func (w *WeaponStore) SellingList() {
	taskCh <- Task{Paragraph: "[WEAPON] Our selling", Options: w.weapons}
}

func (w *WeaponStore) BeSoldList() {}

func NewWeaponStore() *WeaponStore {
	w := &WeaponStore{}

	w.weapons = append(w.weapons, Weapon{Name: "A sword"})
	w.weapons = append(w.weapons, Weapon{Name: "B sword"})
	w.weapons = append(w.weapons, Weapon{Name: "C sword"})
	w.weapons = append(w.weapons, Weapon{Name: "D sword"})
	w.weapons = append(w.weapons, Weapon{Name: "E sword"})
	w.weapons = append(w.weapons, Weapon{Name: "F sword"})

	w.wopts = append(w.wopts, WelcomeOption{Description: "Buy something", Action: w.SellingList})
	w.wopts = append(w.wopts, WelcomeOption{Description: "Buy back ", Action: w.BeSoldList})

	return w
}
