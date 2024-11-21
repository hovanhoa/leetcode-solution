func permute(nums []int) [][]int {
    ans := [][]int{}
    if len(nums) == 1 {
        return [][]int{nums}
    }

    for i := 0; i < len(nums); i++ {
        n := nums[i]
        perms := permute(append(append([]int{}, nums[:i]...), nums[i+1:]...))

        for j := range perms {
            perms[j] = append(perms[j], n)
        }

        ans = append(ans, perms...)
    }

    return ans
}