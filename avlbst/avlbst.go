package avlbst

func (bst *BST[T]) Insert(value T) {
	if bst == nil {
		return
	}
	new := &Node[T]{value: value}
	if bst.root == nil {
		bst.root = new
		return
	}
	bst.insert(bst.root, new)
}

func (bst *BST[T]) insert(node *Node[T], newNode *Node[T]) {
	if node == nil {
		return
	}
	if newNode.value > node.value {
		if node.right == nil {
			node.right = newNode
			return
		}
		bst.insert(node.right, newNode)
	} else {
		if node.left == nil {
			node.left = newNode
			return
		}
		bst.insert(node.left, newNode)
	}
}

func (bst *BST[T]) Remove(key T) (ok bool) {
	if bst.root == nil {
		return false
	}
	var removed bool
	bst.root, removed = bst.remove(bst.root, key)
	return removed
}

func (bst *BST[T]) remove(node *Node[T], key T) (*Node[T], bool) {
	if node == nil {
		return nil, false
	}

	if key > node.value { // key on right subtree
		var removed bool
		node.right, removed = bst.remove(node.right, key)
		return node, removed
	} else if key < node.value { // key on left subtree
		var removed bool
		node.left, removed = bst.remove(node.left, key)
		return node, removed
	} else {
		// Node to remove found
		if node.left == nil && node.right == nil { // no children
			return nil, true
		} else if node.left == nil { // right child only
			return node.right, true
		} else if node.right == nil { // left child only
			return node.left, true
		} else { // node has two children
			minNode := bst.findMin(node.right)
			node.value = minNode.value
			node.right, _ = bst.remove(node.right, minNode.value)
			return node, true
		}
	}
}

func (bst *BST[T]) findMin(node *Node[T]) *Node[T] {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (bst *BST[T]) Search(key T) (T, bool) {
	if bst == nil {
		return key, false
	}
	return bst.search(bst.root, key)
}

func (bst *BST[T]) search(node *Node[T], key T) (T, bool) {
	if node == nil {
		return key, false
	}
	if node.value == key {
		return node.value, true
	}
	// recursive cases
	if key > node.value {
		return bst.search(node.right, key)
	} else {
		return bst.search(node.left, key)
	}
}

func (bst *BST[T]) Size() uint {
	if bst == nil {
		return 0
	}
	return bst.size(bst.root)
}

func (bst *BST[T]) size(node *Node[T]) uint {
	if node == nil {
		return 0
	}
	return 1 + bst.size(node.left) + bst.size(node.right)
}

func (bst *BST[T]) Height() uint {
	if bst == nil {
		return 0
	}
	return bst.height(bst.root)
}

func (bst *BST[T]) height(node *Node[T]) uint {
	if node == nil {
		return 0
	}
	max := func(a, b uint) uint {
		if a > b {
			return a
		}
		return b
	}
	return 1 + max(bst.height(node.right), bst.height(node.left))
}
