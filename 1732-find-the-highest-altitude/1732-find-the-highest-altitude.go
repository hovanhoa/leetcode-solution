func largestAltitude(gain []int) int {
    res, cur := 0, 0
    for _, v := range gain {
        cur += v
        if res < cur {
            res = cur
        }
    }

    return res
}