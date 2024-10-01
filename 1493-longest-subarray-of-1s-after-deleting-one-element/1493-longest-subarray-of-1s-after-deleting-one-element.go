func longestSubarray(nums []int) int {
    l, k := 0, 0
    ans := 0
    for r := 0; r < len(nums); r++ {
        if nums[r] == 0 {
            k++
        }

        for k > 1 {
            if nums[l] == 0 {
                k--
            }

            l++
        }

        if r - l > ans {
            ans = r - l
        }
    }

    return ans
}