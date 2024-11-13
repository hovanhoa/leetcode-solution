func pivotIndex(nums []int) int {
    right := 0
    for _, v := range nums {
        right += v
    }

    left := 0
    for i, v := range nums {
        right -= v
        if left == right {
            return i
        }

        left += v
    }

    return -1
}