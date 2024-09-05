func pivotIndex(nums []int) int {
    total := 0
    for _, val := range nums {
        total += val
    }

    leftSide := 0
    for i, val := range nums {
        if leftSide == total - leftSide - val {
            return i
        }

        leftSide += val
    }

    return -1
}