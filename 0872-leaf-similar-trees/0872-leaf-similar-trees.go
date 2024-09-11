/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getLeaf(root *TreeNode, arr *[]int) {
    if root == nil {
        return
    }

    if root.Left == nil && root.Right == nil {
        *arr = append(*arr, root.Val)
    }

    getLeaf(root.Left, arr)
    getLeaf(root.Right, arr)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    arr1, arr2 := []int{}, []int{}
    getLeaf(root1, &arr1)
    getLeaf(root2, &arr2)

    fmt.Println(arr1, arr2)

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