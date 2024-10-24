func maxProfit(prices []int) int {
    ans := 0
    curMin := prices[0]
    for _, v := range prices {
        if v < curMin {
            curMin = v
        } else if v - curMin > ans {
            ans = v - curMin
        }
    }

    return ans
}