func maxProfit(prices []int, fee int) int {
    dp1, dp2 := make([]int, len(prices)), make([]int, len(prices))
    dp2[0] = - prices[0]

    for i := 1; i < len(prices); i++ {
        dp1[i] = max(dp1[i-1], prices[i] + dp2[i-1] - fee)
        dp2[i] = max(dp2[i-1], dp1[i-1] - prices[i])
    }

    return dp1[len(prices) - 1]
}