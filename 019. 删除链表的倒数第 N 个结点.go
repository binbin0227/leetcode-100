package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p1, p2 := dummy, dummy

	for range n {
		p2 = p2.Next
	}

	for p2.Next != nil {
		p1, p2 = p1.Next, p2.Next
	}

	p1.Next = p1.Next.Next
	return dummy.Next
}
