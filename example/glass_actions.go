package main

import (
	"fmt"

	"github.com/alevinval/gop"
)

const (
	EMPTY_GLASS action = iota
	FILL_GLASS
)

type action byte

func (a action) PreConditions(world gop.World) []gop.State {
	switch a {
	case EMPTY_GLASS:
		return []gop.State{GLASS_IS_FULL}
	case FILL_GLASS:
		return []gop.State{GLASS_IS_EMPTY}
	default:
		return []gop.State{}
	}
}

func (a action) PostConditions(world gop.World) []gop.State {
	switch a {
	case EMPTY_GLASS:
		return []gop.State{GLASS_IS_EMPTY}
	case FILL_GLASS:
		return []gop.State{GLASS_IS_FULL}
	default:
		return []gop.State{}
	}
}

func (a action) Name() string {
	switch a {
	case FILL_GLASS:
		return "fill-glass"
	case EMPTY_GLASS:
		return "empty-glass"
	default:
		return "{action not implemented}"
	}
}

func (a action) String() string {
	return a.Name()
}

type Drink struct {
	name string
}

func (d *Drink) Name() string {
	return fmt.Sprintf("%q drinks", d.name)
}

func (d *Drink) String() string {
	return d.Name()
}

func (d *Drink) PreConditions(world gop.World) []gop.State {
	return []gop.State{NewPerson(d.name, THIRSTY), GLASS_IS_FULL}
}

func (d *Drink) PostConditions(world gop.World) []gop.State {
	return []gop.State{NewPerson(d.name, NOT_THIRSTY), GLASS_IS_EMPTY}
}
