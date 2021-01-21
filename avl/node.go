package avl

import (
	"fmt"
)

// Interface - element
type Interface interface {
	Value() interface{}
	Less(Interface) bool
}

// Tree - av tree
type Tree struct {
}

// Node - avl node
type Node struct {
	key   int
	value int

	height int
	left   *Node
	right  *Node
}

func (n *Node) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node) getValue() int {
	return n.value
}

func (n *Node) search(key int) *Node {
	if n == nil {
		return nil
	}
	if key < n.key {
		return n.left.search(key)
	} else if key > n.key {
		return n.right.search(key)
	} else {
		return n
	}
}

func (n *Node) displayNodesInOrder() {
	if n.left != nil {
		n.left.displayNodesInOrder()
	}
	fmt.Print(n.key, " ")
	if n.right != nil {
		n.right.displayNodesInOrder()
	}
}

func (n *Node) remove(key int) *Node {
	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = n.left.remove(key)
	} else if key > n.key {
		n.right = n.right.remove(key)
	} else {
		if n.left != nil && n.right != nil {
			rightMinNode := n.right.findSmallest()
			n.key = rightMinNode.key
			n.value = rightMinNode.value
			n.right = n.right.remove(rightMinNode.key)

		} else if n.left != nil {
			// node only has left child
			n = n.left
		} else if n.right != nil {
			// node only has right child
			n = n.right
		} else {
			// node has no children
			n = nil
			return n
		}
	}

	return n.rebalanceTree()
}

func (n *Node) findSmallest() *Node {
	if n.left != nil {
		return n.left.findSmallest()
	}
	return n
}

func (n *Node) add(key, value int) *Node {
	if n == nil {
		return &Node{key, value, 1, nil, nil}
	}

	if key < n.key {
		n.left = n.left.add(key, value)
	} else if key > n.key {
		n.right = n.right.add(key, value)
	} else {
		n.value = value
	}

	return n.rebalanceTree()
}

func (n *Node) rebalanceTree() *Node {
	if n == nil {
		return n
	}

	n.recalculateHeight()

	balanceFactor := n.left.getHeight() - n.right.getHeight()
	if balanceFactor == -2 {
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor == 2 {
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}

	return n
}

func (n *Node) rotateLeft() *Node {
	root := n.right
	n.right = root.left
	root.left = n

	n.recalculateHeight()
	root.recalculateHeight()
	return root
}

func (n *Node) rotateRight() *Node {
	root := n.left
	n.left = root.right
	root.right = n

	n.recalculateHeight()
	root.recalculateHeight()
	return root
}

func (n *Node) recalculateHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

// Returns max number
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
