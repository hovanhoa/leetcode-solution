/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, targetSum int, total int) int {
    if root == nil {
        return 0
    }

    total += root.Val
    if total == targetSum {
        return 1 + dfs(root.Left, targetSum, total) + dfs(root.Right, targetSum, total) 
    }

    return dfs(root.Left, targetSum, total) + dfs(root.Right, targetSum, total) 
}

func pathSum(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }

    return dfs(root, targetSum, 0) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}