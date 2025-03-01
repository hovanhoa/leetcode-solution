func removeDuplicates(nums []int) int {
    m := map[int]int{}
    l := 0
    for _, v := range nums {
        if m[v] < 2 {
            m[v] += 1
            nums[l] = v
            l += 1
        }
    }

    return l
}