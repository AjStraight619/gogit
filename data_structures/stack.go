package datastructures

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Stack[T any] struct {
	length int
	head   *Node[T]
}

func (s *Stack[T]) Push(value T) {
	newNode := &Node[T]{value: value, next: s.head}
	s.head = newNode
	s.length++
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.length == 0 {
		var zeroValue T
		return zeroValue, false
	}

	value := s.head.value
	s.head = s.head.next
	s.length--
	return value, true

}

func (s *Stack[T]) Peek() (T, bool) {
	if s.length == 0 {
		var zeroValue T
		return zeroValue, false
	}

	return s.head.value, true
}

func (s *Stack[T]) Print() {

	current := s.head

	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}

}
