func productExceptSelf(nums []int) []int {
    left, right := make([]int, len(nums)), make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        if i == 0 {
            left[i] = 1
        } else {
            left[i] = left[i-1] * nums[i-1]
        }
    }

    for i := len(nums) - 1; i >= 0; i-- {
        if i == len(nums) - 1 {
            right[i] = 1
        } else {
            right[i] = right[i+1] * nums[i+1]
        }
    }

    ans := make([]int, len(nums))
    for i := range ans {
        ans[i] = left[i] * right[i] 
    }

    return ans
}