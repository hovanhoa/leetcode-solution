func removeDuplicates(nums []int) int {
    m := map[int]bool{}
    l := 0
    for i, v := range nums {
        if !m[v] {
            m[v] = true
            nums[l] = nums[i]
            l += 1
        }
    }

    return l
}