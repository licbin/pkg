package avl

import (
	"testing"
)

func TestTree_Add(t *testing.T) {
	var tr *Tree
	tr = NewTree()
	tr.Add(&testElement{10, 10})
	tr.Add(&testElement{2, 2})
	tr.Add(&testElement{9, 9})
	tr.Add(&testElement{9, 9})
	tr.Add(&testElement{8, 8})

	e := tr.root.getValue()
	et := e.(*testElement)
	expectedValue := 9
	expectedHeight := 3
	if expectedHeight != tr.root.height {
		t.Errorf("expected height:%d, actual:%d", expectedHeight, tr.root.height)
	}
	if expectedValue != et.value {
		t.Errorf("expected value:%d, actual:%d", expectedValue, et.value)
	}
}

func TestTree_Remove(t *testing.T) {
	var tr *Tree
	tr = NewTree()
	tr.Add(&testElement{10, 10})
	tr.Add(&testElement{2, 2})
	tr.Add(&testElement{9, 9})
	tr.Add(&testElement{9, 9})
	tr.Add(&testElement{8, 8})
	tr.Remove(&testElement{9, 9})

	e := tr.root.getValue()
	et := e.(*testElement)
	expectedValue := 8
	expectedHeight := 2
	if expectedHeight != tr.root.height {
		t.Errorf("expected height:%d, actual:%d", expectedHeight, tr.root.height)
	}
	if expectedValue != et.value {
		t.Errorf("expected value:%d, actual:%d", expectedValue, et.value)
	}

}
