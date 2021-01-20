package avl

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

func (n *Node) rotateLeft() *Node {
	root := n.right
	n.right = root.left
	root.left = n
	return root
}
