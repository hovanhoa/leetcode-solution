/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    arr := []*TreeNode{root}
    ans := []int{}
    for len(arr) > 0 {
        ans = append(ans, arr[len(arr)-1].Val)
        newArr := make([]*TreeNode, 0)
        for _, v := range arr {
            if v.Left != nil {
                newArr = append(newArr, v.Left)
            }

            if v.Right != nil {
                newArr = append(newArr, v.Right)
            }
        }

        arr = newArr
    }

    return ans
}