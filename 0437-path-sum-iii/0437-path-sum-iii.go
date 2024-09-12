/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }

    count := 0
    if targetSum == root.Val {
        count = 1
    }

    return count + dfs(root.Left, targetSum - root.Val) + dfs(root.Right, targetSum - root.Val)
}

func pathSum(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }

    return dfs(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}