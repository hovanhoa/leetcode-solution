func productExceptSelf(nums []int) []int {
    countZero := 0
    product := 1
    for _, v := range nums {
        if v == 0 {
            countZero += 1
        } else {
            product *= v
        }
    }

    if countZero > 1 {
        return make([]int, len(nums))
    } else if countZero == 1 {
        for i := range nums {
            if nums[i] != 0 {
                nums[i] = 0
            } else {
                nums[i] = product
            }
        }

        return nums
    }

    for i := range nums {
        nums[i] = product / nums[i]
    }

    return nums
}