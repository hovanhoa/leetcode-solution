/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    root := &TreeNode{Val: preorder[0]}
    i := slices.Index(inorder, preorder[0])
    root.Left = buildTree(preorder[1:1+i], inorder[:i])
    root.Right = buildTree(preorder[1+i:], inorder[i+1:])
    return root
}