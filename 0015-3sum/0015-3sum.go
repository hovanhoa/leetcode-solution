import "slices"

func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    res := [][]int{}

    for i := 0; i < len(nums) - 2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        l, r := i + 1, len(nums) - 1
        for l < r {
            target := nums[l] + nums[r] + nums[i]
            if target < 0 {
                l += 1
            } else if target > 0 {
                r -= 1
            } else {
                res = append(res, []int{nums[i], nums[l], nums[r]})
                l += 1
                for l < r && nums[l] == nums[l-1] {
                    l += 1
                }
            }
        }
    }

    return res
}