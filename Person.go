package main

type person struct {
	name     string
	lastname string
}

func NewPerson(name, lastname string) *person {
	return &person{name, lastname}
}

func (p *person) GetName() string {
	return p.name
}
