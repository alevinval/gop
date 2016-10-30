package main

import (
	"fmt"

	"github.com/alevinval/gop"
)

func main() {
	ini := gop.NewStack()
	ini.Push(GLASS_IS_EMPTY)
	ini.Push(NewPerson("john", THIRSTY))
	ini.Push(NewPerson("doe", THIRSTY))

	fin := gop.NewStack()
	fin.Push(GLASS_IS_FULL)
	fin.Push(NewPerson("doe", NOT_THIRSTY))
	fin.Push(NewPerson("john", NOT_THIRSTY))

	plan := gop.BuildPlan(ini, fin)
	for !plan.Empty() {
		a, _ := plan.Pop().(gop.Action)
		fmt.Printf("- %s\n", a.Name())
	}

}
