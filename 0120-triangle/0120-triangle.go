func minimumTotal(triangle [][]int) int {
    dp := make([]int, len(triangle) + 1)
    for i := len(triangle) - 1; i >= 0; i-- {
        for j, v := range triangle[i] {
            dp[j] = v + min(dp[j], dp[j+1])
        }
    }

    return dp[0]
}