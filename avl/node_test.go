package avl

import (
	"testing"
)

type testElement struct {
	key   int
	value int
}

func (t *testElement) Less(e Element) int {
	et := e.(*testElement)
	if t.key < et.key {
		return -1
	} else if t.key > et.key {
		return 1
	}
	return 0
}

func TestNode_add(t *testing.T) {
	var n *Node
	n = n.add(&testElement{10, 10})
	n = n.add(&testElement{2, 2})
	n = n.add(&testElement{9, 9})
	n = n.add(&testElement{9, 9})
	n = n.add(&testElement{8, 8})

	e := n.getValue()
	et := e.(*testElement)
	expectedValue := 9
	expectedHeight := 3
	if expectedHeight != n.height {
		t.Errorf("expected height:%d, actual:%d", expectedHeight, n.height)
	}
	if expectedValue != et.value {
		t.Errorf("expected value:%d, actual:%d", expectedValue, et.value)
	}
}

func TestNode_remove(t *testing.T) {
	var n *Node
	n = n.add(&testElement{10, 10})
	n = n.add(&testElement{2, 2})
	n = n.add(&testElement{9, 9})
	n = n.add(&testElement{9, 9})
	n = n.add(&testElement{8, 8})
	n = n.remove(&testElement{9, 9})

	e := n.getValue()
	et := e.(*testElement)

	expectedValue := 8
	expectedHeight := 2
	if expectedHeight != n.height {
		t.Errorf("expected height:%d, actual:%d", expectedHeight, n.height)
	}
	if expectedValue != et.value {
		t.Errorf("expected value:%d, actual:%d", expectedValue, et.value)
	}
}
