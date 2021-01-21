package avl

import "sync"

// Tree - av tree
type Tree struct {
	root *Node
	lock sync.RWMutex
}

// NewTree - new tree
func NewTree() *Tree {
	return &Tree{
		lock: sync.RWMutex{},
	}
}

// Add - add
func (t *Tree) Add(e Element) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.root = t.root.add(e)
}

// Search - search
func (t *Tree) Search(e Element) Element {
	t.lock.RLock()
	defer t.lock.RUnlock()
	if node := t.root.search(e); node != nil {
		return node.e
	}
	return nil
}

// Remove - remove
func (t *Tree) Remove(e Element) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.root = t.root.remove(e)
}
