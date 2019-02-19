package main

import "fmt"

type Place interface {
	Enter(t *Team)
}

type Store struct {
	Name        string
	Description string
	Action      ActionFunc
}

type Stores []Store

func (s Stores) List() {
	for key, value := range s {
		fmt.Printf("[%d] %s\n", key+1, value.Name)
	}
}

func (s Stores) Len() int { return len(s) }

func (s Stores) Action(i int) { s[i-1].Action() }

func (s Stores) IsShopping() bool { return false }

func (s Stores) Enter(t *Team) {
}

type Villeage struct {
	Name string
}

func (v Villeage) Enter(t *Team) {
	//var p Places
	var s Stores

	weaponStore := NewWeaponStore()
	armorStore := NewArmorStore()
	miscStore := NewMiscStore()

	s = append(s, Store{Name: "Weapon Store", Description: "Buy / Sell Weapons", Action: weaponStore.Welcome})
	s = append(s, Store{Name: "Armor Store", Description: "Buy / Sell Armors", Action: armorStore.Welcome})
	s = append(s, Store{Name: "Misc Store", Description: "Buy / Sell Misc Items", Action: miscStore.Welcome})

	storeTask := Task{Paragraph: "Where u want to go", Options: s}

	taskCh <- storeTask
}

func NewVilleage(name string) *Villeage {
	v := &Villeage{
		Name: name,
	}

	return v
}
