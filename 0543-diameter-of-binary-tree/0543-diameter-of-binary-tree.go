/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    ans := 0
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        left, right := dfs(node.Left), dfs(node.Right)
        ans = max(ans, left+right)
        return 1 + max(left, right)
    }

    dfs(root)
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}