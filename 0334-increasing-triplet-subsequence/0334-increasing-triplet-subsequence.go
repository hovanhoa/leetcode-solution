func increasingTriplet(nums []int) bool {
    if len(nums) < 3 {
        return false
    }

    f, s := math.MaxInt, math.MaxInt
    for _, v := range nums {
        if v <= f {
            f = v
        } else if v <= s {
            s = v
        } else {
            return true
        }
    }

    return false
}