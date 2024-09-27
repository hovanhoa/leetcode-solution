func tribonacci(n int) int {
    if n < 1 {
        return 0
    }

    if n < 3 {
        return 1 
    }
    
    dp := make([]int, n+1)
    dp[1] = 1
    dp[2] = 1
    for i := 3; i < n+1; i++ {
        dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
    }

    return dp[n]
}