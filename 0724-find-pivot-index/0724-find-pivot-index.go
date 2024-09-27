func pivotIndex(nums []int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }

    curSum := 0
    for i, v := range nums {
        sum -= v
        if sum == curSum {
            return i
        }

        curSum += v
    } 

    return -1
}