package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// 找中点，反转后半段，一一比对
	l, r := head, head
	for r.Next != nil && r.Next.Next != nil {
		r = r.Next.Next
		l = l.Next
	}

	l.Next = reverse(l.Next)

	l, r = head, l.Next
	for r != nil {
		if l.Val != r.Val {
			return false
		}
		l, r = l.Next, r.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
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
