func canJump(nums []int) bool {
    n := len(nums) - 1
    for i := n; i >= 0; i-- {
        if i + nums[i] >= n {
            n = i
        }
    }

    return n == 0
}