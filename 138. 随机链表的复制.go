package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node, 0)

	curr := head
	for curr != nil {
		m[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	dummy := &Node{}
	p1, p2 := head, dummy
	for p1 != nil {

		p2.Next = m[p1]
		p2.Next.Val = p1.Val
		p2.Next.Next = m[p1.Next]
		p2.Next.Random = m[p1.Random]

		p1, p2 = p1.Next, p2.Next
	}

	return dummy.Next
}
