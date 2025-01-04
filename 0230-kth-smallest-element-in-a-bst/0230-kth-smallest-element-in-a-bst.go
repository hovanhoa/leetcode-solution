/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    cur := root
    n := 0
    stack := []*TreeNode{}

    for cur != nil || len(stack) > 0 {
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }

        cur = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        n += 1
        if n == k {
            return cur.Val
        }

        cur = cur.Right
    }

    return -1
}