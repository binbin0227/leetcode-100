# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[Optional[ListNode]]
        :rtype: Optional[ListNode]
        """

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

        k = len(lists)
        if k == 0:
            return None
        if k == 1:
            return lists[0]

        result = lists[0]
        t = 2  # 表示即将进行前t个链表的合并
        while t <= k:
            result = merge(result, lists[t - 1])
            t += 1
        return result
