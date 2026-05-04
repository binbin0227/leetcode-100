# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution(object):
    def detectCycle(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        # 先用快慢指针确认是否为环形链表
        fast = slow = head
        while fast != None and fast.next != None:
            fast = fast.next.next
            slow = slow.next
            # 【从“链表起点”走到“环入口”的距离】 等于 【从“相遇点”继续往前走到“环入口”的距离】
            # 每次都只走 1 步，最终必定会在环的“入口节点”极其完美地迎面撞上！
            if fast == slow:
                slow = head
                while fast != slow:
                    fast = fast.next
                    slow = slow.next
                return fast
        return None
