package stack

import "github.com/ChristianSchleifer/datastructures/list"

type LinkedListStack struct {
	l *list.DoublyLinkedList
}

// NewLinkedListStack constructs a ready to use stack
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		l: &list.DoublyLinkedList{},
	}
}

// Pop removes an element from the stack
func (stack *LinkedListStack) Pop() (int, error) {
	return stack.l.Pop()
}

// Push adds an element to the stack
func (stack *LinkedListStack) Push(val int) {
	stack.l.Push(val)
}
