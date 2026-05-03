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
	newNode := &Node{Value: value, Next: nil, Prev: nil}
	return &DoublyLinkedList{
		head: newNode,
		tail: newNode,
		size: 1,
	}
}

func (l *DoublyLinkedList) Append(value int) bool {
	newNode := &Node{Value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.size = 1
		return true
	}
	newNode.Prev = l.tail
	l.tail.Next = newNode
	l.tail = newNode
	l.size++
	return true
}

func (l *DoublyLinkedList) Prepend(value int) bool {
	newNode := &Node{Value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.size = 1
		return true
	}
	l.head.Prev = newNode
	newNode.Next = l.head
	l.head = newNode
	l.size++
	return true
}

func (l *DoublyLinkedList) PopLast() bool {
	if l.size == 0 {
		return false
	}
	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0
		return true
	}
	l.tail = l.tail.Prev
	l.tail.Next = nil
	l.size--
	return true
}

func (l *DoublyLinkedList) PopFirst() bool {
	if l.size == 0 {
		return false
	}
	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0
		return true
	}
	l.head = l.head.Next
	l.head.Prev = nil
	l.size--
	return true
}

func (l *DoublyLinkedList) Get(index int) *Node {
	if index < 0 || index >= l.size {
		return nil
	}
	var temp *Node
	if index < l.size/2 {
		temp = l.head
		for i := 0; i < index; i++ {
			temp = temp.Next
		}
	} else {
		temp = l.tail
		for i := l.size - 1; i > index; i-- {
			temp = temp.Prev
		}
	}
	return temp
}

func (l *DoublyLinkedList) Set(index int, value int) bool {
	temp := l.Get(index)
	if temp == nil {
		return false
	}
	temp.Value = value
	return true
}

func (l *DoublyLinkedList) Insert(index int, value int) bool {
	if index < 0 || index > l.size {
		return false
	}
	newNode := &Node{Value: value}

	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
		l.size = 1
		return true
	}

	if index == 0 {
		l.Prepend(value)
		return true
	}

	if index == l.size {
		l.Append(value)
		return true
	}

	current := l.Get(index)
	newNode.Prev = current.Prev
	newNode.Next = current
	current.Prev.Next = newNode
	current.Prev = newNode
	l.size++
	return true
}

func (l *DoublyLinkedList) Delete(index int) bool {
	if index < 0 || index >= l.size {
		return false
	}
	if index == 0 {
		return l.PopFirst()
	}
	if index == l.size-1 {
		return l.PopLast()
	}
	current := l.Get(index)
	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev
	l.size--
	return true
}

func (l *DoublyLinkedList) Print() {
	temp := l.head
	for temp != nil {
		fmt.Println(temp.Value)
		temp = temp.Next
	}
}

func (l *DoublyLinkedList) PrintReverse() {
	temp := l.tail
	for temp != nil {
		fmt.Println(temp.Value)
		temp = temp.Prev
	}
}

func main() {
	newDoublyLinkedList := createDoublyLinkedList(1)
	newDoublyLinkedList.Append(2)
	newDoublyLinkedList.Append(3)
	newDoublyLinkedList.Append(4)
	newDoublyLinkedList.Insert(1, 100)
	newDoublyLinkedList.Delete(4)
	newDoublyLinkedList.Print()
}
