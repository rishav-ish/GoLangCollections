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

type StackV2[T comparable] []T

func (s *StackV2[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *StackV2[T]) Pop() (T, error) {
	if len(*s) == 0 {
		var zero_val T
		return zero_val, errors.New("stack is empty")
	}

	top := (*s)[len(*s)-1]
	*s = append([]T{}, (*s)[:len(*s)-1]...)
	return top, nil
}

func (s *StackV2[T]) Top() (T, error) {
	if len(*s) == 0 {
		var zero_val T
		return zero_val, errors.New("stack is empty")
	}
	return (*s)[len(*s)-1], nil
}

func (s *StackV2[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *StackV2[T]) Size() int {
	return len(*s)
}

func (s *StackV2[T]) Sort(compare func(a, b T) bool) {
	if s.Size() == 0 || s.Size() == 1 {
		return
	}

	data, _ := s.Pop()
	s.Sort(compare)
	s.insert(data, compare)
}

func (s *StackV2[T]) insert(val T, compare func(a, b T) bool) {
	if s.IsEmpty() {
		s.Push(val)
		return
	}

	if top_element, _ := s.Top(); compare(top_element, val) {
		s.Push(val)
		return
	}

	temp, _ := s.Pop()
	s.insert(val, compare)
	s.Push(temp)
}
