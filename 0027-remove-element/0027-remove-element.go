func removeElement(nums []int, val int) int {
    indexCount := 0
    for i, value := range nums {
        if val != value {
            nums[indexCount], nums[i] = nums[i], nums[indexCount]
            indexCount ++
        }
    }

    return indexCount
}