func maxProfit(prices []int) int {
    curPrice := prices[0]
    ans := 0
    for i := 1; i < len(prices); i++ {
        if curPrice > prices[i] {
            curPrice = prices[i]
        } else {
            ans += prices[i] - curPrice
            curPrice = prices[i]
        }
    }

    return ans
}