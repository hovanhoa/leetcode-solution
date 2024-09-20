func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 || j == 0 {
                dp[i][j] = 1
                continue
            }

            num := 0
            if i > 0 {
                num += dp[i-1][j] 
            }

            if j > 0 {
                num += dp[i][j-1] 
            }

            dp[i][j] = num
        }
    }

    return dp[m-1][n-1]
}