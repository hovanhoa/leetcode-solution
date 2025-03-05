func rotate(nums []int, k int)  {
    n := k % len(nums)
    if n == 0 {
        return
    }

    
    copy(nums, append(nums[len(nums)-n:], nums[:len(nums)-n]...))
}