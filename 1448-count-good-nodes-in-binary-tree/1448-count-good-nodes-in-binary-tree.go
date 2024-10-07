/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodesWithVal(root *TreeNode, val int) int {
    if root == nil {
        return 0
    }

    if root.Val >= val {
        return 1 + goodNodesWithVal(root.Left, root.Val) + goodNodesWithVal(root.Right, root.Val)
    }

    return goodNodesWithVal(root.Left, val) + goodNodesWithVal(root.Right, val)
}

func goodNodes(root *TreeNode) int {
    return goodNodesWithVal(root, root.Val)
}