package tree

import (
	"fmt"
	"github.com/ChristianSchleifer/datastructures/queue"
)

type node struct {
	val int

	left  *node
	right *node
}

func newNode(val int) *node {
	return &node{val: val}
}

type binarySearchTree struct {
	root *node
}

func NewBinarySearchTree() Tree {
	return &binarySearchTree{}
}

func (t *binarySearchTree) Insert(val int) {
	newChild := newNode(val)

	if t.root == nil {
		t.root = newChild
		return
	}

	parent := traverse(newChild.val, t.root)

	if newChild.val < parent.val {
		parent.left = newChild
	} else if newChild.val > parent.val {
		parent.right = newChild
	}
}

func (t *binarySearchTree) Find(val int) bool {
	if t.root == nil {
		return false
	}

	node := traverse(val, t.root)

	if node != nil && node.val == val {
		return true
	}

	return false
}

func traverse(val int, node *node) *node {
	if val < node.val && node.left != nil {
		return traverse(val, node.left)
	}

	if val > node.val && node.right != nil {
		return traverse(val, node.right)
	}

	return node
}

func (t *binarySearchTree) TraverseBreadthFirst() {

	q := queue.NewLinkedListQueue()
	q.Enqueue(t.root)

	for {
		current, err := q.Dequeue()
		if err != nil {
			break
		}

		node, ok := current.(*node)
		if !ok {
			panic("encountered a unknown element in the queue")
		}

		fmt.Println(node.val)

		if node.left != nil {
			q.Enqueue(node.left)
		}

		if node.right != nil {
			q.Enqueue(node.right)
		}
	}
}

func (t *binarySearchTree) TraversePreOrder() {
	if t.root == nil {
		return
	}

	traversePreOrder(t.root)
}

func traversePreOrder(node *node) {
	fmt.Println(node.val)

	if node.left != nil {
		traversePreOrder(node.left)
	}

	if node.right != nil {
		traversePreOrder(node.right)
	}
}

func (t *binarySearchTree) TraverseInOrder() {
	if t.root == nil {
		return
	}

	traverseInOrder(t.root)
}

func traverseInOrder(node *node) {
	if node.left != nil {
		traverseInOrder(node.left)
	}

	fmt.Println(node.val)

	if node.right != nil {
		traverseInOrder(node.right)
	}
}

func (t *binarySearchTree) TraversePostOrder() {
	if t.root == nil {
		return
	}

	traversePostOrder(t.root)
}

func traversePostOrder(node *node) {
	if node.left != nil {
		traversePostOrder(node.left)
	}

	if node.right != nil {
		traversePostOrder(node.right)
	}

	fmt.Println(node.val)
}
