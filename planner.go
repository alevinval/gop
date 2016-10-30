package gop

import (
	_ "fmt"
)

const UPPER_BOUND = 1000000

type State interface {
	Name() string
	Actions(goal Stack) []Action
}

type Action interface {
	Name() string
	PreConditions(world Stack) []State
	PostConditions(world Stack) []State
}

func EqualStacks(w1, w2 Stack) (eq bool) {
	if len(w1.List()) != len(w2.List()) {
		return false
	}
	if len(w1.List()) == 0 {
		return true
	}
	w1restore := NewStack()
	w2restore := NewStack()
	eq = true
	for !w1.Empty() && !w2.Empty() {
		s1, _ := w1.Pop().(State)
		w1restore.Push(s1)
		var s2 State
		for !w2.Empty() {
			s2, _ = w2.Pop().(State)
			w2restore.Push(s2)
			if s1.Name() == s2.Name() {
				break
			}
		}
		if s1.Name() != s2.Name() {
			eq = false
		}
		for !w2restore.Empty() {
			w2.Push(w2restore.Pop())
		}
	}
	for !w1restore.Empty() {
		w1.Push(w1restore.Pop())
	}
	return
}

func StatesArePresent(w Stack, states []State) (allPresent bool, present []State, missing []State) {
	if states == nil || len(states) == 0 {
		return true, []State{}, []State{}
	}
	var pCount, mCount int
	present = make([]State, len(states))
	missing = make([]State, len(states))
	allPresent = true
	for _, s := range states {
		isPresent := StateIsPresent(w, s)
		allPresent = allPresent && isPresent
		if isPresent {
			present[pCount] = s
			pCount++
		} else {
			missing[mCount] = s
			mCount++
		}
	}
	return allPresent, present[:pCount], missing[:mCount]
}

func StateIsPresent(w Stack, s State) bool {
	for _, wS := range w.List() {
		wS, _ := wS.(State)
		if wS.Name() == s.Name() {
			return true
		}
	}
	return false
}

func delStates(w Stack, states []State) {
	tmp := NewStack()
	for _, state := range states {
		for !w.Empty() {
			wS, _ := w.Pop().(State)
			if wS.Name() == state.Name() {
				//fmt.Printf("Removing state: %q\n", state)
				break
			}
			tmp.Push(wS)
		}
		for !tmp.Empty() {
			w.Push(tmp.Pop())
		}
	}
}

func addStates(w Stack, states []State) {
	for _, state := range states {
		//fmt.Printf("	Adding: %q\n", state)
		w.Push(state)
	}
}

func BuildPlan(world, goal Stack) Stack {
	plan := NewStack()
	pending := NewStack()

	var N int
	for !EqualStacks(world, goal) && N < UPPER_BOUND {
		N++
		if pending.Empty() {
			for _, e := range goal.List() {
				pending.Push(e)
			}
		}

		//fmt.Printf("  World: %s\n", world.List())
		//fmt.Printf("Pending: %s\n", pending.List())
		//fmt.Printf("  Goal: %s\n", goal.List())
		//fmt.Println("")

		desiredState, _ := pending.Peek().(State)

		//fmt.Printf("Is state: %q satisfied? %t\n", desiredState.Name(), StatesArePresent(world, desiredState))
		if StateIsPresent(world, desiredState) {
			pending.Pop()
			//fmt.Printf("	Removing: %q\n", p)
			continue
		}

		var action Action
		actions := desiredState.Actions(goal)
		if len(actions) == 0 {
			pending.Pop()
			continue
		}
		for _, a := range actions {
			action = a
			break
		}
		preconditions := action.PreConditions(world)
		postconditions := action.PostConditions(world)

		ok, present, missing := StatesArePresent(world, preconditions)
		if ok {
			delStates(world, preconditions)
			addStates(world, postconditions)
			plan.Push(action)
			//fmt.Printf("Pushing action: %s\n", action)
		} else {
			delStates(pending, present)
			addStates(pending, missing)
		}
		//fmt.Println("")
	}

	// Sort plan.
	sortedPlan := NewStack()
	for !plan.Empty() {
		sortedPlan.Push(plan.Pop())
	}
	return sortedPlan
}
