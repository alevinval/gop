package main

import (
	"fmt"

	"github.com/alevinval/gop"
)

const (
	GLASS_IS_EMPTY state = iota
	GLASS_IS_FULL

	THIRSTY
	NOT_THIRSTY
)

type state byte

func (s state) Name() string {
	switch s {
	case GLASS_IS_EMPTY:
		return "glass-is-empty"
	case GLASS_IS_FULL:
		return "glass-is-full"
	case THIRSTY:
		return "thirsty"
	case NOT_THIRSTY:
		return "not-thirsty"
	default:
		return "{state not implemented}"
	}
}

func (s state) String() string {
	return s.Name()
}

func (s state) Action(_ gop.World) (a gop.Action) {
	switch s {
	case GLASS_IS_FULL:
		a = FILL_GLASS
	}
	return
}

type Person struct {
	name  string
	state gop.State
}

func NewPerson(name string, state gop.State) *Person {
	return &Person{name: name, state: state}
}

func (p *Person) PersonName() string {
	return p.name
}

func (p *Person) Name() string {
	return fmt.Sprintf("Person{%s, %s}", p.name, p.state)
}

func (p *Person) String() string {
	return p.Name()
}

func (p *Person) Action(_ gop.World) (a gop.Action) {
	switch p.state {
	case NOT_THIRSTY:
		a = &Drink{p.PersonName()}
	}
	return
}
