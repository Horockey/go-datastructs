package stack

import "errors"

var ErrEmptyStack = errors.New("stack is empty")

type Stack[T any] struct {
	top *node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack[T]) Push(val T) {
	t := newNode(val)
	t.next = s.top
	s.top = t
}
func (s *Stack[T]) Pop() (T, error) {
	if s.top == nil {
		return *new(T), ErrEmptyStack
	}
	t := s.top.payload
	s.top = s.top.next
	return t, nil
}
func (s *Stack[T]) Get() (T, error) {
	if s.top == nil {
		return *new(T), ErrEmptyStack
	}
	return s.top.payload, nil
}
