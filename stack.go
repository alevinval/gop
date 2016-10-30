package gop

type Type interface{}

type Stack interface {
	Push(e Type)
	Pop() (e Type)
	Peek() (e Type)
	Empty() bool
	List() []Type
}

type stack struct {
	list []Type
	pos  int
	cap  int
}

func NewStack() Stack {
	return &stack{
		list: make([]Type, 1),
		pos:  0,
		cap:  1,
	}
}

func (s *stack) Empty() bool {
	return s.pos == 0
}

func (s *stack) Push(e Type) {
	if s.pos == s.cap {
		s.grow()
	}
	s.list[s.pos] = e
	s.pos++
}

func (s *stack) Pop() (r Type) {
	s.pos--
	r = s.list[s.pos]
	return
}

func (s *stack) Peek() Type {
	return s.list[s.pos-1]
}

func (s *stack) List() []Type {
	return s.list[:s.pos]
}

func (s *stack) grow() {
	s.cap <<= 2
	l := make([]Type, s.cap)
	copy(l, s.list)
	s.list = l
}
