func moveZeroes(nums []int)  {
    j := 0
    for i := range nums {
        if nums[i] != 0 {
            nums[j] = nums[i]
            j += 1
        }
    }

    for j < len(nums) {
        nums[j] = 0
        j += 1
    }
}