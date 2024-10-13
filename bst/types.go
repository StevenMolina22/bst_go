package bst

type BSTInterface[T comparable] interface {
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

type Color int

const (
	Red Color = iota
	Black
)

type Node[T comparable] struct {
	value T
	left  *Node[T]
	right *Node[T]
	color Color
}

type CmpFn[T comparable] func(a, b T) int

type BST[T comparable] struct {
	root *Node[T]
	cmp  CmpFn[T]
}

func NewBST[T comparable](f CmpFn[T]) *BST[T] {
	return &BST[T]{root: nil, cmp: f}
}
