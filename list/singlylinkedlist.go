package list

import "errors"

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

// Get returns the value of an element at a certain position
func (sll *SinglyLinkedList) Get(index int) (int, error) {
	node, err := sll.getNode(index)
	if err != nil {
		return 0, err
	}

	return node.value, nil
}

// Set changes the value of an element at a certain position
func (sll *SinglyLinkedList) Set(index int, val int) error {
	node, err := sll.getNode(index)

	if err != nil {
		return err
	}

	node.value = val
	return nil
}

// Insert adds an element at the specified position to the list
func (sll *SinglyLinkedList) Insert(index int, val int) error {
	if index < 0 || index > sll.length {
		return errors.New("index out of range")
	}

	if index == 0 {
		sll.Unshift(val)
		return nil
	}

	if index == sll.length {
		sll.Push(val)
		return nil
	}

	nodeToInsert := newNode(val)
	node, err := sll.getNode(index - 1)

	if err != nil {
		return err
	}

	nodeToInsert.next = node.next
	node.next = nodeToInsert
	sll.length++
	return nil
}

// Remove deletes an element at the specified position from the list
func (sll *SinglyLinkedList) Remove(index int) error {
	if index < 0 || index >= sll.length {
		return errors.New("index out of range")
	}

	if index == 0 {
		sll.Shift()
		return nil
	}

	if index == sll.length-1 {
		sll.Pop()
		return nil
	}

	node, err := sll.getNode(index - 1)
	if err != nil {
		return err
	}

	node.next = node.next.next
	sll.length--
	return nil
}

func (sll *SinglyLinkedList) isEmpty() bool {
	return sll.length == 0
}

func (sll *SinglyLinkedList) getNode(index int) (*node, error) {
	if index < 0 || index >= sll.length {
		return nil, errors.New("index out of range")
	}

	cursor := sll.head

	for i := 0; i < index; {
		cursor = cursor.next
		i++
	}

	return cursor, nil
}
