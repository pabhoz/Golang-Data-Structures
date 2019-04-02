package stack

import (
	dll "../DLL"
)

type (
	//Stack Basic stack struct
	Stack struct {
		top      *dll.Element
		elements *dll.List
		max      int
	}
)

// New Creates a new stack
func New(max int) *Stack {
	list := dll.New()
	return &Stack{nil, list, max}
}

// Push add new element to stack
func (s *Stack) Push(element interface{}) *dll.Element {
	s.top = s.elements.PushBack(element)
	return s.top
}

// Pop pop the first element and returns its value
func (s *Stack) Pop() interface{} {
	toRemove := s.top
	//fmt.Println("toRemove: ", toRemove)
	s.top = s.top.Prev()
	//fmt.Println("Prev: ", s.top)
	return s.elements.Remove(toRemove)
}

//Top returns top element of the stack
func (s *Stack) Top() *dll.Element {
	return s.top
}

// Elements return dll with stack elements
func (s *Stack) Elements() *dll.List {
	return s.elements
}

/********************************
*****Implement these methods*****
********************************/

// Empty clears the stack
func (s *Stack) Empty() {

}

// Size returns stack size
func (s *Stack) Size() int {
	size := 0
	return size
}
