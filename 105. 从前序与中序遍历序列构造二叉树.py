# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def buildTree(self, preorder, inorder):
        """
        :type preorder: List[int]
        :type inorder: List[int]
        :rtype: Optional[TreeNode]
        """
        if not preorder or not inorder:
            return None

        # 前序名单第一位是根节点
        root_val = preorder[0]
        root = TreeNode(root_val)  # type: ignore

        # 把切割好的左右两份名单，分别交给左右手去递归建树
        mid_index = inorder.index(root_val)
        root.left = self.buildTree(preorder[1 : mid_index + 1], inorder[:mid_index])
        root.right = self.buildTree(preorder[mid_index + 1 :], inorder[mid_index + 1 :])

        return root
