package avl

import (
	"testing"
)

func TestNode_add(t *testing.T) {
	var n *Node
	n = n.add(10, 10)
	n = n.add(2, 2)
	n = n.add(9, 9)
	n = n.add(9, 9)
	n = n.add(8, 8)

	expectedValue := 9
	expectedHeight := 3
	if expectedHeight != n.height {
		t.Errorf("expected height:%d, actual:%d", expectedHeight, n.height)
	}
	if expectedValue != n.value {
		t.Errorf("expected value:%d, actual:%d", expectedValue, n.value)
	}
}

func TestNode_remove(t *testing.T) {
	var n *Node
	n = n.add(10, 10)
	n = n.add(2, 2)
	n = n.add(9, 9)
	n = n.add(9, 9)
	n = n.add(8, 8)

	n = n.remove(9)

	expectedValue := 8
	expectedHeight := 2
	if expectedHeight != n.height {
		t.Errorf("expected height:%d, actual:%d", expectedHeight, n.height)
	}
	if expectedValue != n.value {
		t.Errorf("expected value:%d, actual:%d", expectedValue, n.value)
	}
}
