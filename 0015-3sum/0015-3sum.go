func threeSum(nums []int) [][]int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    ans := [][]int{}

    for i := 0; i < len(nums) - 2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        l, r := i + 1, len(nums) - 1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum == 0 {
                ans = append(ans, []int{nums[i], nums[l], nums[r]})

                l += 1
                for l < r && nums[l] == nums[l-1] {
                    l += 1
                }
            } else if sum < 0 {
                l += 1
            } else {
                r -= 1
            }
        }
    }

    return ans
}