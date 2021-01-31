package list

import "errors"

type singlyLinkedNode struct {
	value interface{}
	next  *singlyLinkedNode
}

func newNode(value interface{}) *singlyLinkedNode {
	return &singlyLinkedNode{value: value}
}

type singlyLinkedList struct {
	head *singlyLinkedNode
	tail *singlyLinkedNode

	length int
}

// NewSinglyLinkedList returns a singly linked list based implementation of a list
func NewSinglyLinkedList() List {
	return &singlyLinkedList{}
}

// Push adds a new element to the end of the list
func (list *singlyLinkedList) Push(val interface{}) {

	nextNode := newNode(val)

	if list.isEmpty() {
		list.head = nextNode
	} else {
		list.tail.next = nextNode
	}

	list.tail = nextNode
	list.length++
}

// Pop removes the last element of the list and returns it
func (list *singlyLinkedList) Pop() (interface{}, error) {

	if list.isEmpty() {
		return 0, errors.New("list is empty")
	}

	pre := list.head
	cursor := list.head.next

	if cursor == nil {
		val := list.head.value
		list.head = nil
		list.tail = nil
		list.length--

		return val, nil
	}

	for cursor.next != nil {
		pre = cursor
		cursor = cursor.next
	}

	val := cursor.value
	pre.next = nil
	list.tail = pre
	list.length--

	return val, nil
}

// Unshift adds a new element to the beginning of the list
func (list *singlyLinkedList) Unshift(val interface{}) {
	nextNode := newNode(val)

	nextNode.next = list.head
	list.head = nextNode

	if list.isEmpty() {
		list.tail = nextNode
	}

	list.length++
}

// Shift removes the element at the beginning of the list and returns it
func (list *singlyLinkedList) Shift() (interface{}, error) {

	if list.isEmpty() {
		return 0, errors.New("list is empty")
	}

	val := list.head.value

	list.head = list.head.next
	list.length--

	if list.isEmpty() {
		list.tail = nil
	}

	return val, nil
}

// Get returns the value of an element at a certain position
func (list *singlyLinkedList) Get(index int) (interface{}, error) {
	node, err := list.getNode(index)
	if err != nil {
		return 0, err
	}

	return node.value, nil
}

// Set changes the value of an element at a certain position
func (list *singlyLinkedList) Set(index int, val interface{}) error {
	node, err := list.getNode(index)

	if err != nil {
		return err
	}

	node.value = val
	return nil
}

// Insert adds an element at the specified position to the list
func (list *singlyLinkedList) Insert(index int, val interface{}) error {
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

	nodeToInsert := newNode(val)
	node, err := list.getNode(index - 1)

	if err != nil {
		return err
	}

	nodeToInsert.next = node.next
	node.next = nodeToInsert
	list.length++
	return nil
}

// Remove deletes an element at the specified position from the list
func (list *singlyLinkedList) Remove(index int) error {
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

	node, err := list.getNode(index - 1)
	if err != nil {
		return err
	}

	node.next = node.next.next
	list.length--
	return nil
}

// Reverse inverts the singly linked list in place (i.e. not making a copy)
func (list *singlyLinkedList) Reverse() {
	if list.length <= 1 {
		return
	}

	var tmp *singlyLinkedNode
	var previousNode *singlyLinkedNode
	currentNode := list.head

	for i := 0; i < list.length; i++ {
		tmp = currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = tmp
	}

	tmp = list.head
	list.head = list.tail
	list.tail = tmp
}

func (list *singlyLinkedList) isEmpty() bool {
	return list.length == 0
}

func (list *singlyLinkedList) getNode(index int) (*singlyLinkedNode, error) {
	if index < 0 || index >= list.length {
		return nil, errors.New("index out of range")
	}

	cursor := list.head

	for i := 0; i < index; {
		cursor = cursor.next
		i++
	}

	return cursor, nil
}
