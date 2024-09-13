# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def longestZigZag(self, root: Optional[TreeNode]) -> int:
        maxLength = 0
        def dfs(node, direction, length):
            nonlocal maxLength
            if not node:
                return
            maxLength = max(maxLength, length)
            if direction == "left":
                dfs(node.left, "right", length + 1)
                dfs(node.right, "left", 1)
            else:
                dfs(node.right, "left", length + 1)
                dfs(node.left, "right", 1)
        if not root:
            return 0
        
        dfs(root, "left", 0)
        dfs(root, "right", 0)
        return maxLength
