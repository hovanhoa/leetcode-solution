func maximalSquare(matrix [][]byte) int {
    m, n := len(matrix), len(matrix[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    ans := 0
    for i := range matrix {
        for j := range matrix[i] {
            if matrix[i][j] == '1' {
                left, top, diag := 0, 0, 0
                if i > 0 {
                    left = dp[i-1][j]
                } 

                if j > 0 {
                    top = dp[i][j-1]
                }

                if i > 0 && j > 0 {
                    diag = dp[i-1][j-1]
                }

                dp[i][j] = min(left, min(top, diag)) + 1
                ans = max(ans, dp[i][j])
            }
        }
    }

    return ans * ans
}