# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def swapPairs(self, head):
        """
        :type head: Optional[ListNode]
        :rtype: Optional[ListNode]
        """
        dummy = ListNode(0) # type: ignore
        dummy.next = head
        pre = dummy
        
        # 只要 pre 后面还有完整的两个节点，就说明可以凑成一对进行交换
        while pre.next and pre.next.next:
            # 在循环内部安全地定义 a 和 b，绝对不会报空指针错误
            a = pre.next
            b = pre.next.next
            
            # 你的核心交换逻辑（一模一样，完美运行）
            a.next = b.next
            b.next = a
            pre.next = b
            
            # 关键修复：为下一轮交换做准备
            # 交换后变成了 pre -> b -> a -> 下一对
            # 所以下一轮的前驱节点应该是 a！
            pre = a
            
        return dummy.next

    # 冗余解法
    def swapPairs2(self, head):
        """
        :type head: Optional[ListNode]
        :rtype: Optional[ListNode]
        """
        dummy = ListNode(0)  # type: ignore
        dummy.next = head
        pre = dummy
        a = dummy.next if dummy else None
        b = a.next if a else None
        while b:
            a.next = b.next
            b.next = a
            pre.next = b
            pre = a
            a = pre.next if pre else None
            b = a.next if a else None
        return dummy.next
