package main

type Attribute struct {
}

type Character struct {
	Name string
	Attr Attribute
}

type Team struct {
	Characters []Character
	Money      int64
}
