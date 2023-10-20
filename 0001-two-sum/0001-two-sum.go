func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i := 0; i < len(nums); i ++ {
        j, ok := m[target - nums[i]]
        if ok{
            res := []int{j, i}
            return res
        } else {
            m[nums[i]] = i
        }
    }

    res := []int{-1, -1}
    return res
}