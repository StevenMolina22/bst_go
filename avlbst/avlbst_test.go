package avlbst

import "testing"

func TestBST_Search(t *testing.T) {
	bst := NewBST[int]()
	if _, ok := bst.Search(1); ok {
		t.Error("Expected empty BST to return false")
	}
}

func TestBST_Insert(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(1)
	if _, ok := bst.Search(1); !ok {
		t.Error("Expected BST to contain 1")
	}
}

func TestBST_Remove(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(1)
	bst.Remove(0)
	if _, ok := bst.Search(1); !ok {
		t.Error("Expected BST to not contain 1")
	}
}

func TestBST_InsertMultiple(t *testing.T) {
	bst := NewBST[int]()
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	for _, v := range values {
		bst.Insert(v)
	}
	for _, v := range values {
		if _, ok := bst.Search(v); !ok {
			t.Errorf("Expected BST to contain %d", v)
		}
	}
}

func TestBST_InsertEdgeCases(t *testing.T) {
	bst := NewBST[int]()

	// Test inserting a very large number
	bst.Insert(1<<31 - 1)
	if _, ok := bst.Search(1<<31 - 1); !ok {
		t.Error("Expected BST to contain the maximum int value")
	}

	// Test inserting a very small number
	bst.Insert(-(1<<31 - 1))
	if _, ok := bst.Search(-(1<<31 - 1)); !ok {
		t.Error("Expected BST to contain the minimum int value")
	}

	// Test inserting zero
	bst.Insert(0)
	if _, ok := bst.Search(0); !ok {
		t.Error("Expected BST to contain 0")
	}

	// Test inserting duplicate values
	bst.Insert(1)
	bst.Insert(1)
	count := 0
	bst.InOrder(func(value int) {
		if value == 1 {
			count++
		}
	})
	if count != 2 {
		t.Error("Expected BST to contain 2")
	}
}

func TestBST_RemoveExisting(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(3)
	bst.Remove(2)

	if _, ok := bst.Search(2); ok {
		t.Error("Expected BST to not contain 2 after removal")
	}
	if _, ok := bst.Search(1); !ok {
		t.Error("Expected BST to still contain 1")
	}
	if _, ok := bst.Search(3); !ok {
		t.Error("Expected BST to still contain 3")
	}
}

func TestBST_Size(t *testing.T) {
	bst := NewBST[int]()
	if bst.Size() != 0 {
		t.Error("Expected size of empty BST to be 0")
	}

	bst.Insert(1)
	if bst.Size() != 1 {
		t.Error("Expected size of BST with one element to be 1")
	}

	bst.Insert(2)
	if bst.Size() != 2 {
		t.Error("Expected size of BST with two elements to be 2")
	}

	bst.Insert(0)
	if bst.Size() != 3 {
		t.Error("Expected size of BST with three elements to be 3")
	}

	bst.Remove(1)
	if bst.Size() != 2 {
		t.Error("Expected size of BST after removing one element to be 2")
	}
}

func TestBST_Height(t *testing.T) {
	bst := NewBST[int]()
	if bst.Height() != 0 {
		t.Error("Expected height of empty BST to be 0")
	}

	bst.Insert(1)
	if bst.Height() != 1 {
		t.Error("Expected height of BST with one element to be 1")
	}

	bst.Insert(2)
	if bst.Height() != 2 {
		t.Error("Expected height of BST with two elements to be 2")
	}

	bst.Insert(0)
	if bst.Height() != 2 {
		t.Error("Expected height of BST with three elements to be 2")
	}

	bst.Insert(-1)
	if bst.Height() != 3 {
		t.Error("Expected height of BST with four elements to be 3")
	}
}
