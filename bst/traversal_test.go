package bst

import "testing"

// ----- TRAVERSAL -----
func TestBST_InOrder(t *testing.T) {
	bst := NewBST(intCmp)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(3)
	bst.InOrder(func(value int) {
		t.Log(value)
	})
}

func TestBST_Traversals(t *testing.T) {
	bst := NewBST(intCmp)
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	for _, v := range values {
		bst.Insert(v)
	}

	inOrder := bst.InOrderVec()
	expectedInOrder := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
	for i, v := range inOrder {
		if v != expectedInOrder[i] {
			t.Errorf("Expected in-order value %d, got %d", expectedInOrder[i], v)
		}
	}

	preOrder := bst.PreOrderVec()
	expectedPreOrder := []int{3, 1, 1, 2, 4, 5, 3, 5, 9, 6, 5}
	for i, v := range preOrder {
		if v != expectedPreOrder[i] {
			t.Errorf("Expected pre-order value %d, got %d", expectedPreOrder[i], v)
		}
	}

	postOrder := bst.PostOrderVec()
	expectedPostOrder := []int{1, 2, 1, 3, 5, 5, 5, 6, 9, 4, 3}
	for i, v := range postOrder {
		if v != expectedPostOrder[i] {
			t.Errorf("Expected post-order value %d, got %d", expectedPostOrder[i], v)
		}
	}
}
