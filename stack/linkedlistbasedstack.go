package stack

import "github.com/ChristianSchleifer/datastructures/list"

type linkedListStack struct {
	l list.List
}

// NewLinkedListStack constructs a ready to use stack
func NewLinkedListStack() Stack {
	return &linkedListStack{
		l: list.NewSinglyLinkedList(),
	}
}

// Pop removes an element from the stack
func (stack *linkedListStack) Pop() (int, error) {
	return stack.l.Shift()
}

// Push adds an element to the stack
func (stack *linkedListStack) Push(val int) {
	stack.l.Unshift(val)
}
