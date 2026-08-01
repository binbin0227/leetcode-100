package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	flag := 0
	dummy := &ListNode{}
	curr := dummy
	sum := 0

	for p1 != nil && p2 != nil {
		sum = p1.Val + p2.Val + flag
		if sum >= 10 {
			flag = 1
			sum = sum % 10
		} else {
			flag = 0
		}
		curr.Next = &ListNode{Val: sum}
		curr = curr.Next
		p1 = p1.Next
		p2 = p2.Next
	}

	if flag == 1 {
		if p1 == nil && p2 == nil {
			curr.Next = &ListNode{Val: 1}
		} else if p1 == nil {
			for p2 != nil {
				sum = p2.Val + flag
				if sum >= 10 {
					flag = 1
					sum = sum % 10
				} else {
					flag = 0
				}
				curr.Next = &ListNode{Val: sum}
				curr = curr.Next
				p2 = p2.Next
			}
		} else {
			for p1 != nil {
				sum = p1.Val + flag
				if sum >= 10 {
					flag = 1
					sum = sum % 10
				} else {
					flag = 0
				}
				curr.Next = &ListNode{Val: sum}
				curr = curr.Next
				p1 = p1.Next
			}
		}
	} else {
		if p1 == nil && p2 == nil {
			return dummy.Next
		} else if p1 == nil {
			for p2 != nil {
				sum = p2.Val + flag
				if sum >= 10 {
					flag = 1
					sum = sum % 10
				} else {
					flag = 0
				}
				curr.Next = &ListNode{Val: sum}
				curr = curr.Next
				p2 = p2.Next
			}
		} else {
			for p1 != nil {
				sum = p1.Val + flag
				if sum >= 10 {
					flag = 1
					sum = sum % 10
				} else {
					flag = 0
				}
				curr.Next = &ListNode{Val: sum}
				curr = curr.Next
				p1 = p1.Next
			}
		}
	}
	if flag == 1 {
		curr.Next = &ListNode{Val: 1}
	}
	return dummy.Next
}
