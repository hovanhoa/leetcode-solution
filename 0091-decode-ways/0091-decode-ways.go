func numDecodings(s string) int {
    if s[0] == '0' {
        return 0
    }

    size := len(s)
    dp := make([]int, size + 1)
    dp[0], dp[1] = 1, 1

    for i := 2; i <= size; i++ {
        sub := s[i-2:i]
        if "10" <= sub && sub <= "26" {
            dp[i] = dp[i-2]
        }

        if s[i-1] != '0' {
            dp[i] += dp[i-1]
        }
    }

    return dp[size]
}