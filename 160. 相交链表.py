# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution(object):
    def getIntersectionNode(self, headA, headB):
        """
        :type head1, head1: ListNode
        :rtype: ListNode
        """
        # pa 走完 链表A 后去走 链表B ，pb 走完 链表B 后去走 链表A 
        # 有相交节点 -> return 循环结束时的 pa
        # 无相交节点 -> return None
        pa, pb = headA, headB
        while pa != pb:
            if pa is None:
                pa = headB
            else:
                pa = pa.next
            if pb is None:
                pb = headA
            else:
                pb = pb.next
        return pa

    # 超时解法
    def getIntersectionNode2(self, headA, headB):
        """
        :type head1, head1: ListNode
        :rtype: ListNode
        """
        pa = headA
        while pa is not None:
            pb = headB
            while pb is not None:
                if pa == pb:
                    return pa
                pb = pb.next
            pa = pa.next
