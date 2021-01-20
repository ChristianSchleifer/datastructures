package list

type node struct {
	value int
	next  *node
}

func newNode(value int) *node {
	return &node{value: value}
}

// SinglyLinkedList is an implementation of a linked list data structure
type SinglyLinkedList struct {
	head *node
	tail *node

	length int
}

// Push adds a new element to the end of the list
func (sll *SinglyLinkedList) Push(val int) {

	nextNode := newNode(val)

	if sll.isEmpty() {
		sll.head = nextNode
	} else {
		sll.tail.next = nextNode
	}

	sll.tail = nextNode
	sll.length++
}

// Pop removes the last element of the list and returns it
func (sll *SinglyLinkedList) Pop() (int, bool) {

	pre := sll.head
	if pre == nil {
		return 0, false
	}

	cursor := sll.head.next
	if cursor == nil {
		val := sll.head.value
		sll.head = nil
		sll.tail = nil
		sll.length--

		return val, true
	}

	for cursor.next != nil {
		pre = cursor
		cursor = cursor.next
	}

	val := cursor.value
	pre.next = nil
	sll.tail = pre
	sll.length--

	return val, true
}

// Unshift adds a new element to the beginning of the list
func (sll *SinglyLinkedList) Unshift(val int) {
	nextNode := newNode(val)

	nextNode.next = sll.head
	sll.head = nextNode

	if sll.isEmpty() {
		sll.tail = nextNode
	}

	sll.length++
}

// Shift removes the element at the beginning of the list and returns it
func (sll *SinglyLinkedList) Shift() (int, bool) {

	if sll.isEmpty() {
		return 0, false
	}

	val := sll.head.value

	sll.head = sll.head.next
	sll.length--

	if sll.isEmpty() {
		sll.tail = nil
	}

	return val, true
}
