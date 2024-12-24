func rob(nums []int) int {
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    
    for i := 1; i < len(nums); i++ {
        if i == 1 {
            dp[i] = max(nums[1], dp[0])
        } else {
            dp[i] = max(dp[i-1], dp[i-2] + nums[i])
        }
    }

    return dp[len(dp)-1]
}