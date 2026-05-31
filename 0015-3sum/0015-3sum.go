import "slices"

func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    ans := [][]int{}
    i := 0
    fmt.Println(nums)
    for i <= len(nums) - 3 {
        left, right := i + 1, len(nums) - 1
        for left < right {
            s := nums[i] + nums[left] + nums[right]
            if s == 0 {
                fmt.Println(i, left, right)
                ans = append(ans, []int{nums[i], nums[left], nums[right]})
                left += 1
                for nums[left] == nums[left-1] && left < right {
                    left += 1
                }
            } else if s < 0 {
                left += 1
            } else {
                right -= 1
            }
        }

        i += 1
        for i <= len(nums) - 3 && nums[i] == nums[i-1] { 
            i += 1
        }
    }

    return ans
}
