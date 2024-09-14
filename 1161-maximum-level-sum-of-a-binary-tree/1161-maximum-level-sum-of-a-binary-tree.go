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
    maxLevel, level := 1, 1
    q := []*TreeNode{root}
    for len(q) > 0 {
        newQ := []*TreeNode{}
        sum := 0
        for i := 0; i < len(q); i++ {
            sum += q[i].Val
            if q[i].Left != nil {
                newQ = append(newQ, q[i].Left)
            }

            if q[i].Right != nil {
                newQ = append(newQ, q[i].Right)
            }
        }

        if sum > max {
            maxLevel = level
            max = sum
        }

        level++
        q = newQ
    }

    return maxLevel
}