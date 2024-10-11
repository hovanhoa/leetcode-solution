/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    node := head
    arr := []int{}
    for node != nil {
        arr = append(arr, node.Val)
        node = node.Next
    }

    ans := 0
    for i := 0; i < len(arr)/2; i++ {
        if arr[i] + arr[len(arr)-1-i] > ans {
            ans = arr[i] + arr[len(arr)-1-i]
        }
    }

    return ans
}