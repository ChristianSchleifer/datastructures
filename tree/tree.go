package tree

type Tree interface {
	Insert(int)
	Find(int) bool
	TraverseBreadthFirst()
	TraversePreOrder()
	TraverseInOrder()
	TraversePostOrder()
}
