package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	h := &ListNode{}

	p := head
	for p != nil {
		pNext := p.Next
		p.Next = h.Next
		h.Next = p
		p = pNext
	}

	return h.Next
}
