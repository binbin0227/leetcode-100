"""
# Definition for a Node.
class Node:
    def __init__(self, x, next=None, random=None):
        self.val = int(x)
        self.next = next
        self.random = random
"""


class Solution(object):
    def copyRandomList(self, head):
        """
        :type head: Node
        :rtype: Node
        """
        if not head:  # 处理空链表
            return None
        # =========================================
        # 第一步：创造影分身 (A -> A' -> B -> B' -> C -> C')
        # =========================================
        curr = head
        while curr:
            clone = Node(curr.val)  # type: ignore
            clone.next = curr.next
            curr.next = clone
            curr = clone.next
        # =========================================
        # 第二步：灵魂注入 (复制 random 指针)
        # =========================================
        curr = head
        while curr:
            if curr.random:
                # A'.random = A.random.next (即 C')
                curr.next.random = curr.random.next
            curr = curr.next.next
        # =========================================
        # 第三步：物理剥离 (拉链分离)
        # =========================================
        curr = head
        clone_head = head.next
        while curr:
            clone = curr.next
            curr.next = curr.next.next  # 1. 恢复老链表：A 绕过 A' 指向 B
            clone.next = (
                clone.next.next if clone.next else None
            )  # 2. 提取新链表：A' 绕过 B 指向 B' (要注意判断 B 是不是 None)
            curr = curr.next  # 3. 履带前推：走到 B 继续拆
        return clone_head

    # 字典哈希表 (空间复杂度 O(N))
    def copyRandomList2(self, head):
        """
        :type head: Node
        :rtype: Node
        """
        dict = {}

        # 第一遍：造肉身
        curr = head
        while curr:
            dict[curr] = Node(curr.val)  # 不复制 next 和 random # type: ignore
            curr = curr.next

        # 第二遍：装灵魂（连线）
        curr = head
        while curr:
            clone = dict[curr]

            # 通过旧节点的 next/random 去字典里查到对应的新节点
            clone.next = dict.get(curr.next)
            clone.random = dict.get(curr.random)  # .get()找不到不会报错

            curr = curr.next
        return dict.get(head)
