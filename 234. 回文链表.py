# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def isPalindrome(self, head):
        """
        :type head: Optional[ListNode]
        :rtype: bool
        """
    # 用快慢指针找到中点，翻转后半段，再用双指针判断两个链表的数据是否相等
        def reverseList(head):
            curr = head
            pre = None
            while curr:
                n = curr.next
                curr.next = pre
                pre = curr
                curr = n
            return pre

        fast = slow = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
        
        p2 = reverseList(slow)
        p1 = head
        while p2:
            if p1.val != p2.val:
                return False
            p1, p2 = p1.next, p2.next
        return True
