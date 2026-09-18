/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	check := func(node *ListNode) bool {
		// 检查 node 后面有没有 k 个节点
		curr := node
		for i := 0; i < k; i++ {
			if curr.Next != nil {
				curr = curr.Next
			} else {
				return false
			}
		}
        return true
	}

	reverse := func(pre *ListNode) *ListNode {
		// 反转 pre 后面的 k 个节点
		curr := pre.Next
		futureTail := pre.Next
		for i := 0; i < k; i++ {
			n := curr.Next
			curr.Next = pre.Next
			pre.Next = curr
			curr = n
		}
		futureTail.Next = curr
		return futureTail
	}

	dummy := &ListNode{
		Next: head,
	}
	curr := dummy
	for check(curr) {
		curr = reverse(curr)
	}
	return dummy.Next
}