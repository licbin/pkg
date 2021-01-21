package avl

// Element - element
type Element interface {
	Less(Element) int // -1 less , 0 equal , 1 greatear
}

// Node - avl node
type Node struct {
	e Element

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

func (n *Node) getValue() interface{} {
	return n.e
}

func (n *Node) search(e Element) *Node {
	if n == nil {
		return nil
	}
	rst := e.Less(n.e)
	switch rst {
	case 0:
		return n
	case -1:
		return n.left.search(e)
	case 1:
		return n.right.search(e)
	default:
		return nil
	}

}

// func (n *Node) displayNodesInOrder() {
// 	if n.left != nil {
// 		n.left.displayNodesInOrder()
// 	}
// 	fmt.Print(e.Value(), " ")
// 	if n.right != nil {
// 		n.right.displayNodesInOrder()
// 	}
// }

func (n *Node) remove(e Element) *Node {
	if n == nil {
		return nil
	}

	rst := e.Less(n.e)
	switch rst {
	case 0:
		if n.left != nil && n.right != nil {
			rightMinNode := n.right.findSmallest()
			n.e = rightMinNode.e
			n.right = n.right.remove(n.e)

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
	case -1:
		n.left = n.left.remove(e)
	case 1:
		n.right = n.right.remove(e)
	default:
		return nil
	}

	return n.rebalanceTree()
}

func (n *Node) findSmallest() *Node {
	if n.left != nil {
		return n.left.findSmallest()
	}
	return n
}

func (n *Node) add(e Element) *Node {
	if n == nil {
		return &Node{e, 1, nil, nil}
	}

	rst := e.Less(n.e)
	switch rst {
	case 0:
		n.e = e
	case -1:
		n.left = n.left.add(e)
	case 1:
		n.right = n.right.add(e)
	default:
		return nil
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
