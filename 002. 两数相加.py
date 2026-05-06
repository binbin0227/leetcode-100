# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: Optional[ListNode]
        :type l2: Optional[ListNode]
        :rtype: Optional[ListNode]
        """
        res = ListNode(0) # pyright: ignore[reportUndefinedVariable]
        curr = res
        p1 = l1
        p2 = l2
        carry = 0
        while p1 or p2 or carry:

            # 如果节点为空，就当它是 0，不影响加法
            val1 = p1.val if p1 else 0
            val2 = p2.val if p2 else 0
            total = val1 + val2 + carry

            # 算进位和余数
            carry = total // 10
            remain = total % 10

            # 创建新节点
            curr.next = ListNode(remain) # type: ignore
            curr = curr.next

            # p1 和 p2 往前走（注意要判断它们是不是已经为空了）
            if p1:
                p1 = p1.next
            if p2:
                p2 = p2.next

        return res.next

        # res = []
        # p1 = l1
        # p2 = l2
        # temp = 0
        # while p1 and p2:
        #     s = p1.val + p2.val + temp
        #     if s < 10:
        #         res.append(s)
        #         temp = 0
        #     else:
        #         res.append(s % 10)
        #         temp = 1
        #     p1,p2=p1.next,p2.next
        # if temp ==0:
        #     if p1 is None and p2 is None:
        #         return res
        #     elif p1:
        #         while p1:
        #             res.append(p1.val)
        #             p1=p1.next
        #         return res
        #     else:
        #         while p2:
        #             res.append(p2.val)
        #             p2=p2.next
        #         return res
        # if temp == 1:
        #     if p1 is None and p2 is None:
        #         res.append(1)
        #         return res
        #     elif p1:
        #         while p1:
        #             s = p1.val+temp
        #             if s<10:
        #                 res.append(s)
        #             else:
        #                 res.append(s%10)
        #                 temp=1
        #         return res
        #     else:
        #         while p2:
        #             s = p2.val+temp
        #             if s<10:
        #                 res.append(s)
        #             else:
        #                 res.append(s%10)
        #                 temp=1
        #         return res
