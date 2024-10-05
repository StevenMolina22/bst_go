package main

import (
	"fmt"

	"github.com/StevenMolina22/bst_go/bst"
)

func main() {

	// Create a new BST
	bst := bst.NewBST(func(a, b int) int {
		return a - b
	})

	// Insert elements
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println("Inserting values:", values)
	for _, v := range values {
		bst.Insert(v)
	}

	// Search for elements
	searchValues := []int{3, 7, 5}
	fmt.Println("Searching for values:", searchValues)
	for _, v := range searchValues {
		if _, ok := bst.Search(v); ok {
			fmt.Printf("Value %d found in BST\n", v)
		} else {
			fmt.Printf("Value %d not found in BST\n", v)
		}
	}

	// Remove elements
	removeValues := []int{1, 9, 5}
	fmt.Println("Removing values:", removeValues)
	for _, v := range removeValues {
		if ok := bst.Remove(v); ok {
			fmt.Printf("Value %d removed from BST\n", v)
		} else {
			fmt.Printf("Value %d not found in BST\n", v)
		}
	}

	// Perform in-order traversal
	fmt.Println("In-order traversal:")
	bst.InOrder(func(value int) {
		fmt.Print(value, " ")
	})
	fmt.Println()

	// Perform pre-order traversal
	fmt.Println("Pre-order traversal:")
	bst.PreOrder(func(value int) {
		fmt.Print(value, " ")
	})
	fmt.Println()

	// Perform post-order traversal
	fmt.Println("Post-order traversal:")
	bst.PostOrder(func(value int) {
		fmt.Print(value, " ")
	})
	fmt.Println()

	// Print size and height of the BST
	fmt.Printf("Size of BST: %d\n", bst.Size())
	fmt.Printf("Height of BST: %d\n", bst.Height())
}
