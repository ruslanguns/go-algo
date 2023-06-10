package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) insert(key int) {
	newNode := &Node{key: key, left: nil, right: nil}

	if bst.root == nil {
		bst.root = newNode
	} else {
		bst.insertNode(bst.root, newNode)
	}
}

func (bst *BinarySearchTree) insertNode(node *Node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			bst.insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			bst.insertNode(node.right, newNode)
		}
	}
}

func (bst *BinarySearchTree) search(key int) bool {
	return bst.searchNode(bst.root, key)
}

func (bst *BinarySearchTree) searchNode(node *Node, key int) bool {
	if node == nil {
		return false
	}

	if key == node.key {
		return true
	} else if key < node.key {
		return bst.searchNode(node.left, key)
	} else {
		return bst.searchNode(node.right, key)
	}
}

func (bst *BinarySearchTree) remove(key int) {
	bst.root = bst.removeNode(bst.root, key)
}

func (bst *BinarySearchTree) removeNode(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = bst.removeNode(node.left, key)
	} else if key > node.key {
		node.right = bst.removeNode(node.right, key)
	} else {
		// Case 1: No child or only one child
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		// Case 2: Two children
		minNode := bst.findMinNode(node.right)
		node.key = minNode.key
		node.right = bst.removeNode(node.right, minNode.key)
	}

	return node
}

func (bst *BinarySearchTree) findMinNode(node *Node) *Node {
	if node == nil || node.left == nil {
		return node
	}

	return bst.findMinNode(node.left)
}

func (bst *BinarySearchTree) inOrderTraversal() {
	bst.inOrder(bst.root)
	fmt.Println()
}

func (bst *BinarySearchTree) inOrder(node *Node) {
	if node != nil {
		bst.inOrder(node.left)
		fmt.Printf("%d ", node.key)
		bst.inOrder(node.right)
	}
}

func main() {
	bst := BinarySearchTree{}

	// Insert nodes
	bst.insert(50)
	bst.insert(30)
	bst.insert(20)
	bst.insert(40)
	bst.insert(70)
	bst.insert(60)
	bst.insert(80)

	// Perform in-order traversal
	fmt.Println("In-order Traversal:")
	bst.inOrderTraversal()

	// Search for a key
	key := 40
	found := bst.search(key)
	if found {
		fmt.Printf("Key %d found in the Binary Search Tree\n", key)
	} else {
		fmt.Printf("Key %d not found in the Binary Search Tree\n", key)
	}

	// Remove a key
	keyToRemove := 30
	bst.remove(keyToRemove)
	fmt.Printf("Key %d removed from the Binary Search Tree\n", keyToRemove)

	// Perform in-order traversal after removal
	fmt.Println("In-order Traversal after Removal:")
	bst.inOrderTraversal()
}
