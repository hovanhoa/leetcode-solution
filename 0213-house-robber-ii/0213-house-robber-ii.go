func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    
    return max(findRob(nums[1:]), findRob(nums[:len(nums)-1]))
}

func findRob(nums []int) int {
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        if i == 1 {
            dp[i] = max(dp[0], nums[i])
        } else {
            dp[i] = max(dp[i-2] + nums[i], dp[i-1])
        }
    }

    return dp[len(nums)-1]
}