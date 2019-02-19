package main

type Attribute struct {
}

type Character struct {
	Name string
	Attr Attribute
}

//type Team []Character
//
//func (t Team) List()            {}
//func (t Team) Action()          {}
//func (t Team) Len() int         { return 0 }
//func (t Team) IsShopping() bool { return false }

type Team struct {
	Characters []Character
	Money      int64
}
