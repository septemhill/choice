package main

import (
	"fmt"
	"math/rand"
)

const (
	EVT_POISON_TRAP = iota
	EVT_NEEDLE_TRAP
	EVT_MONEY_STACK
	EVT_MAX
)

var eventList = map[int]CoordinateEvent{
	EVT_POISON_TRAP: PoisonTrap{},
	EVT_NEEDLE_TRAP: NeedleTrap{},
	EVT_MONEY_STACK: MoneyStack{},
}

type CoordinateEvent interface {
	Trigger(t *Team)
}

type Removable interface {
	TakeAllOrTearDown(t *Team)
}

type PoisonTrap struct {
}

func (p PoisonTrap) Trigger(t *Team) {
	for i := 0; i < len(t.Characters); i++ {
		fmt.Printf("%s, ", t.Characters[i].Name)
	}
	fmt.Println("are poisoning.")
}

func (p PoisonTrap) TakeAllOrTearDown(t *Team) {
}

type NeedleTrap struct {
	RemoveProbability float32
}

func (n NeedleTrap) Trigger(t *Team) {
	for i := 0; i < len(t.Characters); i++ {
		fmt.Printf("%s, ", t.Characters[i].Name)
	}
	fmt.Println("are damaged.")
}

func (n NeedleTrap) TakeAllOrTearDown(t *Team) {
}

type MoneyStack struct {
}

func (m MoneyStack) Trigger(t *Team) {
	fmt.Println("WOWOWOWOWOWOW")
	m.TakeAllOrTearDown(t)
}

func (m MoneyStack) TakeAllOrTearDown(t *Team) {
	t.Money += rand.Int63n(5000)
}
