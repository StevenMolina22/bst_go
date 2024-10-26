package avlbst

/*
Example of Left Rotation:

	 A
	  \
	   B
	  / \
	T1   C

After left rotation around node A:

	  B
	 / \
	A   C
	 \
	 T1
*/
func (bst *BST[T]) RotateLeft(node *Node[T]) {
	B := node.right
	T1 := B.left
	B.left = node
	node.right = T1
}

/*
Example of Left Rotation:

	  B
	 / \
	A	C
	 \
	  T2

After left rotation around node A:

	A
	 \
	  B
	 / \
	T2  C
*/
func (bst *BST[T]) RotateRight(node *Node[T]) {
	A := node.left
	T2 := A.right
	A.right = node
	node.left = T2
}
