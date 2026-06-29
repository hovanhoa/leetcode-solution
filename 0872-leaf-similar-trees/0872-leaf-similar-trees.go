/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    arr1, arr2 := []int{}, []int{}
    getLeafs(root1, &arr1)
    getLeafs(root2, &arr2)

    if len(arr1) != len(arr2) {
        return false
    }

    for i := range arr1 {
        if arr1[i] != arr2[i] {
            return false
        }
    }

    return true
}

func getLeafs(root *TreeNode, arr *[]int) {
    if root == nil {
        return 
    }

    if root.Left == nil && root.Right == nil {
        *arr = append(*arr, root.Val)
        return
    }

    getLeafs(root.Left, arr)
    getLeafs(root.Right, arr)
}