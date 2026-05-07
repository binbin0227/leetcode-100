# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def removeNthFromEnd(self, head, n):
        """
        :type head: Optional[ListNode]
        :type n: int
        :rtype: Optional[ListNode]
        """
        # 快慢指针，使慢指针到达被删结点的上一个结点
        dummy = ListNode(0) # type: ignore
        dummy.next = head
        fast = slow = dummy
        for i in range(n + 1):
            fast = fast.next
        while fast:
            slow, fast = slow.next, fast.next
        slow.next = slow.next.next
        return dummy.next

    def removeNthFromEnd2(self, head, n):
        """
        :type head: Optional[ListNode]
        :type n: int
        :rtype: Optional[ListNode]
        """
        # 1. 设置空头结点
        dummy = ListNode(0)  # type: ignore
        dummy.next = head
        d = {}
        i = 0

        # 2. 从空头结点开始，给所有人发号码牌（存入字典）
        curr = dummy
        while curr:
            d[i] = curr
            i += 1
            curr = curr.next

        # 循环结束后，i 是队伍的总人数（包含了空头结点）。
        # 因为 i 在最后一次循环多加了 1，所以最后一个人的编号是 i - 1。
        # 倒数第 1 个人是 (i - 1) - 0
        # 倒数第 n 个人是 (i - 1) - (n - 1) = i - n
        # 我们要找的是被删人的【前一个人】，所以再往前退一步：i - n - 1

        prev_idx = i - n - 1

        # 3. 极其暴力的物理切断：前一个人直接连上下下个人
        # 因为 Python 会自动处理 next.next，就算被删的是最后一个节点，它的 next 也是 None，完美衔接！
        d[prev_idx].next = d[prev_idx].next.next

        # 4. 返回空头结点身后的真正队伍
        return dummy.next
