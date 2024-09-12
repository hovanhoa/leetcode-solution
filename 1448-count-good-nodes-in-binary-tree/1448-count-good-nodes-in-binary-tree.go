/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func godNodes(root *TreeNode, max int) int {
    if root == nil {
        return 0
    }

    count := 0
    if root.Val >= max {
        count = 1
        max = root.Val
    }

    return count + godNodes(root.Left, max) + godNodes(root.Right, max)
}

func goodNodes(root *TreeNode) int {
    return godNodes(root, root.Val)
}