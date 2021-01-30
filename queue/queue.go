package queue

// Queue is the interface describing methods provided by queue implementations
type Queue interface {
	Enqueue(int)
	Dequeue() (int, error)
}
