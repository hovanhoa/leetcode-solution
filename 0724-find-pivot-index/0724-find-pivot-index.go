func pivotIndex(nums []int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }

    cur := 0
    for i, v := range nums {
        sum -= v
        if cur == sum {
            return i
        }

        cur += v
    }

    return -1
}