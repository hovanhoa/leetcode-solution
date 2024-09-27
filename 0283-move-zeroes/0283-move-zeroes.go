func moveZeroes(nums []int)  {
    count := 0
    for _, v := range nums {
        if v != 0 {
            nums[count] = v
            count++
        }
    }

    for count < len(nums) {
        nums[count] = 0
        count++
    }
}