package main

import "fmt"

type Node struct {
	Value    int
	Level    int
	Children []*Node
}

type Tree struct {
	Root *Node
}

func NewTree(value int) *Tree {
	return &Tree{
		Root: &Node{
			Value: value,
			Level: 0,
		},
	}
}

func (n *Node) Add(value int) *Node {
	child := &Node{
		Value: value,
		Level: n.Level + 1,
	}

	n.Children = append(n.Children, child)
	return child
}

func (t *Tree) Print() {
	printNode(t.Root)
}

func printNode(node *Node) {
	if node == nil {
		return
	}

	for i := 0; i < node.Level; i++ {
		fmt.Print("  ")
	}

	fmt.Printf("%d (level=%d)\n", node.Value, node.Level)

	for _, child := range node.Children {
		printNode(child)
	}
}

func (n *Node) Search(value int) *Node {

	if n == nil {
		return nil
	}

	if n.Value == value {
		return n
	}

	for _, child := range n.Children {
		if result := child.Search(value); result != nil {
			return result
		}
	}
	return nil
}

func (t *Tree) Search(value int) *Node {
	return t.Root.Search(value)
}

func (t *Tree) Delete(value int) bool {
	if t.Root == nil {
		return false
	}

	if t.Root.Value == value {
		t.Root = nil
		return true
	}

	return deleteNode(t.Root, value)
}

func deleteNode(parent *Node, value int) bool {
	for i, child := range parent.Children {
		if child.Value == value {
			// delete from slice
			parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
			return true
		}

		if deleteNode(child, value) {
			return true
		}
	}

	return false
}

func main() {
	tree := NewTree(1)

	n2 := tree.Root.Add(2)
	n3 := tree.Root.Add(3)
	tree.Root.Add(4)

	n2.Add(5)
	n2.Add(6)

	n3.Add(7)

	tree.Print()
}
