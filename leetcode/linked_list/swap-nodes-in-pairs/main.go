package main

import "fmt"

type LinkedList struct {
	head *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *LinkedList) swapPairs(head *ListNode) {

	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		// swap
		prev.Next = second
		first.Next = second.Next
		second.Next = first

		// move forward
		prev = first
		head = first.Next
	}

	for dummy != nil {
		fmt.Println(dummy.Val)
		dummy = dummy.Next
	}

}

func createLinkedList(value int) *LinkedList {
	newNode := &ListNode{Val: value}
	return &LinkedList{
		head: newNode,
	}
}

func (l *LinkedList) Prepend(value int) bool {
	newNode := &ListNode{Val: value, Next: nil}

	newNode.Next = l.head
	l.head = newNode
	return true
}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func main() {
	mylist := createLinkedList(4)
	mylist.Prepend(3)
	mylist.Prepend(2)
	mylist.Prepend(1)

	//	mylist.Print()

	mylist.swapPairs(mylist.head)

}
