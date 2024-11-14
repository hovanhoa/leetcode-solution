func arrangeCoins(n int) int {
    ans, count := 0, 1
    for n > 0 {
        n -= count
        count += 1

        if n >= 0 {
            ans += 1
        }
    }

    return ans
}