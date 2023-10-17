# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def maxDepth(self, root):
        if not root:
            return 0
        return 1 + max(self.maxDepth(root.left), self.maxDepth(root.right))

    def isBalanced(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if not root:
            return True

        depthLeft = self.maxDepth(root.left)
        depthRight = self.maxDepth(root.right)
        res = abs(depthLeft - depthRight) <= 1
        return res and self.isBalanced(root.left) and self.isBalanced(root.right)
