func largestAltitude(gain []int) int {
    ans := 0
    cur := 0
    for _, v := range gain {
        cur += v
        ans = max(ans, cur)
    }

    return ans
}