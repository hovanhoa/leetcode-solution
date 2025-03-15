func maxProfit(prices []int) int {
    res := 0
    cur := prices[0]
    for _, v := range prices {
        if v < cur {
            cur = v
        } else {
            res += (v- cur)
            cur = v
        }
    }

    return res
}