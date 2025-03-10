func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    return max(findRob(nums[1:]), findRob(nums[:len(nums)-1]))
}

func findRob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    dp := make([]int, len(nums))
    dp[0], dp[1] = nums[0], max(nums[0], nums[1])
    for i := 2; i < len(nums); i++ {
        dp[i] = max(dp[i-1], dp[i-2] + nums[i])
    }

    return dp[len(nums)-1]
}