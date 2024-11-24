func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
    for i := range dp {
        dp[i] = amount + 1
    }
    dp[0] = 0

    for n := 1; n < amount + 1; n++ {
        for _, c := range coins {
            if n >= c {
                dp[n] = min(dp[n], 1 + dp[n-c])
            }
        }

    }

    if dp[amount] != amount + 1 {
        return dp[amount]
    }

    return -1
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}