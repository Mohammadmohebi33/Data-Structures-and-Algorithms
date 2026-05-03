package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func createDoublyLinkedList(value int) *DoublyLinkedList {
	newNode := &Node{Value: value}
	return &DoublyLinkedList{
		head: newNode,
		tail: newNode,
		size: 1,
	}
}

func main() {
	newDoublyLinkedList := createDoublyLinkedList(1)
	fmt.Println(newDoublyLinkedList.head.Value)
}
