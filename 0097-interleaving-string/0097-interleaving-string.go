func isInterleave(s1 string, s2 string, s3 string) bool {
    m, n, l := len(s1), len(s2), len(s3)
    if m + n != l {
        return false
    }

    dp := make([][]bool, m + 1)
    for i := range dp {
        dp[i] = make([]bool, n + 1)
    }
    dp[0][0] = true

    for i := 1; i <= m; i++ {
        dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
    }

    for i := 1; i <= n; i++ {
        dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
        }
    }

    return dp[m][n]
}