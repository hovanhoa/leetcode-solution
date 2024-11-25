func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
    for i := range dp {
        dp[i] = amount + 1
    }  

    dp[0] = 0 
    for i := 1; i < amount + 1; i++ {
        for _, c := range coins {
            if i >= c {
                dp[i] = min(dp[i], dp[i-c] + 1)
            }
        }
    }

    if dp[amount] != amount + 1 {
        return dp[amount]
    }

    return -1
}