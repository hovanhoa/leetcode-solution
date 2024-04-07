# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def goodNodes(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        return 1 + self.goodNodeWithMax(root.left, root.val) + self.goodNodeWithMax(root.right, root.val)
        
    def goodNodeWithMax(self, root, maxNum):
        if root is None:
            return 0
        
        count = 0
        if root.val >= maxNum:
            count = 1
            maxNum = root.val

        return count + self.goodNodeWithMax(root.left, maxNum) + self.goodNodeWithMax(root.right, maxNum)