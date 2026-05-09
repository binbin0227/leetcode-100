# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def reverseKGroup(self, head, k):
        """
        :type head: Optional[ListNode]
        :type k: int
        :rtype: Optional[ListNode]
        """
        dummy = ListNode(0)  # type: ignore
        dummy.next = head

        def reverse(pre, k):
            """反转 pre 后面的 k 个结点"""
            # 当 k =3 时，pre -> 1 -> 2 -> 3 -> 4 ===> pre -> 3 -> 2 -> 1 -> 4
            head = pre.next
            for i in range(k - 1):
                curr = head.next
                head.next = curr.next
                curr.next = pre.next
                pre.next = curr
            return head  # 反转后的尾节点

        def judge(node, k):
            """判断node后面是否至少还有 k 个结点"""
            temp = node
            count = 0
            while temp.next and count < k:
                temp = temp.next
                count += 1
            return count == k

        curr = dummy
        while judge(curr, k):
            curr = reverse(curr, k)
        return dummy.next
