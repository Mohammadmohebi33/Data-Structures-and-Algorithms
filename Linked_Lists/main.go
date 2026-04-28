package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

func createLinkedList(value int) *LinkedList {
	newNode := &Node{value: value}
	return &LinkedList{
		head: newNode,
		tail: newNode,
		len:  1,
	}
}

func (l *LinkedList) Prepend(value int) bool {
	newNode := &Node{value: value, next: nil}
	if l.len == 0 {
		l.head = newNode
		l.tail = newNode
		return true
	}
	newNode.next = l.head
	l.head = newNode
	l.len++
	return true
}

func (l *LinkedList) Append(value int) bool {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return true
	}
	l.tail.next = newNode
	l.tail = newNode
	l.len++
	return true
}

func (l *LinkedList) PopFirst() bool {
	if l.len == 0 {
		return false
	}
	l.head = l.head.next
	l.len--
	return true
}

func (l *LinkedList) PopLast() bool {
	if l.len == 0 {
		return false
	}
	temp, pre := l.head, l.head
	for temp.next != nil {
		pre = temp
		temp = temp.next
	}
	pre.next = nil
	l.len--
	return true
}

func (l *LinkedList) Get(index int) *Node {
	if l.len < index || index < 0 {
		return &Node{value: -1}
	}
	temp := l.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp
}

func (l *LinkedList) Set(index int, value int) bool {
	temp := l.Get(index)
	if temp == nil {
		return false
	}
	temp.value = value
	l.len++
	return true
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) Insert(index int, value int) bool {
	if l.len < index || index < 0 {
		return false
	}
	if index == 0 {
		return l.Prepend(value)
	}
	if index == l.len {
		return l.Append(value)
	}
	newNode := &Node{value: value}
	temp := l.Get(index - 1)
	newNode.next = temp.next
	temp.next = newNode
	l.len++
	return true
}

func (l *LinkedList) Remove(index int) bool {
	if l.len < index || index < 0 {
		return false
	}
	if index == 0 {
		return l.PopFirst()
	}
	if index == l.len {
		return l.PopLast()
	}
	prev := l.Get(index - 1)
	temp := prev.next
	prev.next = temp.next
	temp.next = nil
	l.len--
	return true
}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func main() {
	my_node := createLinkedList(10)

	my_node.Prepend(40)
	my_node.Append(50)
	my_node.Append(100)
	my_node.Append(200)

	my_node.Print()
	fmt.Println("--------")

	my_node.PopLast()
	my_node.Print()

	//my_node.Prepend(30)
	//my_node.Prepend(33)
	//
	//my_node.Append(100)
	//
	////	my_node.Pop()
	////my_node.Print()
	//my_node.PopFirst()
	//my_node.Print()

	//fmt.Println(my_node.Get(1))
	//my_node.Set(1, 150)
	//fmt.Println("--------")
	//my_node.Print()
	//fmt.Println("--------")
	//my_node.Insert(1, 44)
	//my_node.Print()
	//fmt.Println("--------")
	//my_node.Remove(1)
	//my_node.Print()
}
