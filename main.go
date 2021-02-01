package main

import (
	"fmt"
	"github.com/ChristianSchleifer/datastructures/tree"
)

func main() {

	t := tree.NewBinarySearchTree()


	t.Insert(10)
	t.Insert(6)
	t.Insert(15)
	t.Insert(3)
	t.Insert(8)
	t.Insert(20)


	fmt.Println("Breadth first")

	t.TraverseBreadthFirst()

	fmt.Println("-------------------")

	fmt.Println("Pre Order")
	t.TraversePreOrder()

	fmt.Println("-------------------")

	fmt.Println("In Order")
	t.TraverseInOrder()

	fmt.Println("-------------------")

	fmt.Println("Post Order")
	t.TraversePostOrder()
}
