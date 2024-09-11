/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    arr := []int{}
    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }

    l, r := 0, len(arr) - 1
    max := 0
    for l < r {
        if arr[l] + arr[r] > max {
            max = arr[l] + arr[r]
        }
        l++
        r--
    }

    return max
}