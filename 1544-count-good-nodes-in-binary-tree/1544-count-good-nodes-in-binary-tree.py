# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def goodNodes(self, root: TreeNode) -> int:
        return self.getGoodNodes(root, float("-inf"))
        
    def getGoodNodes(self, root: TreeNode, newMax: int) -> int:
        if not root:
            return 0
        
        newMax = max(newMax, root.val)
        return (1 if root.val >= newMax else 0) + self.getGoodNodes(root.left, newMax) + self.getGoodNodes(root.right, newMax)
    