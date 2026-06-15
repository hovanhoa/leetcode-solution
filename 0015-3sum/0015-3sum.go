import "slices"

func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    ans := [][]int{}
    i := 0
    for i <= len(nums) - 3 {
        l, r := i + 1, len(nums) - 1
        for l < r {
            s := nums[i] + nums[l] + nums[r]
            if s == 0 {
                ans = append(ans, []int{nums[i], nums[l], nums[r]})
                l += 1
                for l < r && nums[l] == nums[l-1] {
                    l += 1
                }
            } else if s < 0 {
                l += 1
            } else {
                r -= 1
            }
        }

        i += 1
        for i <= len(nums) - 3 && nums[i] == nums[i-1] {
            i += 1
        }
    }

    return ans    
}