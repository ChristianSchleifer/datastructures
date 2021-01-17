package main

import (
	"fmt"
	"github.com/ChristianSchleifer/datastructures/list"
)

func main() {
	fmt.Println("Datastructures: ")

	fmt.Println("Singly linked list: ")

	sll := list.SinglyLinkedList{}

	sll.Push(1)
	sll.Push(2)
	sll.Push(3)
	sll.Push(4)
	sll.Push(5)

	fmt.Println(sll.Pop())
	fmt.Println(sll.Pop())
	fmt.Println(sll.Pop())
	fmt.Println(sll.Pop())
	fmt.Println(sll.Pop())
	fmt.Println(sll.Pop())
	fmt.Println(sll.Pop())
	fmt.Println(sll.Pop())
	fmt.Println(sll.Pop())

	sll.Unshift(5)
	sll.Unshift(4)
	sll.Unshift(3)
	sll.Unshift(2)
	sll.Unshift(1)

	fmt.Println(sll.Shift())
	fmt.Println(sll.Shift())
	fmt.Println(sll.Shift())
	fmt.Println(sll.Shift())
	fmt.Println(sll.Shift())
	fmt.Println(sll.Shift())
	fmt.Println(sll.Shift())
	fmt.Println(sll.Shift())
	fmt.Println(sll.Shift())
}
