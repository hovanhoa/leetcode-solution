# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def leafSimilar(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> bool:
        def getLeafs(root):
            if not root:
                return
            if not root.left and not root.right:
                yield root.val
                return
            yield from getLeafs(root.left)
            yield from getLeafs(root.right)

        return list(getLeafs(root1)) == list(getLeafs(root2))  