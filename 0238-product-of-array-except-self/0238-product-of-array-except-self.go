func productExceptSelf(nums []int) []int {
    ans := make([]int, len(nums))
    for i := range ans {
        ans[i] = 1
    }

    prefix := 1
    for i := range nums {
        ans[i] = prefix
        prefix *= nums[i]
    }

    postfix := 1
    for i := len(nums) - 1; i >= 0; i-- {
        ans[i] *= postfix
        postfix *= nums[i]
    }

    return ans
}