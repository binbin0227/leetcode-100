package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	reverse := func(pre, n2 *ListNode) {
		// 反转 (pre, n2]
		n1 := pre.Next
		curr := n1
		tail := n2.Next
		stop := n2.Next

		for {
			if curr == stop {
				break
			}

			n := curr.Next
			curr.Next = tail
			tail = curr
			curr = n
		}

		pre.Next = n2
	}

	dummy := &ListNode{Next: head}

	curr := dummy
	var pre, n2 *ListNode
	for i := 1; i <= right; i++ {
		if i == left {
			pre = curr
		}

		curr = curr.Next

		if i == right {
			n2 = curr
		}
	}

	reverse(pre, n2)
	return dummy.Next
}