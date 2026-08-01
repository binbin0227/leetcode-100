package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	p1, p2 := dummy, head

	for {
		p1.Next = p2.Next
		p2.Next = p2.Next.Next
		p1.Next.Next = p2

		if p2.Next == nil || p2.Next.Next == nil {
			break
		}

		p1 = p2
		p2 = p2.Next
	}

	return dummy.Next
}
