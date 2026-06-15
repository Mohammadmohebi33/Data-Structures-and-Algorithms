package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(value int) {
	if t.root == nil {
		t.root = &Node{value: value}
		return
	}
	t.root.insert(value)
}

func (t *Tree) Search(value int) bool {
	if t.root == nil {
		return false
	}
	return t.root.search(value)
}

func (n *Node) insert(value int) {
	if value < n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

func (n *Node) search(value int) bool {
	if n == nil {
		return false
	}
	if value == n.value {
		return true
	}
	if value < n.value {
		return n.left.search(value)
	}
	return n.right.search(value)
}

func main() {
	myTree := &Tree{}

	myTree.Insert(1)
	myTree.Insert(2)
	myTree.Insert(3)

	if myTree.Search(5) {
		fmt.Println("3 is exist")
	} else {
		fmt.Println("5 is not exist")
	}
}
