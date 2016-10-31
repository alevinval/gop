package gop

type World interface {
	Push(e State)
	Pop() (e State)
	Peek() (e State)
	Empty() bool
	List() []State
	Reset()
}

type stateStack struct {
	list []State
	pos  int
	cap  int
}

func NewWorld() World {
	return NewWorldSize(1)
}

func NewWorldSize(size int) World {
	return &stateStack{
		list: make([]State, size),
		pos:  0,
		cap:  size,
	}
}

func (s *stateStack) Empty() bool {
	return s.pos == 0
}

func (s *stateStack) Push(e State) {
	if s.pos == s.cap {
		s.grow()
	}
	s.list[s.pos] = e
	s.pos++
}

func (s *stateStack) Pop() (r State) {
	s.pos--
	r = s.list[s.pos]
	return
}

func (s *stateStack) Peek() State {
	return s.list[s.pos-1]
}

func (s *stateStack) List() []State {
	return s.list[:s.pos]
}

func (s *stateStack) Reset() {
	s.pos = 0
}

func (s *stateStack) grow() {
	s.cap <<= 2
	l := make([]State, s.cap)
	copy(l, s.list)
	s.list = l
}
