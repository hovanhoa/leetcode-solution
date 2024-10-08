/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(root *TreeNode, isLeft bool, count int) int {
    if root == nil {
        return count
    }

    if isLeft {
        return max(dfs(root.Right, false, count+1), dfs(root.Left, true, 0))
    }

    return max(dfs(root.Left, true, count+1), dfs(root.Right, false, 0))
}

func longestZigZag(root *TreeNode) int {
    return max(dfs(root.Left, true, 0), dfs(root.Right, false, 0))
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}