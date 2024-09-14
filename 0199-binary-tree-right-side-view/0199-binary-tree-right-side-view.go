/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }

    q := []*TreeNode{root}
    for len(q) > 0 {
        size := len(q)
        newQ := []*TreeNode{}
        for i := 0; i < size; i ++ {
            if i == size - 1 {
                res = append(res, q[i].Val)
            }

            if q[i].Left != nil {
                newQ = append(newQ, q[i].Left)
            }

            if q[i].Right != nil {
                newQ = append(newQ, q[i].Right)
            }
        }

        q = newQ
    }

    return res
}