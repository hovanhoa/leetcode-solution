func longestSubarray(nums []int) int {
    l, res := 0, 0
    countZero := 0

    for r := 0; r < len(nums); r++ {
        if nums[r] == 0 {
            countZero++
        }

        for countZero == 2 {
            if nums[l] == 0 {
                countZero--
            }

            l++
        }

        if res < r - l {
            res = r - l
        }
    }

    return res
}