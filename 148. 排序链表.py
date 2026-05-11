# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def sortList(self, head):
        """
        :type head: Optional[ListNode]
        :rtype: Optional[ListNode]
        """

        def cutMiddle(head):
            """返回列表另一半的开头"""
            if not head:
                return None
            if not head.next:
                return head
            fast = slow = head
            while fast.next and fast.next.next:
                fast = fast.next.next
                slow = slow.next
            mid = slow.next
            slow.next = None  # 切断链表
            return mid

        def merge(l1, l2):
            """将两个链表中的元素有序连接"""
            dummy = ListNode(0)  # type: ignore
            curr = dummy
            while l1 and l2:
                if l1.val <= l2.val:
                    curr.next = l1
                    l1 = l1.next
                else:
                    curr.next = l2
                    l2 = l2.next
                curr = curr.next
            if l1:
                curr.next = l1
            else:
                curr.next = l2

            return dummy.next

        if not head or not head.next:
            return head
        mid = cutMiddle(head)

        # 让左右两边各自排序
        left_sorted = self.sortList(head)
        right_sorted = self.sortList(mid)
        return merge(left_sorted, right_sorted)
