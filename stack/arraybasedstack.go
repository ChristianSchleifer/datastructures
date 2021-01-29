package stack

import "errors"

// ArrayStack is a ready to use stack implementation
type ArrayStack struct {
	array []int
}

// Pop removes an element from the stack
func (stack *ArrayStack) Pop() (int, error) {
	if len(stack.array) == 0 {
		return 0, errors.New("stack is empty")
	}

	val := stack.array[len(stack.array)-1]
	stack.array = stack.array[:len(stack.array)-1]
	return val, nil
}

// Push adds an element to the stack
func (stack *ArrayStack) Push(val int) {
	stack.array = append(stack.array, val)
	return
}
