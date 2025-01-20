func permute(nums []int) [][]int {
    ans := [][]int{}
    if len(nums) == 1 {
        return [][]int{nums}
    }

    for i := range nums {
        newNums := append(append([]int{}, nums[:i]...), nums[i+1:]...)
        perms := permute(newNums)
        for _, v := range perms {
            ans = append(ans, append(v, nums[i]))
        }
    }

    return ans
}