func findDuplicate(nums []int) int {
    m := map[int]bool{}
    for _, v := range nums {
        if m[v] {
            return v
        }

        m[v] = true
    }

    return -1
}