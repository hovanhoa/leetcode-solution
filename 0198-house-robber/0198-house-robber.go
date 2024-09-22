func rob(nums []int) int {
    if len(nums) < 3 {
        return max(nums...)
    }

    dp := make([]int, len(nums))
    dp[0], dp[1] = nums[0], max(nums[0], nums[1])
    for i := 2; i < len(nums); i++ {
        dp[i] = max(dp[i-1], nums[i] + dp[i-2])
    }

    return dp[len(nums)-1]
}

func max(nums ...int) int {
    ans := nums[0]
    for _, v := range nums {
        if ans < v {
            ans = v
        }
    }

    return ans
}