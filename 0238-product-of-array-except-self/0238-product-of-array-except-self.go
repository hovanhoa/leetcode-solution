func productExceptSelf(nums []int) []int {
    left := make([]int, len(nums))
    left[0] = 1
    for i := 1; i < len(nums); i++ {
        left[i] = left[i-1] * nums[i-1] 
    }

    right := make([]int, len(nums))
    right[len(nums)-1] = 1
    for i := len(nums) - 2; i >= 0; i-- {
        right[i] = right[i+1] * nums[i+1]
    }

    ans := make([]int, len(nums))
    for i := range ans {
        ans[i] = left[i] * right[i]
    }

    return ans
}