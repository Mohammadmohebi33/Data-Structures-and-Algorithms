package main

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(value int) {

	t.root = insertRecursive(t.root, value)
}

func insertRecursive(node *Node, value int) *Node {

	if node == nil {
		return &Node{value: value}
	}

	if value < node.value {
		node.left = insertRecursive(node.left, value)
	} else if value > node.value {
		node.right = insertRecursive(node.right, value)
	}

	return node
}

func (t *Tree) Search(value int) bool {
	return searchRecursive(t.root, value)
}

func searchRecursive(node *Node, value int) bool {

	if node == nil {
		return false
	}

	if value == node.value {
		return true
	}

	if value < node.value {
		return searchRecursive(node.left, value)
	}

	return searchRecursive(node.right, value)
}

func (t *Tree) Delete(value int) {
	t.root = deleteRecursive(t.root, value)
}

func deleteRecursive(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.value {
		node.left = deleteRecursive(node.left, value)
	} else if value > node.value {
		node.right = deleteRecursive(node.right, value)
	} else {

		if node.left == nil {
			return node.right
		}

		if node.right == nil {
			return node.left
		}

		minNode := minValueNode(node.right)
		node.value = minNode.value
		node.right = deleteRecursive(node.right, minNode.value)
	}

	return node
}

func minValueNode(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

func main() {

}
