func canJump(nums []int) bool {
    n := len(nums) - 1
    for i := n - 1; i >= 0; i-- {
        if nums[i] + i >= n {
            n = i
        }
    }

    return n == 0
}