package benchmarks

import (
	"fmt"

	"github.com/alevinval/gop"
)

type Flubber int

func (f Flubber) Name() string {
	return fmt.Sprintf("flubber: %d", f)
}
func (f Flubber) String() string { return f.Name() }

func (f Flubber) Actions(goal gop.Stack) []gop.Action {
	target := goal.Peek().(Flubber)
	if f%2 == 0 && f/2*2 <= target {
		return []gop.Action{Double(f / 2)}
	} else {
		return []gop.Action{Increment(f - 1)}
	}
}

type Double int

func (sf Double) PreConditions(w gop.Stack) []gop.State {
	return []gop.State{Flubber(sf)}
}

func (sf Double) PostConditions(w gop.Stack) []gop.State {
	return []gop.State{Flubber(sf * 2)}
}

func (sf Double) Name() string {
	return fmt.Sprintf("double %d", sf)
}

func (sf Double) String() string {
	return sf.Name()
}

type Increment int

func (mf Increment) PreConditions(w gop.Stack) []gop.State {
	return []gop.State{Flubber(mf)}
}

func (mf Increment) PostConditions(w gop.Stack) []gop.State {
	return []gop.State{Flubber(mf + 1)}
}

func (mf Increment) Name() string {
	return fmt.Sprintf("increment %d", mf)
}

func (mf Increment) String() string {
	return mf.Name()
}
