package gop

const UPPER_BOUND = 1000000

var (
	r1        = NewWorld()
	r2        = NewWorld()
	satisfied = NewWorld()
	missing   = NewWorld()
)

type State interface {
	Name() string
	Action(goal World) Action
}

type Action interface {
	Name() string
	PreConditions(world World) []State
	PostConditions(world World) []State
}

func EqualStacks(w1, w2 World) (eq bool) {
	if len(w1.List()) != len(w2.List()) {
		return false
	}
	if len(w1.List()) == 0 {
		return true
	}
	eq = true
	var s1, s2 State
	for !w1.Empty() && !w2.Empty() {
		s1 = w1.Pop()
		r1.Push(s1)
		for !w2.Empty() {
			s2 = w2.Pop()
			r2.Push(s2)
			if s1.Name() == s2.Name() {
				break
			}
		}
		if s1.Name() != s2.Name() {
			eq = false
		}
		for !r2.Empty() {
			w2.Push(r2.Pop())
		}
	}
	for !r1.Empty() {
		w1.Push(r1.Pop())
	}
	return
}

func StatesArePresent(w World, states []State) (allPresent bool) {
	satisfied.Reset()
	missing.Reset()
	if states == nil || len(states) == 0 {
		return true
	}
	allPresent = true
	for _, s := range states {
		isPresent := StateIsPresent(w, s)
		allPresent = allPresent && isPresent
		if isPresent {
			satisfied.Push(s)
		} else {
			missing.Push(s)
		}
	}
	return allPresent
}

func StateIsPresent(w World, s State) bool {
	for _, wS := range w.List() {
		if wS.Name() == s.Name() {
			return true
		}
	}
	return false
}

func delStates(w World, states []State) {
	var ws State
	for _, state := range states {
		for !w.Empty() {
			ws = w.Pop()
			if ws.Name() == state.Name() {
				//fmt.Printf("Removing state: %q\n", state)
				break
			}
			r1.Push(ws)
		}
		for !r1.Empty() {
			w.Push(r1.Pop())
		}
	}
}

func addStates(w World, states []State) {
	for _, state := range states {
		//fmt.Printf("	Adding: %q\n", state)
		w.Push(state)
	}
}

func BuildPlan(world, goal World) []Action {
	plan := []Action{}
	pending := NewWorldSize(len(goal.List()))

	var N int
	for !EqualStacks(world, goal) && N < UPPER_BOUND {
		N++
		if pending.Empty() {
			for _, s := range goal.List() {
				pending.Push(s)
			}
		}

		//fmt.Printf("  World: %s\n", world.List())
		//fmt.Printf("Pending: %s\n", pending.List())
		//fmt.Printf("  Goal: %s\n", goal.List())
		//fmt.Println("")

		desiredState := pending.Peek()

		//fmt.Printf("Is state: %q satisfied? %t\n", desiredState.Name(), StatesArePresent(world, desiredState))
		if StateIsPresent(world, desiredState) {
			pending.Pop()
			//fmt.Printf("	Removing: %q\n", p)
			continue
		}

		action := desiredState.Action(goal)
		if action == nil {
			pending.Pop()
			continue
		}

		preconditions := action.PreConditions(world)
		ok := StatesArePresent(world, preconditions)
		if ok {
			delStates(world, preconditions)
			addStates(world, action.PostConditions(world))
			plan = append(plan, action)
			//fmt.Printf("Pushing action: %s\n", action)
		} else {
			delStates(pending, satisfied.List())
			addStates(pending, missing.List())
		}
		//fmt.Println("")
	}
	return plan
}
