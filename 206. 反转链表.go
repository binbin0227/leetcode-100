/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{}
	curr := head

	for curr != nil {
		n := curr.Next
		curr.Next = dummy.Next
		dummy.Next = curr
		curr = n
	}

	return dummy.Next
}