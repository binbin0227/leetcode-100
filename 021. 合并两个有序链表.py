# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def mergeTwoLists(self, list1, list2):
        """
        :type list1: Optional[ListNode]
        :type list2: Optional[ListNode]
        :rtype: Optional[ListNode]
        """
        p1, p2 = list1, list2
        head = ListNode() # pyright: ignore[reportUndefinedVariable]
        p = head
        while p1 != None and p2 != None:
            if p1.val < p2.val:
                p.next = p1
                p1 = p1.next
                p = p.next
            elif p1.val == p2.val:
                p.next = p1
                p1 = p1.next  # 没说去重，p2不用走
                p = p.next
            else:
                p.next = p2
                p2 = p2.next
                p = p.next

        if p1 != None:  # p1有剩
            p.next = p1
        else:  # p2有剩
            p.next = p2
        return head.next
