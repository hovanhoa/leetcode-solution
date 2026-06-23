func longestSubarray(nums []int) int {
    return longestOnes(nums, 1) - 1
}

func longestOnes(nums []int, k int) int {
    ans := 0
    l, r := 0, 0
    cnt := 0
    for r < len(nums) {
        if nums[r] == 0 {
            cnt += 1
            for cnt > k {
                if nums[l] == 0 {
                    cnt -= 1
                }

                l += 1
            }
        }

        r += 1

        ans = max(ans, r - l)
    }

    return ans
}