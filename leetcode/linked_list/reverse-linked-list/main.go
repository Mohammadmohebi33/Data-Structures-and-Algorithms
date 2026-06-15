package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {

	for {
		swapped := false

		curr := head

		for curr != nil && curr.Next != nil {
			if curr.Val > curr.Next.Val {
				curr.Val, curr.Next.Val =
					curr.Next.Val, curr.Val

				swapped = true
			}

			curr = curr.Next
		}

		if !swapped {
			break
		}
	}
	return head
}

func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// پیدا کردن وسط
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// نصف کردن لیست
	right := slow.Next
	slow.Next = nil

	// سورت بازگشتی دو نیمه
	left := sortList(head)
	right = sortList(right)

	// ادغام دو لیست مرتب
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}

		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}

	if l2 != nil {
		curr.Next = l2
	}

	return dummy.Next
}
func main() {

}
