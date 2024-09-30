func longestOnes(nums []int, k int) int {
    l, res := 0, 0

    for r := 0; r < len(nums); r++ {
        if nums[r] == 0 {
            k--
        }

        for k < 0 {
            if nums[l] == 0 {
                k++
            }

            l++
        }

        if res < r - l + 1 {
            res = r - l + 1
        }
    }  

    return res
}