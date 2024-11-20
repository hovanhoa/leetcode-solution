func threeSum(nums []int) [][]int {
    ans := [][]int{}
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    for i := 0; i < len(nums) - 2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
 
        l, r := i + 1, len(nums) - 1
        for l < r {
            total := nums[i] + nums[l] + nums[r]
            if total < 0 {
                l += 1
            } else if total > 0 {
                r -= 1
            } else {
                ans = append(ans, []int{nums[i], nums[l], nums[r]})
                l += 1
                for l < r && nums[l] == nums[l-1] {
                    l += 1
                }
            }
        }
    }

    return ans
}