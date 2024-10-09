/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    if root == nil {
        return 0
    }

    max := root.Val
    ans, count := 1, 1
    q := []*TreeNode{root}
    for len(q) > 0 {
        newQ := make([]*TreeNode, 0)
        newMax := 0
        for _, v := range q {
            newMax += v.Val
            if v.Left != nil {
                newQ = append(newQ, v.Left)
            }

            if v.Right != nil {
                newQ = append(newQ, v.Right)
            }
        }

        if newMax > max {
            max = newMax
            ans = count
        }

        q = newQ
        count++
    }

    return ans
}