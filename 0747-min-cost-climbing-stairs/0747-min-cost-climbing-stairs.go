func minCostClimbingStairs(cost []int) int {
    dp := make([]int, len(cost) + 1)
    for i := range cost {
        dp[i+1] = cost[i]
    }

    for i := len(cost) - 2; i >= 0; i-- {
        dp[i] += min(dp[i+1], dp[i+2])
    }

    return dp[0]
}