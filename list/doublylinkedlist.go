package list

import "errors"

type doublyLinkedNode struct {
	value int

	next *doublyLinkedNode
	prev *doublyLinkedNode
}

func newDoublyLinkedNode(val int) *doublyLinkedNode {
	return &doublyLinkedNode{value: val}
}

type DoublyLinkedList struct {
	head *doublyLinkedNode
	tail *doublyLinkedNode

	length int
}

// Push adds an element to the end of the list
func (list *DoublyLinkedList) Push(val int) {
	newNode := newDoublyLinkedNode(val)

	if list.isEmpty() {
		list.head = newNode
		list.tail = newNode
		list.length++
		return
	}

	list.tail.next = newNode
	newNode.prev = list.tail
	list.tail = newNode
	list.length++
}

// Pop removes an element from the end of the list and returns it
func (list *DoublyLinkedList) Pop() (int, error) {
	if list.isEmpty() {
		return 0, errors.New("list is empty")
	}

	if list.length == 1 {
		val := list.tail.value
		list.head = nil
		list.tail = nil
		list.length--
		return val, nil
	}

	val := list.tail.value
	list.tail = list.tail.prev
	list.tail.next = nil
	list.length--

	return val, nil
}

// Shift removes an element from the beginning of the list and returns it
func (list *DoublyLinkedList) Shift() (int, error) {
	if list.isEmpty() {
		return 0, errors.New("list is empty")
	}

	if list.length == 1 {
		val := list.tail.value
		list.head = nil
		list.tail = nil
		list.length--
		return val, nil
	}

	val := list.head.value
	list.head = list.head.next
	list.head.prev = nil
	list.length--
	return val, nil
}

// Unshift adds an element to the beginning of the list
func (list *DoublyLinkedList) Unshift(val int) {
	newNode := newDoublyLinkedNode(val)

	if list.isEmpty() {
		list.head = newNode
		list.tail = newNode
		list.length++
		return
	}

	newNode.next = list.head
	list.head.prev = newNode
	list.head = newNode
	list.length++
}

// Get returns the value of an element at a certain position
func (list *DoublyLinkedList) Get(index int) (int, error) {
	node, err := list.getNode(index)

	if err != nil {
		return 0, err
	}

	return node.value, nil
}

// Set changes the value of an element at a certain position
func (list *DoublyLinkedList) Set(index int, val int) error {
	node, err := list.getNode(index)

	if err != nil {
		return err
	}

	node.value = val
	return nil
}

// Insert adds an element at the specified position to the list
func (list *DoublyLinkedList) Insert(index int, val int) error {
	if index < 0 || index > list.length {
		return errors.New("index out of range")
	}

	if index == 0 {
		list.Unshift(val)
		return nil
	}

	if index == list.length {
		list.Push(val)
		return nil
	}

	newNode := newDoublyLinkedNode(val)
	prevNode, err := list.getNode(index - 1)
	if err != nil {
		return err
	}

	newNode.next = prevNode.next
	prevNode.next.prev = newNode
	newNode.prev = prevNode
	prevNode.next = newNode
	list.length++
	return nil
}

// Remove deletes an element at the specified position from the list
func (list *DoublyLinkedList) Remove(index int) error {
	if index < 0 || index >= list.length {
		return errors.New("index out of range")
	}

	if index == 0 {
		_, err := list.Shift()
		return err
	}

	if index == list.length-1 {
		_, err := list.Pop()
		return err
	}

	prevNode, err := list.getNode(index - 1)
	if err != nil {
		return err
	}

	prevNode.next = prevNode.next.next
	prevNode.next.prev = prevNode
	list.length--
	return nil
}

func (list *DoublyLinkedList) getNode(index int) (*doublyLinkedNode, error) {
	if index < 0 || index >= list.length {
		return nil, errors.New("index out of range")
	}

	// search from the end or beginning of the list depending on the index
	if index < list.length/2 {
		cursor := list.head
		for i := 0; i < index; i++ {
			cursor = cursor.next
		}

		return cursor, nil
	}

	cursor := list.tail
	for i := list.length - 1; i > index; i-- {
		cursor = cursor.prev
	}

	return cursor, nil
}

func (list *DoublyLinkedList) isEmpty() bool {
	return list.length == 0
}