package gop_test

import (
	"fmt"
	"testing"

	"github.com/alevinval/gop"
	"github.com/stretchr/testify/assert"
)

func stateIsPresent(w gop.Stack, s ...gop.State) bool {
	ok, _, _ := gop.StatesArePresent(w, s)
	return ok
}

func TestStatesArePresent(t *testing.T) {
	w := stack(1, 2, 3)

	// Correct assertions
	assert.True(t, stateIsPresent(w, S(1)))
	assert.True(t, stateIsPresent(w, S(2)))
	assert.True(t, stateIsPresent(w, S(3)))
	assert.True(t, stateIsPresent(w, S(1), S(2)))
	assert.True(t, stateIsPresent(w, S(1), S(3)))
	assert.True(t, stateIsPresent(w, S(3), S(2)))
	assert.True(t, stateIsPresent(w, S(3), S(1)))
	assert.True(t, stateIsPresent(w, S(2), S(1), S(3)))

	//// Missing states.
	assert.False(t, stateIsPresent(w, S(5)))
	assert.False(t, stateIsPresent(w, S(1), S(5)))
}

func TestEqualStacks(t *testing.T) {
	a := stack()
	b := stack()
	assert.True(t, gop.EqualStacks(a, b))

	a = stack()
	b = stack(1)
	assert.False(t, gop.EqualStacks(a, b))

	a = stack(1)
	b = stack()
	assert.False(t, gop.EqualStacks(a, b))

	a = stack(1)
	b = stack(1)
	assert.True(t, gop.EqualStacks(a, b))

	a = stack(1)
	b = stack(2)
	assert.False(t, gop.EqualStacks(a, b))

	a = stack(1, 2, 3)
	b = stack(3, 2, 1)
	assert.True(t, gop.EqualStacks(a, b))
}

type MockState int

func (ms MockState) Actions(_ gop.Stack) []gop.Action {
	return []gop.Action{}
}

func (ms MockState) Name() string { return fmt.Sprintf("%d", ms) }

func S(n int) gop.State {
	return MockState(n)
}

func stack(nums ...int) gop.Stack {
	w := gop.NewStack()
	for _, n := range nums {
		w.Push(S(n))
	}
	return w
}
