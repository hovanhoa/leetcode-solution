func permute(nums []int) [][]int {
    ans := [][]int{}
    if len(nums) == 1 {
        return [][]int{nums}
    }

    for i := 0; i < len(nums); i++ {
        newNums := append(append([]int{}, nums[:i]...), nums[i+1:]...)
        perms := permute(newNums)
        for j := range perms {
            ans = append(ans, append(perms[j], nums[i]))
        }
    }

    return ans
}