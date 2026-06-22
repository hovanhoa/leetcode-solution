func productExceptSelf(nums []int) []int {
    left, right := make([]int, len(nums)), make([]int, len(nums))
    leftProduct, rightProduct := 1, 1
    for i, v := range nums {
        left[i] = leftProduct
        leftProduct *= v
    }

    for i := len(nums) - 1; i >= 0; i -- {
        right[i] = rightProduct
        rightProduct *= nums[i]
    }

    ans := make([]int, len(nums))
    for i := range nums {
        ans[i] = left[i] * right[i]
    }

    return ans
}