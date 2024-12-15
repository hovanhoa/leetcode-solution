/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return isBST(root, nil, nil)
}

func isBST(root, min, max *TreeNode) bool {
    if root == nil {
        return true
    }

    if min != nil && min.Val >= root.Val {
        return false
    }

    if max != nil && max.Val <= root.Val {
        return false
    }

    return isBST(root.Left, min, root) && isBST(root.Right, root, max)
}