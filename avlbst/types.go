package avlbst

import "cmp"

type BSTInterface[T cmp.Ordered] interface {
	Insert(value T)
	Remove(value T) (ok bool)
	Search(value T) (T, ok bool)
	Size() uint
	Height() uint

	InOrder(f func(T))
	PreOrder(f func(T))
	PostOrder(f func(T))

	InOrderVec() []T
	PreOrderVec() []T
	PostOrderVec() []T
}

type Node[T cmp.Ordered] struct {
	value  T
	parent *Node[T]
	left   *Node[T]
	right  *Node[T]
	height uint
}

type BST[T cmp.Ordered] struct {
	root *Node[T]
}

func NewBST[T cmp.Ordered]() *BST[T] {
	return &BST[T]{root: nil}
}
