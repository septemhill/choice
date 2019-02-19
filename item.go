package main

type Consumable interface {
	Use()
}

type Unconsumable interface {
}

type Item struct {
	Name           string
	Description    string
	CanUseInBattle bool
}

type Sword struct {
	Item
	Damage int
}
