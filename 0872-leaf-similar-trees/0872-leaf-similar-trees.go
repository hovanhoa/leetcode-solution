/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getLeafs(root *TreeNode, arr *[]int) {
    if root == nil {
        return
    }

    if root.Left == nil && root.Right == nil {
        *arr = append(*arr, root.Val)
    }

    getLeafs(root.Left, arr) 
    getLeafs(root.Right, arr) 
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    arr1, arr2 := make([]int, 0), make([]int, 0)
    getLeafs(root1, &arr1)
    getLeafs(root2, &arr2)

    if len(arr1) != len(arr2) {
        return false
    }

    for i := 0; i < len(arr1); i++ {
        if arr1[i] != arr2[i] {
            return false
        }
    }

    return true
}