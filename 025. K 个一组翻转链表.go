package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	curr := dummy

	for check(curr, k) {
		curr = reverse(curr, k)
	}

	return dummy.Next
}

func reverse(n *ListNode, k int) *ListNode {
	first, curr := n.Next, n.Next
	dummy := &ListNode{}
	var next *ListNode
	for range k {
		next = curr.Next
		curr.Next = dummy.Next
		dummy.Next = curr
		curr = next
	}
	first.Next = next
	n.Next = dummy.Next
	return first
}

func check(n *ListNode, k int) bool {
	if n == nil {
		return false
	}

	for range k {
		if n.Next != nil {
			n = n.Next
			continue
		}
		return false
	}

	return true
}
