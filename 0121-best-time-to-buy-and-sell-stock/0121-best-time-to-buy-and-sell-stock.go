func maxProfit(prices []int) int {
    res := 0
    m := prices[0]
    for _, v := range prices {
        m = min(m, v)
        res = max(res, v - m)
    }

    return res
}