package main

import (
	"fmt"

	"github.com/alevinval/gop"
)

func main() {
	ini := gop.NewWorld()
	ini.Push(GLASS_IS_EMPTY)
	ini.Push(NewPerson("john", THIRSTY))
	ini.Push(NewPerson("doe", THIRSTY))

	fin := gop.NewWorld()
	fin.Push(GLASS_IS_FULL)
	fin.Push(NewPerson("doe", NOT_THIRSTY))
	fin.Push(NewPerson("john", NOT_THIRSTY))

	plan := gop.BuildPlan(ini, fin)
	fmt.Printf("Plan: %s\n", plan)
}
