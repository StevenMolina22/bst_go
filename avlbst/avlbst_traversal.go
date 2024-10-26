package avlbst

func (bst *BST[T]) InOrder(f func(T)) {
	if bst == nil {
		return
	}
	bst.inOrder(f, bst.root)
}

func (bst *BST[T]) inOrder(f func(T), node *Node[T]) {
	if node == nil {
		return
	}
	bst.inOrder(f, node.left)
	f(node.value)
	bst.inOrder(f, node.right)
}

func (bst *BST[T]) PreOrder(f func(T)) {
	if bst == nil {
		return
	}
	bst.preOrder(f, bst.root)
}

func (bst *BST[T]) preOrder(f func(T), node *Node[T]) {
	if node == nil {
		return
	}
	f(node.value)
	bst.preOrder(f, node.left)
	bst.preOrder(f, node.right)
}

func (bst *BST[T]) PostOrder(f func(T)) {
	if bst == nil {
		return
	}
	bst.postOrder(f, bst.root)
}

func (bst *BST[T]) postOrder(f func(T), node *Node[T]) {
	if node == nil {
		return
	}
	bst.postOrder(f, node.left)
	bst.postOrder(f, node.right)
	f(node.value)
}

func (bst *BST[T]) InOrderVec() []T {
	if bst == nil {
		return nil
	}
	return nil
}

func (bst *BST[T]) PreOrderVec() []T {
	if bst == nil {
		return nil
	}
	return nil
}

func (bst *BST[T]) PostOrderVec() []T {
	if bst == nil {
		return nil
	}
	return nil
}
