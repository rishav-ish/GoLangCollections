package collections

import "errors"

type Stack struct {
	elements []any
}

func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (any, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("Stack is empty")
	}
	top := s.elements[len(s.elements)-1]
	// s.elements = s.elements[:len(s.elements)-1]
	s.elements = append([]any{}, s.elements[:len(s.elements)-1]...)
	return top, nil
}

func (s *Stack) Top() (any, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("Stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}

type StackV2 []any

func (s *StackV2) Push(val any) {
	*s = append(*s, val)
}

func (s *StackV2) Pop() (any, error) {
	if len(*s) == 0 {
		return 0, errors.New("stack is empty")
	}

	top := (*s)[len(*s)-1]
	*s = append([]any{}, (*s)[:len(*s)-1]...)
	return top, nil
}

func (s *StackV2) Top() (any, error) {
	if len(*s) == 0 {
		return 0, errors.New("stack is empty")
	}
	return (*s)[len(*s)-1], nil
}

func (s *StackV2) IsEmpty() bool {
	return len(*s) == 0
}

func (s *StackV2) Size() int {
	return len(*s)
}
