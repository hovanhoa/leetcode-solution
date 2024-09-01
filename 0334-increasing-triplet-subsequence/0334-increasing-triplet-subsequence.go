func increasingTriplet(nums []int) bool {
    const MaxUint = ^uint(0) 
    const MaxInt = int(MaxUint >> 1) 

    first, second := MaxInt, MaxInt
    fmt.Println(first, second)
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