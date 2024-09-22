func minDistance(word1 string, word2 string) int {
    dp := make([][]int, len(word1) + 1)
    for i := range dp {
        dp[i] = make([]int, len(word2) + 1)
        dp[i][0] = i
    }

    for i := range dp[0] {
        dp[0][i] = i
    }

    for i := 1; i <= len(word1); i++ {
        for j := 1; j <= len(word2); j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
            }
        }
    }

    return dp[len(word1)][len(word2)]
}

func min(nums ...int) int {
    ans := nums[0]
    for _, v := range nums {
        if v < ans {
            ans = v
        }
    }

    return ans
}