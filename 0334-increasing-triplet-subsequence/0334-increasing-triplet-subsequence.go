func increasingTriplet(nums []int) bool {
    first, second := math.MaxInt32, math.MaxInt32
    for _, v := range nums {
        if v <= first {
            first = v
        } else if v <= second {
            second = v
        } else {
            return true
        }
    }
    
    return false
}