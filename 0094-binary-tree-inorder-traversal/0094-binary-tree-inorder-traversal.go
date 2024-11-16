/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    ans := []int{}
    var dfs func(root *TreeNode, arr []int)
    dfs = func(root *TreeNode, arr []int) {
        if root == nil {
            return
        }

        dfs(root.Left, ans)
        ans = append(ans, root.Val)
        dfs(root.Right, ans)
    }

    dfs(root, ans)

    return ans
}