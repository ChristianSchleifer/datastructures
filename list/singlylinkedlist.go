package list

type node struct {
	value int
	next  *node
}

func newNode(value int) *node {
	return &node{value: value}
}

type SinglyLinkedList struct {
	head *node
	tail *node

	length int
}

func (sll *SinglyLinkedList) isEmpty() bool {
	return sll.length == 0
}

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

func (sll *SinglyLinkedList) Pop() (int, bool) {

	pre := sll.head
	if pre == nil {
		return 0, false
	}

	pointer := sll.head.next
	if pointer == nil {
		val := sll.head.value
		sll.head = nil
		sll.tail = nil
		sll.length--

		return val, true
	}

	for pointer.next != nil {
		pre = pointer
		pointer = pointer.next
	}

	val := pointer.value
	pre.next = nil
	sll.tail = pre
	sll.length--

	return val, true
}

func (sll *SinglyLinkedList) Unshift(val int) {
	nextNode := newNode(val)

	nextNode.next = sll.head
	sll.head = nextNode

	if sll.isEmpty() {
		sll.tail = nextNode
	}

	sll.length++
}

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
