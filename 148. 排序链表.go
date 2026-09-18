package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow.Next
	slow.Next = nil

	leftSorted := sortList(head)
	rightSorted := sortList(mid)

	dummy := &ListNode{}
	curr := dummy
	p1, p2 := leftSorted, rightSorted
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			curr.Next = p1
			p1 = p1.Next
		} else {
			curr.Next = p2
			p2 = p2.Next
		}
		curr = curr.Next
	}
	if p1 != nil {
		curr.Next = p1
	} else {
		curr.Next = p2
	}
	return dummy.Next
}