func productExceptSelf(nums []int) []int {
    leftArr := make([]int, len(nums))
    rightArr := make([]int, len(nums))
    for i := 0; i < len(nums); i ++ {
        if i == 0 {
            leftArr[i] = 1
        } else {
            leftArr[i] = leftArr[i-1] * nums[i-1]
        }
    }

    for i := len(nums) - 1; i >= 0; i -- {
        if i == len(nums) - 1 {
            rightArr[i] = 1
        } else {
            rightArr[i] = rightArr[i+1] * nums[i+1]
        }
    }

    for i := 0; i < len(nums); i ++ {
        leftArr[i] = leftArr[i] * rightArr[i]
    }

    return leftArr
}