package stack

// Stack is the interface describing methods provided by stack implementations
type Stack interface {
	Push(int)
	Pop() (int, error)
}
