# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def leafSimilar(self, root1, root2):
        """
        :type root1: TreeNode
        :type root2: TreeNode
        :rtype: bool
        """
        leaves1, leaves2 = [], []

        self.getLeafs(root1, leaves1)
        self.getLeafs(root2, leaves2)

        return leaves1 == leaves2

    def getLeafs(self, root, leaves):
        if not root:
            return
        
        if not root.left and not root.right:
            leaves.append(root.val)
        
        self.getLeafs(root.left, leaves)
        self.getLeafs(root.right, leaves)
        