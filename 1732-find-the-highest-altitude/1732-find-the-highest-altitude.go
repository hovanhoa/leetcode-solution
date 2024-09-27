func largestAltitude(gain []int) int {
    ans := 0
    road := 0
    for _, v := range gain {
        road += v
        if ans < road {
            ans = road
        }
    }

    return ans
}