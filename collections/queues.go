package collections

import "errors"

type Queue[T any] []T

func (q *Queue[T]) Enqueue(val T) {
	*q = append(*q, val)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero_val T
		return zero_val, errors.New("queue is empty")
	}
	front := (*q)[0]
	*q = (*q)[1:]
	return front, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero_val T
		return zero_val, errors.New("queue is empty")
	}
	return (*q)[0], nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Size() int {
	return len(*q)
}
