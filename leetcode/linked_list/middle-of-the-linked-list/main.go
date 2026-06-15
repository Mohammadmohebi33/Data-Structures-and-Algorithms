package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	first, second := head, head
	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next
	}
	return first
}

func main() {

}
