func moveZeroes(nums []int)  {
    index := 0
    for i, val := range nums {
        if val != 0 {
            nums[index], nums[i] = nums[i], nums[index]
            index ++
        }
    }
}