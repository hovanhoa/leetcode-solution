func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    for i := range dp {
        dp[i][0] = 1
    }

    for i := range dp[0] {
        dp[0][i] = 1
    }

    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[i]); j++ {
            dp[i][j] = dp[i][j-1] + dp[i-1][j]
        }
    }

    return dp[m-1][n-1]
}