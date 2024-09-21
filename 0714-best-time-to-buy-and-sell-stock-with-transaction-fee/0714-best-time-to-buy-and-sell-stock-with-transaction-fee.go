func maxProfit(prices []int, fee int) int {
    dp := make([][2]int, len(prices))
    dp[0][1] = -prices[0]
    
    for i := 1; i < len(prices); i++ {
        dp[i][0] = max(dp[i-1][0], prices[i] + dp[i-1][1] - fee)
        dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
    }
    
    return dp[len(prices) - 1][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}